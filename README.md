# sing-box rules

基于最新 beta（sing-box v1.14.0-beta.14）的规则集生成器。把 **clash / host / qx** 三种行格式的
上游规则源，清洗、转换、去重、排除后，输出 sing-box rule-set 的两种产物：

- `<tag>.json` — source 格式（人类可读、可审查，version 5）
- `<tag>.srs` — 二进制格式（供 `rule_set` 的 local/remote 引用）

GitHub Actions 自动在每次推送 / 每日定时 / 手动触发时生成，产物提交到 **`rule-set` 分支根目录**
（官方命名，平铺无子目录），并附带 `README.md` 变更分析报告（GitHub 自动渲染）。

## 构建

```bash
go build -o sbrules.exe ./cmd/sbrules
```

## 用法

```bash
# 在线模式：抓取 .conf 中的 URL 上游
sbrules build --conf-dir . --out-dir rules

# 离线模式：从 <conf-dir>/examples/ 读取本地文件
sbrules build --conf-dir . --out-dir rules --examples

# 校验生成的 .srs（官方回读 + 版本升级检查）
sbrules check rules/binary/direct.srs
```

build 参数：

| 参数 | 默认 | 说明 |
|---|---|---|
| `--conf-dir` | `.` | 含 `*.conf` 的目录（扫描全部，不硬编码） |
| `--out-dir` | `rules` | 输出目录（下分 source/ binary/） |
| `--examples` | false | 用本地 examples/ 文件代替网络抓取 |
| `--strict` | false | 抓取失败即报错退出 |
| `--timeout` | 30s | HTTP 超时 |
| `--retries` | 2 | 每个 URL 重试次数 |

## .conf 格式（sb 风格 JSONC，支持 // 注释）

每个 `.conf` 文件 = 一个规则集定义，**文件名即默认 tag**；可用 `outputs` 按字段类型拆分成多个规则集。

```jsonc
// reject.conf → 输出 reject-domain.json/.srs + reject-ip.json/.srs
{
  "sources": [
    { "url": "https://example.com/filter.list", "format": "qx" }
  ],
  "head_rules": [
    { "query_type": ["PTR", "SVCB", "HTTPS", "ANY", "AXFR", "IXFR", "TSIG", "TKEY"] },
    { "domain_keyword": ["adserv", "advert"] }
  ],
  "outputs": [
    { "tag": "reject-domain", "fields": ["domain", "domain_suffix", "domain_keyword", "domain_regex", "query_type"] },
    { "tag": "reject-ip",     "fields": ["ip_cidr", "source_ip_cidr"] }
  ]
}
```

| 字段 | 说明 |
|---|---|
| `sources[]` | 上游源，`url`/`path` 二选一；`format` 缺省 `clash` |
| `head_rules[]` | 头部规则；**domain 系 / ip 系字段参与语义去重**，query_type / logical 保持独立输出在最前 |
| `exclude[]` | 排除规则，值级语义过滤（见下） |
| `outputs[]` | 输出变体（缺省 = 单输出 `<文件名>` 全部字段）；`fields` 白名单：只保留列出的字段，其余丢弃；空 = 全部保留 |

### 为什么要按类型拆分（domain / ip）

**DNS 规则引用的规则集若含 `ip_cidr` 项，会触发 Legacy 地址过滤模式**（sing-box 1.14 起弃用告警，1.16 硬失败，dns/router.go:1458 按规则集元数据 `ContainsIPCIDRRule` 判断，与是否混有域名无关）。因此：

- **route 引用**：域名 + IP 可共存（一个规则集全匹配）
- **dns 引用**：必须纯域名（查询拒绝、fakeip）或纯 IP + `match_response`（响应匹配）

本仓库产物：

