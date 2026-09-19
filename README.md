# Rule Set Changes

- generated: 2026-09-19 10:51 UTC
- sing-box pinned: v1.14.0-beta.14

## AI evaluation

**本期要点**：更新实质集中在 reject-domain，规则规模小幅增长；direct、global 与 reject-ip 基本保持不变。上游整体健康，新增内容已同步到源文件和二进制产物。

**质量与风险**：无异常。语义去重占比较高但未达到明显异常阈值；仅个别上游存在极少量解析跳过，未见失败、空壳或源文件与二进制体积不同步。

**建议**：暂不需要调整构建策略。持续关注 reject-domain 的语义去重比例及后续规模变化，并跟进存在少量跳过的上游条目。

## Summary

total: +1886890 -1868565

## Per File

```

direct.json: +0 -0
global-domain.json: +0 -0
global-ip.json: +0 -0
reject-domain.json: +2592 -467
reject-ip.json: +0 -0
direct.srs: +0 -0
global-domain.srs: +0 -0
global-ip.srs: +0 -0
reject-domain.srs: +1884298 -1868098
reject-ip.srs: +0 -0
```

## Rule Sets

| tag | rules | values | text-rm | sem-rm | excl-rm | skipped | src bytes | bin bytes | sources |
|---|---|---|---|---|---|---|---|---|---|
| direct | 1 | 24 | 0 | 0 | 0 |  | 847 | 333 | Direct.list(clash,ok,24/24) |
| global-domain | 1 | 4209 | 356 | 295 | 19 |  | 103197 | 28581 | Proxy.list(clash,ok,513/513); gfw.txt(host,ok,4384/4384) |
| global-ip | 1 | 18 | 356 | 295 | 19 |  | 563 | 155 | Proxy.list(clash,ok,513/513); gfw.txt(host,ok,4384/4384) |
| reject-domain | 1 | 226644 | 56962 | 108793 | 0 | unknown:9 | 7193047 | 1884298 | dns-filter.txt(host,ok,179698/179698); chinese-filter.txt(host,ok,6062/6062); AWAvenue-Ads-Rule-Surge.list(host,ok,948/948); domains.txt(host,ok,202953/202954); filter.list(qx,ok,2753/2761) |
| reject-ip | 1 | 29 | 56962 | 108793 | 0 | unknown:9 | 929 | 203 | dns-filter.txt(host,ok,179698/179698); chinese-filter.txt(host,ok,6062/6062); AWAvenue-Ads-Rule-Surge.list(host,ok,948/948); domains.txt(host,ok,202953/202954); filter.list(qx,ok,2753/2761) |

## Upstream Activity & Quality

### Per Rule Set (new vs previous)

| tag | values | text-rm | sem-rm | excl-rm | src bytes | bin bytes |
|---|---|---|---|---|---|---|
| direct | 24 (+0) | 0 (+0) | 0 (+0) | 0 (+0) | 847 (+0) | 333 (+0) |
| global-domain | 4209 (+0) | 356 (+0) | 295 (+0) | 19 (+0) | 103197 (+0) | 28581 (+0) |
| global-ip | 18 (+0) | 356 (+0) | 295 (+0) | 19 (+0) | 563 (+0) | 155 (+0) |
| reject-domain | 226644 (+2125) | 56962 (-325) | 108793 (+9) | 0 (+0) | 7193047 (+66033) | 1884298 (+16200) |
| reject-ip | 29 (+0) | 56962 (-325) | 108793 (+9) | 0 (+0) | 929 (+0) | 203 (+0) |

### Upstream Sources (activity & quality)

