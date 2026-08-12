#!/usr/bin/env bash
# 将生成产物提交到 rule-set 分支根目录。
# 输入：out/ 目录（source/*.json binary/*.srs report.json rulesets.md），在 main 工作树生成。
# 产物：rule-set 分支根目录直接放 <tag>.json <tag>.srs + CHANGES.md
# 依赖：jq 仅用于可选 AI 步骤（GitHub Actions runner 自带）。
set -e

git config user.name "github-actions[bot]"
git config user.email "github-actions[bot]@users.noreply.github.com"

PINNED=$(sed -n 's/.*github.com\/sagernet\/sing-box \([^ ]*\).*/\1/p' go.mod)

# ── 1. 计算每个产物文件的 delta（相对 origin/rule-set） ──
git fetch origin rule-set 2>/dev/null || true
# 远端无 rule-set 时清理本地 stale ref，避免误判旧文件存在
if ! git ls-remote --exit-code origin rule-set >/dev/null 2>&1; then
  git update-ref -d refs/remotes/origin/rule-set 2>/dev/null || true
fi

total_add=0
total_del=0
file_lines=""
for f in out/source/*.json out/binary/*.srs; do
  [ -f "$f" ] || continue
  name=$(basename "$f")
  if git show-ref --verify --quiet refs/remotes/origin/rule-set; then
    git show "origin/rule-set:$name" > /tmp/old_file 2>/dev/null || rm -f /tmp/old_file
  else
    rm -f /tmp/old_file
  fi
  add=0
  del=0
  if [ ! -f /tmp/old_file ]; then
    add=$(wc -c < "$f")
  elif [ "$(cat /tmp/old_file | md5sum | cut -d' ' -f1)" = "$(md5sum < "$f" | cut -d' ' -f1)" ]; then
    add=0
    del=0
  else
    case "$f" in
      *.json)
        add=$(diff /tmp/old_file "$f" 2>/dev/null | grep -c '^>' || true)
        del=$(diff /tmp/old_file "$f" 2>/dev/null | grep -c '^<' || true)
        ;;
      *.srs)
        add=$(wc -c < "$f")
        del=$(wc -c < /tmp/old_file)
        ;;
    esac
  fi
  rm -f /tmp/old_file
  total_add=$((total_add + add))
  total_del=$((total_del + del))
  file_lines="$file_lines"$'\n'"$name: +$add -$del"
done

# ── 2. 生成 README.md（GitHub 会渲染 README.md，CHANGES.md 不会） ──

# 上一期报告从旧 README.md 的隐藏注释中提取（不生成独立 report.json 产物）
DIFF_SECTION=""
SUMMARY_LINES=""
if git show-ref --verify --quiet refs/remotes/origin/rule-set; then
  if git show "origin/rule-set:README.md" 2>/dev/null | sed -n 's/.*<!--report:\([^>]*\)-->.*/\1/p' | base64 -d > /tmp/old_report.json 2>/dev/null; then
    if [ -s /tmp/old_report.json ]; then
      DIFF_SECTION=$(./sbrules diff /tmp/old_report.json out/report.json 2>/dev/null || true)
      SUMMARY_LINES=$(./sbrules diff --summary /tmp/old_report.json out/report.json 2>/dev/null || true)
    fi
  fi
fi

# 先执行 AI 评估（结果放最前）；模型默认 auto（Copilot 自动选择）。
# 可控项（均为 env 变量，可覆盖）：
#   AI_PROMPT     自定义审阅指令（缺省用 .github/prompts/evaluate.md）
#   AI_MODEL      模型名（默认 auto）
#   AI_MAX_CHARS  输出最大字符数（默认 2000）
# 失败不阻塞发布（显示 unavailable）。
# 认证：CI 用 GITHUB_TOKEN（需 copilot-requests: write 权限）；
# 本地运行无 GITHUB_TOKEN 时尝试 copilot 已保存的 OAuth 凭据。
AI_SECTION=""
if [ "${ENABLE_AI:-true}" = "true" ]; then
  if ! command -v copilot >/dev/null 2>&1; then
    AI_SECTION="(unavailable: @github/copilot CLI not installed)"
  else
    AI_MAX_CHARS="${AI_MAX_CHARS:-2000}"
    AI_MODEL="${AI_MODEL:-auto}"
    if [ -n "${AI_PROMPT:-}" ]; then
      prompt="$AI_PROMPT"
    elif [ -f .github/prompts/evaluate.md ]; then
      prompt=$(cat .github/prompts/evaluate.md)
    else
      prompt=""
    fi
    summary=$(cat <<EOF
sing-box pinned: $PINNED
total: +$total_add -$total_del
$(printf '%s\n' "$file_lines")
--- current rule sets ---
$(cat out/rulesets.md)
--- upstream activity & quality (vs previous run) ---
$DIFF_SECTION
EOF
)
    copilot_args=(--yolo --model "$AI_MODEL")
    if [ -n "$prompt" ]; then
      resp=$(copilot "${copilot_args[@]}" -p "$prompt"$'\n\n'"$summary" 2>/dev/null | head -c "$AI_MAX_CHARS") || resp=""
    else
      resp=$(copilot "${copilot_args[@]}" -p "$summary" 2>/dev/null | head -c "$AI_MAX_CHARS") || resp=""
    fi
    if [ -n "$resp" ]; then
      AI_SECTION="$resp"
    else
      AI_SECTION="(AI evaluation unavailable)"
    fi
  fi