| 规则集 | 内容 | route 引用 | dns 引用 |
|---|---|---|---|
| `direct` | 纯域名（精选直连） | 嗅探后直连 | — |
| `global-domain` | 纯域名（代理清单 + gfw） | 嗅探后代理 | fakeip |
| `global-ip` | 纯 IP（Proxy.list 自带 IP-CIDR） | 嗅探前代理 | — |
| `reject-domain` | 纯域名 + query_type | 嗅探后拒绝 | 查询拒绝 |
| `reject-ip` | 纯 IP（fmz200 广告 IP） | 嗅探前拒绝 | match_response 响应拒绝 |

## 输入格式

| format | 行示例 | → sing-box 字段 |
|---|---|---|
| `clash` | `DOMAIN,google.com` | `domain` |
| | `DOMAIN-SUFFIX,google.com` | `domain_suffix` |
| | `DOMAIN-KEYWORD,google` | `domain_keyword` |
| | `DOMAIN-REGEX,^google\.` | `domain_regex` |
| | `IP-CIDR,1.2.3.0/24` / `IP-CIDR6,2001:db8::/32`（`no-resolve` 剥除） | `ip_cidr` |
| | `SRC-IP-CIDR,127.0.0.0/8` | `source_ip_cidr` |
| | `DST-PORT,443` / `SRC-PORT,1080` | `port` / `source_port` |
| `host` | `google.com` / `.google.com`（自动剥点） | `domain_suffix` |
| | `1.2.3.0/24` / `8.8.8.8` / `ip-cidr,10.0.0.0/8` | `ip_cidr` |
| `qx` | `host,google.com,proxy` | `domain` |
| | `host-suffix,google.com,reject` | `domain_suffix` |
| | `host-keyword,google,reject` | `domain_keyword` |
| | `host-wildcard,*.google.com,proxy` | `domain_regex`（`*`→`.*`、`?`→`.`，补锚点） |
| | `host-regex,^stun\.,proxy` | `domain_regex` |
| | `IP-CIDR,1.2.3.0/24,no-resolve` / `ip6-cidr,...` | `ip_cidr`（自动剥策略名） |

格式扩展：`internal/transform` 的 `Register(name, Parser)` 注册表，新增格式只加一个文件。

## 白名单

规则只产出**全平台通用**字段：`domain` / `domain_suffix` / `domain_keyword` /
`domain_regex` / `ip_cidr` / `source_ip_cidr` / `port` / `source_port` /
`port_range` / `source_port_range` / `network` / `query_type`（DNS 专用，来自 conf 内联 head_rules）。

以下输入行被跳过并计入汇总报告（`skipped: geoip(1) process(1) ...`）：

| 类别 | 输入 | 说明 |
|---|---|---|
| `geoip` | `GEOIP,CN` | 建议引用官方 `sing-geoip-*` srs |
| `ip-asn` | `IP-ASN,13335` | 同上 |
| `process` | `PROCESS-NAME` 等 | 平台限定 |
| `unsupported` | `RULE-SET` / `MATCH` / `USER-AGENT` / `URL-REGEX` 等 | 无全平台等价 |
| `unknown` | 无法识别的行 | — |

`exclude` 与 `head_rules` 中若含白名单外字段，直接报错拒绝生成。

## 处理流水线

```
.conf(JSONC) 解析 → 并发抓取(worker pool, 重试) → 行清洗(注释/行尾注释)
  → 格式 Parser → option.HeadlessRule（此后不再碰字符串）
  → 字段归并(同类型合并成单条规则的大数组，head_rules 可归并组并入)
  → 文本去重 → 语义去重 → exclude 值级过滤（最后执行，控制最终可见结果）
  → outputs 字段白名单拆分（query_type / logical 独立 head_rules 保持最前）
  → 渲染: source JSON (v5) + binary SRS (v5, 官方 srs.Write)
  → 自检: 官方 srs.Read 回读 + Upgrade() 验证
```

### 语义去重

与 QX 对齐：每条先按 key 排序（域名按反转标签字典序，祖先在前），再单遍扫描——
每条只与**已保留的更靠前**条目比对，被覆盖即删除（首条保留）；输出保持原输入相对顺序。