| rule set | source | fmt | status | lines | parsed | skip rate |
|---|---|---|---|---|---|---|
| direct | Direct.list | clash | ok | 24 | 24 | 0.0% |
| global-domain | Proxy.list | clash | ok | 513 | 513 | 0.0% |
| global-domain | gfw.txt | host | ok | 4384 | 4384 | 0.0% |
| global-ip | Proxy.list | clash | ok | 513 | 513 | 0.0% |
| global-ip | gfw.txt | host | ok | 4384 | 4384 | 0.0% |
| reject-domain | dns-filter.txt | host | ok | 179698 | 179698 | 0.0% |
| reject-domain | chinese-filter.txt | host | ok | 6062 | 6062 | 0.0% |
| reject-domain | AWAvenue-Ads-Rule-Surge.list | host | ok | 948 | 948 | 0.0% |
| reject-domain | domains.txt | host | ok | 202954 | 202953 | 0.0% |
| reject-domain | filter.list | qx | ok | 2761 | 2753 | 0.3% |
| reject-ip | dns-filter.txt | host | ok | 179698 | 179698 | 0.0% |
| reject-ip | chinese-filter.txt | host | ok | 6062 | 6062 | 0.0% |
| reject-ip | AWAvenue-Ads-Rule-Surge.list | host | ok | 948 | 948 | 0.0% |
| reject-ip | domains.txt | host | ok | 202954 | 202953 | 0.0% |
| reject-ip | filter.list | qx | ok | 2761 | 2753 | 0.3% |

## sing-box update check

- **update available**: pinned v1.14.0-beta.14, latest stable v1.14.1

> bump go.mod: `go get github.com/sagernet/sing-box@v1.14.1` and commit