else
  AI_SECTION="(disabled by ENABLE_AI=false)"
fi

{
  echo "# Rule Set Changes"
  echo
  echo "- generated: $(date -u '+%Y-%m-%d %H:%M UTC')"
  echo "- sing-box pinned: $PINNED"
  echo
  echo "## AI evaluation"
  echo
  printf '%s\n' "$AI_SECTION"
  echo
  echo "## Summary"
  echo
  echo "total: +$total_add -$total_del"
  echo
  echo "## Per File"
  echo
  echo '```'
  printf '%s\n' "$file_lines"
  echo '```'
  echo
  echo "## Rule Sets"
  echo
  cat out/rulesets.md
  echo
  echo "## Upstream Activity & Quality"
  echo
  if [ -n "$DIFF_SECTION" ]; then
    printf '%s\n' "$DIFF_SECTION"
  else
    echo "(no previous report available yet)"
  fi
  echo
  echo "## sing-box update check"
  echo
  # 固定 beta 版本；更新检查仅跟踪正式版（releases/latest 不含 prerelease）。
  # 当高于当前 pinned 的正式版发布时提示升级。
  latest=$(curl -fsSL --max-time 15 https://api.github.com/repos/SagerNet/sing-box/releases/latest 2>/dev/null | sed -n 's/.*"tag_name": *"\([^"]*\)".*/\1/p' | head -1 || true)
  if [ -n "$latest" ] && [ "$latest" != "$PINNED" ]; then
    # 版本号大小比较：beta 固定时只提示正式版更新，不因更高 beta 打扰
    newer=$(printf '%s\n%s\n' "$PINNED" "$latest" | sort -V | tail -1)
    if [ "$newer" = "$latest" ]; then
      echo "- **update available**: pinned $PINNED, latest stable $latest"
      echo
      echo "> bump go.mod: \`go get github.com/sagernet/sing-box@$latest\` and commit"
    else
      echo "- pinned $PINNED (beta), latest stable $latest, beta ahead of stable (no action)"
    fi
  else
    echo "- pinned version is up to date"
  fi
  # 隐藏注释：嵌入本机 report.json（base64），供下一期 diff 对比使用
  echo
  printf '<!--report:%s-->' "$(base64 -w0 out/report.json)"
} > out/README.md

# ── 3. 产物先搬出工作树，再切换分支 ──
TMP_ARTIFACTS=$(mktemp -d)
cp out/source/*.json out/binary/*.srs out/README.md "$TMP_ARTIFACTS"/

rm -rf rules out sbrules sbrules.exe  # 清理未跟踪产物，避免 checkout 冲突
if git show-ref --verify --quiet refs/remotes/origin/rule-set; then
  git checkout -B rule-set origin/rule-set
else
  git checkout --orphan rule-set
  git rm -rf . 2>/dev/null || true
fi

# ── 4. 产物平铺到分支根目录并提交 ──
cp "$TMP_ARTIFACTS"/*.json "$TMP_ARTIFACTS"/*.srs "$TMP_ARTIFACTS"/README.md .
rm -rf "$TMP_ARTIFACTS"

# 只添加预期产物文件（避免 orphan 分支残留的未跟踪文件被误提交）
git add -f *.json *.srs README.md
if git diff --cached --quiet; then
  echo "no changes to commit"
else
  # commit message：第一行总 values 增减，后续每行按 conf（规则集）统计
  if [ -n "$SUMMARY_LINES" ]; then
    first=$(printf '%s\n' "$SUMMARY_LINES" | head -1)
    msg="update rule sets ($first)"
    rest=$(printf '%s\n' "$SUMMARY_LINES" | tail -n +2)
    while IFS= read -r line; do
      [ -n "$line" ] && msg="$msg"$'\n'"$line"
    done <<< "$rest"
  else
    msg="update rule sets (initial build, +$total_add -$total_del)"
  fi
  git commit -m "$msg"
  git push origin rule-set --force
fi