| 场景 | 规则 |
|---|---|
| `domain_suffix` 父子域 | 祖先覆盖后代：`google.com` 删除 `sub.google.com` |
| `domain` | 被 keyword 包含、被 suffix 祖先覆盖或重复则删除 |
| `domain_keyword` | 被更早 keyword 子串覆盖则删除：`adsys` 删除 `adsystem` |
| `ip_cidr` | 宽网段包含窄网段时保留宽段（顺序无关）：`10.0.0.0/8` 删除 `10.1.0.0/16` |
| 其余字段 | 文本去重 |

### exclude 值级语义

`exclude` 作用于**值**而非整条规则（只删匹配的值，不动同规则其他值）：

- `domain: [x]` → 精确相等删除
- `domain_suffix: [x]` → x 及其所有子域删除
- `domain_keyword: [kw]` → 域名含 kw 删除
- `domain_regex: [re]` → 正则命中删除
- `ip_cidr: [cidr]` → 被该网段包含的 CIDR 删除
- `port` / `port_range` / `network` → 精确匹配删除

## 在 sing-box 中使用产物

产物在 `rule-set` 分支根目录（官方命名），可直接以 remote rule_set 引用：

```jsonc
{
  "route": {
    "rule_set": [
      {
        "type": "remote", "format": "binary",
        "url": "https://raw.githubusercontent.com/<user>/sing-box/rule-set/direct.srs",
        "tag": "direct", "update_interval": "1d"
      },
      {
        "type": "remote", "format": "source",
        "url": "https://raw.githubusercontent.com/<user>/sing-box/rule-set/reject.json",
        "tag": "reject", "update_interval": "1d"
      }
    ],
    "rules": [
      { "rule_set": ["direct"], "action": "route", "outbound": "direct" },
      { "rule_set": ["reject"], "action": "reject" }
    ]
  }
}
```

## GitHub Actions

`.github/workflows/generate.yml` 触发时机：

- 推送 `main` 分支
- 每周六定时（UTC 06:06）
- 手动触发（workflow_dispatch）

流程：checkout → 依赖更新到最新 + sing-box 固定 `SB_VERSION`（env 变量，现为
`v1.14.0-beta.14`，正式版发布后在 workflow 升级）→ `sbrules build --out-dir out --strict`
→ `release.sh` 计算每个产物文件的增删 delta → 切到 `rule-set` 分支（首次为 orphan）→
平铺产物到分支根目录并提交，无变化则不提交。

提交信息第一行为**总 values 增减**，后续每行按 **conf 规则集**统计（不再逐文件列）：

```
update rule sets (+5 -0)
direct: values 24 (+2 -0)
global-domain: values 4181 (+0 -0)
reject-domain: values 200482 (+3 -0)
reject-ip: values 29 (+0 -0)
```

首次运行（无上期报告）显示 `update rule sets (initial build, +N -N)`。

### README.md 变更报告（GitHub 自动渲染）

rule-set 分支的产物目录里附 `README.md`（GitHub 只在 README.md 上自动渲染，CHANGES.md 不会），含：

- 生成时间、固定的 sing-box 版本
- 总增减 + 单文件增减
- 各规则集统计表（规则数、匹配值数、跳过分类、体积）
- **sing-box 更新检查**：查询最新 stable，与固定版本做版本号比较（`sort -V`），
  仅当有真正更新时提示升级命令；固定 beta 比 stable 新时显示 "beta ahead"