<!--report:ewogICJnZW5lcmF0ZWRfYXQiOiAiMjAyNi0wOS0xOVQxMDo1MTowMS45NTYzMjU3OTlaIiwKICAic2luZ19ib3hfdmVyc2lvbiI6ICJ2MS4xNC4wLWJldGEuMTQiLAogICJydWxlX3NldHMiOiBbCiAgICB7CiAgICAgICJ0YWciOiAiZGlyZWN0IiwKICAgICAgInJ1bGVzIjogMSwKICAgICAgInZhbHVlcyI6IDI0LAogICAgICAidGV4dF9yZW1vdmVkIjogMCwKICAgICAgInNlbWFudGljX3JlbW92ZWQiOiAwLAogICAgICAiZXhjbHVkZV9yZW1vdmVkIjogMCwKICAgICAgInNvdXJjZV9ieXRlcyI6IDg0NywKICAgICAgImJpbmFyeV9ieXRlcyI6IDMzMywKICAgICAgInNvdXJjZXMiOiBbCiAgICAgICAgewogICAgICAgICAgIm5hbWUiOiAiaHR0cHM6Ly9yYXcuZ2l0aHVidXNlcmNvbnRlbnQuY29tL0Nvbm5lcnNIdWEvUnVsZUdvL21hc3Rlci9TdXJnZS9SdWxlc2V0L0RpcmVjdC5saXN0IiwKICAgICAgICAgICJmb3JtYXQiOiAiY2xhc2giLAogICAgICAgICAgIm9rIjogdHJ1ZSwKICAgICAgICAgICJsaW5lcyI6IDI0LAogICAgICAgICAgInBhcnNlZCI6IDI0CiAgICAgICAgfQogICAgICBdCiAgICB9LAogICAgewogICAgICAidGFnIjogImdsb2JhbC1kb21haW4iLAogICAgICAicnVsZXMiOiAxLAogICAgICAidmFsdWVzIjogNDIwOSwKICAgICAgInRleHRfcmVtb3ZlZCI6IDM1NiwKICAgICAgInNlbWFudGljX3JlbW92ZWQiOiAyOTUsCiAgICAgICJleGNsdWRlX3JlbW92ZWQiOiAxOSwKICAgICAgInNvdXJjZV9ieXRlcyI6IDEwMzE5NywKICAgICAgImJpbmFyeV9ieXRlcyI6IDI4NTgxLAogICAgICAic291cmNlcyI6IFsKICAgICAgICB7CiAgICAgICAgICAibmFtZSI6ICJodHRwczovL3Jhdy5naXRodWJ1c2VyY29udGVudC5jb20vQ29ubmVyc0h1YS9SdWxlR28vbWFzdGVyL1N1cmdlL1J1bGVzZXQvUHJveHkubGlzdCIsCiAgICAgICAgICAiZm9ybWF0IjogImNsYXNoIiwKICAgICAgICAgICJvayI6IHRydWUsCiAgICAgICAgICAibGluZXMiOiA1MTMsCiAgICAgICAgICAicGFyc2VkIjogNTEzCiAgICAgICAgfSwKICAgICAgICB7CiAgICAgICAgICAibmFtZSI6ICJodHRwczovL3Jhdy5naXRodWJ1c2VyY29udGVudC5jb20vTG95YWxzb2xkaWVyL3N1cmdlLXJ1bGVzL3JlbGVhc2UvZ2Z3LnR4dCIsCiAgICAgICAgICAiZm9ybWF0IjogImhvc3QiLAogICAgICAgICAgIm9rIjogdHJ1ZSwKICAgICAgICAgICJsaW5lcyI6IDQzODQsCiAgICAgICAgICAicGFyc2VkIjogNDM4NAogICAgICAgIH0KICAgICAgXQogICAgfSwKICAgIHsKICAgICAgInRhZyI6ICJnbG9iYWwtaXAiLAogICAgICAicnVsZXMiOiAxLAogICAgICAidmFsdWVzIjogMTgsCiAgICAgICJ0ZXh0X3JlbW92ZWQiOiAzNTYsCiAgICAgICJzZW1hbnRpY19yZW1vdmVkIjogMjk1LAogICAgICAiZXhjbHVkZV9yZW1vdmVkIjogMTksCiAgICAgICJzb3VyY2VfYnl0ZXMiOiA1NjMsCiAgICAgICJiaW5hcnlfYnl0ZXMiOiAxNTUsCiAgICAgICJzb3VyY2VzIjogWwogICAgICAgIHsKICAgICAgICAgICJuYW1lIjogImh0dHBzOi8vcmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbS9Db25uZXJzSHVhL1J1bGVHby9tYXN0ZXIvU3VyZ2UvUnVsZXNldC9Qcm94eS5saXN0IiwKICAgICAgICAgICJmb3JtYXQiOiAiY2xhc2giLAogICAgICAgICAgIm9rIjogdHJ1ZSwKICAgICAgICAgICJsaW5lcyI6IDUxMywKICAgICAgICAgICJwYXJzZWQiOiA1MTMKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICJuYW1lIjogImh0dHBzOi8vcmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbS9Mb3lhbHNvbGRpZXIvc3VyZ2UtcnVsZXMvcmVsZWFzZS9nZncudHh0IiwKICAgICAgICAgICJmb3JtYXQiOiAiaG9zdCIsCiAgICAgICAgICAib2siOiB0cnVlLAogICAgICAgICAgImxpbmVzIjogNDM4NCwKICAgICAgICAgICJwYXJzZWQiOiA0Mzg0CiAgICAgICAgfQogICAgICBdCiAgICB9LAogICAgewogICAgICAidGFnIjogInJlamVjdC1kb21haW4iLAogICAgICAicnVsZXMiOiAxLAogICAgICAidmFsdWVzIjogMjI2NjQ0LAogICAgICAic2tpcHBlZCI6IHsKICAgICAgICAidW5rbm93biI6IDkKICAgICAgfSwKICAgICAgInRleHRfcmVtb3ZlZCI6IDU2OTYyLAogICAgICAic2VtYW50aWNfcmVtb3ZlZCI6IDEwODc5MywKICAgICAgImV4Y2x1ZGVfcmVtb3ZlZCI6IDAsCiAgICAgICJzb3VyY2VfYnl0ZXMiOiA3MTkzMDQ3LAogICAgICAiYmluYXJ5X2J5dGVzIjogMTg4NDI5OCwKICAgICAgInNvdXJjZXMiOiBbCiAgICAgICAgewogICAgICAgICAgIm5hbWUiOiAiaHR0cHM6Ly9yYXcuZ2l0aHVidXNlcmNvbnRlbnQuY29tL2dlZWtkYWRhL3N1cmdlLWxpc3QvbWFzdGVyL2RvbWFpbi1zZXQvZG5zLWZpbHRlci50eHQiLAogICAgICAgICAgImZvcm1hdCI6ICJob3N0IiwKICAgICAgICAgICJvayI6IHRydWUsCiAgICAgICAgICAibGluZXMiOiAxNzk2OTgsCiAgICAgICAgICAicGFyc2VkIjogMTc5Njk4CiAgICAgICAgfSwKICAgICAgICB7CiAgICAgICAgICAibmFtZSI6ICJodHRwczovL3Jhdy5naXRodWJ1c2VyY29udGVudC5jb20vZ2Vla2RhZGEvc3VyZ2UtbGlzdC9tYXN0ZXIvZG9tYWluLXNldC9jaGluZXNlLWZpbHRlci50eHQiLAogICAgICAgICAgImZvcm1hdCI6ICJob3N0IiwKICAgICAgICAgICJvayI6IHRydWUsCiAgICAgICAgICAibGluZXMiOiA2MDYyLAogICAgICAgICAgInBhcnNlZCI6IDYwNjIKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICJuYW1lIjogImh0dHBzOi8vcmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbS9URy1Ud2lsaWdodC9BV0F2ZW51ZS1BZHMtUnVsZS9tYWluL0ZpbHRlcnMvQVdBdmVudWUtQWRzLVJ1bGUtU3VyZ2UubGlzdCIsCiAgICAgICAgICAiZm9ybWF0IjogImhvc3QiLAogICAgICAgICAgIm9rIjogdHJ1ZSwKICAgICAgICAgICJsaW5lcyI6IDk0OCwKICAgICAgICAgICJwYXJzZWQiOiA5NDgKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICJuYW1lIjogImh0dHBzOi8vcmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbS9iYWRtb2pyLzFIb3N0cy9tYXN0ZXIvTGl0ZS9kb21haW5zLnR4dCIsCiAgICAgICAgICAiZm9ybWF0IjogImhvc3QiLAogICAgICAgICAgIm9rIjogdHJ1ZSwKICAgICAgICAgICJsaW5lcyI6IDIwMjk1NCwKICAgICAgICAgICJwYXJzZWQiOiAyMDI5NTMsCiAgICAgICAgICAic2tpcHBlZCI6IHsKICAgICAgICAgICAgInVua25vd24iOiAxCiAgICAgICAgICB9CiAgICAgICAgfSwKICAgICAgICB7CiAgICAgICAgICAibmFtZSI6ICJodHRwczovL3Jhdy5naXRodWJ1c2VyY29udGVudC5jb20vZm16MjAwL3dvb2xfc2NyaXB0cy9tYWluL1F1YW50dW11bHRYL2ZpbHRlci9maWx0ZXIubGlzdCIsCiAgICAgICAgICAiZm9ybWF0IjogInF4IiwKICAgICAgICAgICJvayI6IHRydWUsCiAgICAgICAgICAibGluZXMiOiAyNzYxLAogICAgICAgICAgInBhcnNlZCI6IDI3NTMsCiAgICAgICAgICAic2tpcHBlZCI6IHsKICAgICAgICAgICAgInVua25vd24iOiA4CiAgICAgICAgICB9CiAgICAgICAgfQogICAgICBdCiAgICB9LAogICAgewogICAgICAidGFnIjogInJlamVjdC1pcCIsCiAgICAgICJydWxlcyI6IDEsCiAgICAgICJ2YWx1ZXMiOiAyOSwKICAgICAgInNraXBwZWQiOiB7CiAgICAgICAgInVua25vd24iOiA5CiAgICAgIH0sCiAgICAgICJ0ZXh0X3JlbW92ZWQiOiA1Njk2MiwKICAgICAgInNlbWFudGljX3JlbW92ZWQiOiAxMDg3OTMsCiAgICAgICJleGNsdWRlX3JlbW92ZWQiOiAwLAogICAgICAic291cmNlX2J5dGVzIjogOTI5LAogICAgICAiYmluYXJ5X2J5dGVzIjogMjAzLAogICAgICAic291cmNlcyI6IFsKICAgICAgICB7CiAgICAgICAgICAibmFtZSI6ICJodHRwczovL3Jhdy5naXRodWJ1c2VyY29udGVudC5jb20vZ2Vla2RhZGEvc3VyZ2UtbGlzdC9tYXN0ZXIvZG9tYWluLXNldC9kbnMtZmlsdGVyLnR4dCIsCiAgICAgICAgICAiZm9ybWF0IjogImhvc3QiLAogICAgICAgICAgIm9rIjogdHJ1ZSwKICAgICAgICAgICJsaW5lcyI6IDE3OTY5OCwKICAgICAgICAgICJwYXJzZWQiOiAxNzk2OTgKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICJuYW1lIjogImh0dHBzOi8vcmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbS9nZWVrZGFkYS9zdXJnZS1saXN0L21hc3Rlci9kb21haW4tc2V0L2NoaW5lc2UtZmlsdGVyLnR4dCIsCiAgICAgICAgICAiZm9ybWF0IjogImhvc3QiLAogICAgICAgICAgIm9rIjogdHJ1ZSwKICAgICAgICAgICJsaW5lcyI6IDYwNjIsCiAgICAgICAgICAicGFyc2VkIjogNjA2MgogICAgICAgIH0sCiAgICAgICAgewogICAgICAgICAgIm5hbWUiOiAiaHR0cHM6Ly9yYXcuZ2l0aHVidXNlcmNvbnRlbnQuY29tL1RHLVR3aWxpZ2h0L0FXQXZlbnVlLUFkcy1SdWxlL21haW4vRmlsdGVycy9BV0F2ZW51ZS1BZHMtUnVsZS1TdXJnZS5saXN0IiwKICAgICAgICAgICJmb3JtYXQiOiAiaG9zdCIsCiAgICAgICAgICAib2siOiB0cnVlLAogICAgICAgICAgImxpbmVzIjogOTQ4LAogICAgICAgICAgInBhcnNlZCI6IDk0OAogICAgICAgIH0sCiAgICAgICAgewogICAgICAgICAgIm5hbWUiOiAiaHR0cHM6Ly9yYXcuZ2l0aHVidXNlcmNvbnRlbnQuY29tL2JhZG1vanIvMUhvc3RzL21hc3Rlci9MaXRlL2RvbWFpbnMudHh0IiwKICAgICAgICAgICJmb3JtYXQiOiAiaG9zdCIsCiAgICAgICAgICAib2siOiB0cnVlLAogICAgICAgICAgImxpbmVzIjogMjAyOTU0LAogICAgICAgICAgInBhcnNlZCI6IDIwMjk1MywKICAgICAgICAgICJza2lwcGVkIjogewogICAgICAgICAgICAidW5rbm93biI6IDEKICAgICAgICAgIH0KICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICJuYW1lIjogImh0dHBzOi8vcmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbS9mbXoyMDAvd29vbF9zY3JpcHRzL21haW4vUXVhbnR1bXVsdFgvZmlsdGVyL2ZpbHRlci5saXN0IiwKICAgICAgICAgICJmb3JtYXQiOiAicXgiLAogICAgICAgICAgIm9rIjogdHJ1ZSwKICAgICAgICAgICJsaW5lcyI6IDI3NjEsCiAgICAgICAgICAicGFyc2VkIjogMjc1MywKICAgICAgICAgICJza2lwcGVkIjogewogICAgICAgICAgICAidW5rbm93biI6IDgKICAgICAgICAgIH0KICAgICAgICB9CiAgICAgIF0KICAgIH0KICBdCn0=-->