- **AI 评估**（可选，**默认开启**）：用 **GitHub Copilot CLI** 生成变更评估。**无需额外 token**——
  GitHub Actions 内置 `GITHUB_TOKEN` 直接认证 Copilot，workflow 已声明
  `copilot-requests: write` 权限并自动安装 `@github/copilot` CLI。
  前提：你的账号/组织启用了 Copilot（含 Copilot Free 免费层）。

  唯一需配置的 Variables：`ENABLE_AI`（置 `false` 关闭，缺省开启）。
  其余行为由仓库内默认控制：

  | 项 | 默认 | 自定义方式 |
  |---|---|---|
  | 模型 | `auto`（Copilot 自动选择） | 改 release.sh 的 `AI_MODEL="${AI_MODEL:-auto}"` |
  | 提示词 | `.github/prompts/evaluate.md` | 直接编辑该文件（随 main 推送生效） |
  | 输出长度 | 2000 字符 | 改 release.sh 的 `AI_MAX_CHARS="${AI_MAX_CHARS:-2000}"` |

  AI 评估输出在 README.md **最前**（`## AI evaluation` 段），其余部分为程序生成。
  Copilot CLI 以 `--yolo` 非交互模式运行（见
  [官方文档](https://docs.github.com/en/copilot/how-tos/copilot-cli/use-copilot-cli-in-actions)）；
  评估失败不阻塞产物发布，该段落显示 "(AI evaluation unavailable)"。

  AI 只填充 README.md 的 **`## AI evaluation` 段落**，不会改写其他部分（版本/统计/文件表均为程序生成）。
  Copilot CLI 以 `--yolo` 非交互模式运行（见
  [官方文档](https://docs.github.com/en/copilot/how-tos/copilot-cli/use-copilot-cli-in-actions)）；
  评估失败不阻塞产物发布，该段落显示 "(AI evaluation unavailable)"。

> 仓库需在 GitHub 上创建后推送即可启用；`rule-set` 分支仅含产物，由 Actions 维护，请勿手动修改。

## DNS 与 Route 的规则集区别

**规则集文件本身是同一套格式**（headless rule，本仓库产出的产物两者通用），
区别在于**被引用时的匹配上下文**（sing-box v1.14.0-beta.14 源码 dns/router.go）：

| | Route 规则引用 `rule_set` | DNS 规则引用 `rule_set` |
|---|---|---|
| 匹配对象 | 连接（域名/IP/端口/协议） | DNS 查询（域名/来源/查询类型） |
| `domain_suffix` 等域名系 | 匹配连接目标域名 | 匹配被查询域名 |
| `ip_cidr`（无 `match_response`） | 匹配连接目标 IP | **Legacy 模式**：匹配查询响应 IP（已弃用，1.16 移除） |
| `ip_cidr` + `match_response` | 不适用 | 匹配 evaluate 求得的响应地址 |
| `query_type`（DNS 专用 headless 项） | 不适用 | 匹配查询类型；含该规则集时自动禁用 Legacy 模式 |
| 规则集更新校验 | 无额外约束 | 校验元数据（含 IP-CIDR 项 / 查询类型项）决定 DNS 模式 |

要点：

- **同一规则集可同时被 route 和 dns 规则引用**，但行为按上表区分。
- 本仓库产物按类型拆分：DNS 引用的规则集是纯域名（`*-domain`）或纯 IP（`*-ip`，配合 `match_response`），
  因此**不会触发 Legacy 告警**。
- 纯域名规则集（`*-domain`）在 DNS 规则中使用无此顾虑。

## 项目结构

```
cmd/sbrules/          CLI（build / check / diff）
internal/confparse/   .conf 解析（JSONC）
internal/fetch/       并发抓取（worker pool + 重试 + examples 模式）
internal/transform/   格式注册表 + clash / host / qx parser + 白名单校验
internal/dedup/       文本去重 + 语义去重（域名父子 / IP 子网 / keyword 覆盖）
internal/render/      归并、exclude、字段过滤（outputs）、source JSON / binary SRS 输出、回读自检
internal/pipeline/    编排：一个 .conf → 多个 output 规则集
.github/workflows/    generate.yml：自动生成并发布到 rule-set 分支
.github/scripts/      release.sh：delta 计算、README.md 生成、分支提交
direct.conf 等        规则集定义（文件名即 tag，outputs 可拆分多规则集）
```
