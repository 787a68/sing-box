# Rule Set Changes

- generated: 2026-10-31 12:14 UTC
- sing-box pinned: v1.14.0-beta.14

## AI evaluation

**本期要点**：规则内容整体稳定，global-domain 略有扩充、reject-domain 小幅收缩；二进制文件变化远大于规则本身的变化。

**质量与风险**：上游均正常，少量解析跳过不构成异常；语义去重占比较高但未过半，且变化平稳。无明显误删信号。

**建议**：核查二进制文件大幅变化是否由构建或编码差异引起；无需因本期上游状态调整源顺序。

## Summary

total: +2062425 -2061561

## Per File

```

direct.json: +0 -0
global-domain.json: +14 -0
global-ip.json: +0 -0
reject-domain.json: +12360 -13000
reject-ip.json: +0 -0
direct.srs: +0 -0
global-domain.srs: +29003 -28869
global-ip.srs: +0 -0
reject-domain.srs: +2021048 -2019692
reject-ip.srs: +0 -0
```

## Rule Sets

| tag | rules | values | text-rm | sem-rm | excl-rm | skipped | src bytes | bin bytes | sources |
|---|---|---|---|---|---|---|---|---|---|
| direct | 1 | 24 | 0 | 0 | 0 |  | 847 | 333 | Direct.list(clash,ok,24/24) |
| global-domain | 1 | 4274 | 357 | 295 | 19 |  | 104749 | 29003 | Proxy.list(clash,ok,513/513); gfw.txt(host,ok,4448/4448) |
| global-ip | 1 | 18 | 357 | 295 | 19 |  | 563 | 155 | Proxy.list(clash,ok,513/513); gfw.txt(host,ok,4448/4448) |
| reject-domain | 1 | 241552 | 47092 | 108932 | 0 | unknown:9 | 7691865 | 2021048 | dns-filter.txt(host,ok,184844/184844); chinese-filter.txt(host,ok,6069/6069); AWAvenue-Ads-Rule-Surge.list(host,ok,972/972); domains.txt(host,ok,202953/202954); filter.list(qx,ok,2753/2761) |
| reject-ip | 1 | 29 | 47092 | 108932 | 0 | unknown:9 | 929 | 203 | dns-filter.txt(host,ok,184844/184844); chinese-filter.txt(host,ok,6069/6069); AWAvenue-Ads-Rule-Surge.list(host,ok,972/972); domains.txt(host,ok,202953/202954); filter.list(qx,ok,2753/2761) |

## Upstream Activity & Quality

### Per Rule Set (new vs previous)

| tag | values | text-rm | sem-rm | excl-rm | src bytes | bin bytes |
|---|---|---|---|---|---|---|
| direct | 24 (+0) | 0 (+0) | 0 (+0) | 0 (+0) | 847 (+0) | 333 (+0) |
| global-domain | 4274 (+14) | 357 (+0) | 295 (+0) | 19 (+0) | 104749 (+350) | 29003 (+134) |
| global-ip | 18 (+0) | 357 (+0) | 295 (+0) | 19 (+0) | 563 (+0) | 155 (+0) |
| reject-domain | 241552 (-640) | 47092 (-4785) | 108932 (+22) | 0 (+0) | 7691865 (-339) | 2021048 (+1356) |
| reject-ip | 29 (+0) | 47092 (-4785) | 108932 (+22) | 0 (+0) | 929 (+0) | 203 (+0) |

### Upstream Sources (activity & quality)

| rule set | source | fmt | status | lines | parsed | skip rate |
|---|---|---|---|---|---|---|
| direct | Direct.list | clash | ok | 24 | 24 | 0.0% |
| global-domain | Proxy.list | clash | ok | 513 | 513 | 0.0% |
| global-domain | gfw.txt | host | ok | 4448 | 4448 | 0.0% |
| global-ip | Proxy.list | clash | ok | 513 | 513 | 0.0% |
| global-ip | gfw.txt | host | ok | 4448 | 4448 | 0.0% |
| reject-domain | dns-filter.txt | host | ok | 184844 | 184844 | 0.0% |
| reject-domain | chinese-filter.txt | host | ok | 6069 | 6069 | 0.0% |
| reject-domain | AWAvenue-Ads-Rule-Surge.list | host | ok | 972 | 972 | 0.0% |
| reject-domain | domains.txt | host | ok | 202954 | 202953 | 0.0% |
| reject-domain | filter.list | qx | ok | 2761 | 2753 | 0.3% |
| reject-ip | dns-filter.txt | host | ok | 184844 | 184844 | 0.0% |
| reject-ip | chinese-filter.txt | host | ok | 6069 | 6069 | 0.0% |
| reject-ip | AWAvenue-Ads-Rule-Surge.list | host | ok | 972 | 972 | 0.0% |
| reject-ip | domains.txt | host | ok | 202954 | 202953 | 0.0% |
| reject-ip | filter.list | qx | ok | 2761 | 2753 | 0.3% |

## sing-box update check

- **update available**: pinned v1.14.0-beta.14, latest stable v1.14.4

> bump go.mod: `go get github.com/sagernet/sing-box@v1.14.4` and commit

<!--report:ewogICJnZW5lcmF0ZWRfYXQiOiAiMjAyNi0xMC0zMVQxMjoxNDozOS4yNjU5NzEwMzlaIiwKICAic2luZ19ib3hfdmVyc2lvbiI6ICJ2MS4xNC4wLWJldGEuMTQiLAogICJydWxlX3NldHMiOiBbCiAgICB7CiAgICAgICJ0YWciOiAiZGlyZWN0IiwKICAgICAgInJ1bGVzIjogMSwKICAgICAgInZhbHVlcyI6IDI0LAogICAgICAidGV4dF9yZW1vdmVkIjogMCwKICAgICAgInNlbWFudGljX3JlbW92ZWQiOiAwLAogICAgICAiZXhjbHVkZV9yZW1vdmVkIjogMCwKICAgICAgInNvdXJjZV9ieXRlcyI6IDg0NywKICAgICAgImJpbmFyeV9ieXRlcyI6IDMzMywKICAgICAgInNvdXJjZXMiOiBbCiAgICAgICAgewogICAgICAgICAgIm5hbWUiOiAiaHR0cHM6Ly9yYXcuZ2l0aHVidXNlcmNvbnRlbnQuY29tL0Nvbm5lcnNIdWEvUnVsZUdvL21hc3Rlci9TdXJnZS9SdWxlc2V0L0RpcmVjdC5saXN0IiwKICAgICAgICAgICJmb3JtYXQiOiAiY2xhc2giLAogICAgICAgICAgIm9rIjogdHJ1ZSwKICAgICAgICAgICJsaW5lcyI6IDI0LAogICAgICAgICAgInBhcnNlZCI6IDI0CiAgICAgICAgfQogICAgICBdCiAgICB9LAogICAgewogICAgICAidGFnIjogImdsb2JhbC1kb21haW4iLAogICAgICAicnVsZXMiOiAxLAogICAgICAidmFsdWVzIjogNDI3NCwKICAgICAgInRleHRfcmVtb3ZlZCI6IDM1NywKICAgICAgInNlbWFudGljX3JlbW92ZWQiOiAyOTUsCiAgICAgICJleGNsdWRlX3JlbW92ZWQiOiAxOSwKICAgICAgInNvdXJjZV9ieXRlcyI6IDEwNDc0OSwKICAgICAgImJpbmFyeV9ieXRlcyI6IDI5MDAzLAogICAgICAic291cmNlcyI6IFsKICAgICAgICB7CiAgICAgICAgICAibmFtZSI6ICJodHRwczovL3Jhdy5naXRodWJ1c2VyY29udGVudC5jb20vQ29ubmVyc0h1YS9SdWxlR28vbWFzdGVyL1N1cmdlL1J1bGVzZXQvUHJveHkubGlzdCIsCiAgICAgICAgICAiZm9ybWF0IjogImNsYXNoIiwKICAgICAgICAgICJvayI6IHRydWUsCiAgICAgICAgICAibGluZXMiOiA1MTMsCiAgICAgICAgICAicGFyc2VkIjogNTEzCiAgICAgICAgfSwKICAgICAgICB7CiAgICAgICAgICAibmFtZSI6ICJodHRwczovL3Jhdy5naXRodWJ1c2VyY29udGVudC5jb20vTG95YWxzb2xkaWVyL3N1cmdlLXJ1bGVzL3JlbGVhc2UvZ2Z3LnR4dCIsCiAgICAgICAgICAiZm9ybWF0IjogImhvc3QiLAogICAgICAgICAgIm9rIjogdHJ1ZSwKICAgICAgICAgICJsaW5lcyI6IDQ0NDgsCiAgICAgICAgICAicGFyc2VkIjogNDQ0OAogICAgICAgIH0KICAgICAgXQogICAgfSwKICAgIHsKICAgICAgInRhZyI6ICJnbG9iYWwtaXAiLAogICAgICAicnVsZXMiOiAxLAogICAgICAidmFsdWVzIjogMTgsCiAgICAgICJ0ZXh0X3JlbW92ZWQiOiAzNTcsCiAgICAgICJzZW1hbnRpY19yZW1vdmVkIjogMjk1LAogICAgICAiZXhjbHVkZV9yZW1vdmVkIjogMTksCiAgICAgICJzb3VyY2VfYnl0ZXMiOiA1NjMsCiAgICAgICJiaW5hcnlfYnl0ZXMiOiAxNTUsCiAgICAgICJzb3VyY2VzIjogWwogICAgICAgIHsKICAgICAgICAgICJuYW1lIjogImh0dHBzOi8vcmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbS9Db25uZXJzSHVhL1J1bGVHby9tYXN0ZXIvU3VyZ2UvUnVsZXNldC9Qcm94eS5saXN0IiwKICAgICAgICAgICJmb3JtYXQiOiAiY2xhc2giLAogICAgICAgICAgIm9rIjogdHJ1ZSwKICAgICAgICAgICJsaW5lcyI6IDUxMywKICAgICAgICAgICJwYXJzZWQiOiA1MTMKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICJuYW1lIjogImh0dHBzOi8vcmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbS9Mb3lhbHNvbGRpZXIvc3VyZ2UtcnVsZXMvcmVsZWFzZS9nZncudHh0IiwKICAgICAgICAgICJmb3JtYXQiOiAiaG9zdCIsCiAgICAgICAgICAib2siOiB0cnVlLAogICAgICAgICAgImxpbmVzIjogNDQ0OCwKICAgICAgICAgICJwYXJzZWQiOiA0NDQ4CiAgICAgICAgfQogICAgICBdCiAgICB9LAogICAgewogICAgICAidGFnIjogInJlamVjdC1kb21haW4iLAogICAgICAicnVsZXMiOiAxLAogICAgICAidmFsdWVzIjogMjQxNTUyLAogICAgICAic2tpcHBlZCI6IHsKICAgICAgICAidW5rbm93biI6IDkKICAgICAgfSwKICAgICAgInRleHRfcmVtb3ZlZCI6IDQ3MDkyLAogICAgICAic2VtYW50aWNfcmVtb3ZlZCI6IDEwODkzMiwKICAgICAgImV4Y2x1ZGVfcmVtb3ZlZCI6IDAsCiAgICAgICJzb3VyY2VfYnl0ZXMiOiA3NjkxODY1LAogICAgICAiYmluYXJ5X2J5dGVzIjogMjAyMTA0OCwKICAgICAgInNvdXJjZXMiOiBbCiAgICAgICAgewogICAgICAgICAgIm5hbWUiOiAiaHR0cHM6Ly9yYXcuZ2l0aHVidXNlcmNvbnRlbnQuY29tL2dlZWtkYWRhL3N1cmdlLWxpc3QvbWFzdGVyL2RvbWFpbi1zZXQvZG5zLWZpbHRlci50eHQiLAogICAgICAgICAgImZvcm1hdCI6ICJob3N0IiwKICAgICAgICAgICJvayI6IHRydWUsCiAgICAgICAgICAibGluZXMiOiAxODQ4NDQsCiAgICAgICAgICAicGFyc2VkIjogMTg0ODQ0CiAgICAgICAgfSwKICAgICAgICB7CiAgICAgICAgICAibmFtZSI6ICJodHRwczovL3Jhdy5naXRodWJ1c2VyY29udGVudC5jb20vZ2Vla2RhZGEvc3VyZ2UtbGlzdC9tYXN0ZXIvZG9tYWluLXNldC9jaGluZXNlLWZpbHRlci50eHQiLAogICAgICAgICAgImZvcm1hdCI6ICJob3N0IiwKICAgICAgICAgICJvayI6IHRydWUsCiAgICAgICAgICAibGluZXMiOiA2MDY5LAogICAgICAgICAgInBhcnNlZCI6IDYwNjkKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICJuYW1lIjogImh0dHBzOi8vcmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbS9URy1Ud2lsaWdodC9BV0F2ZW51ZS1BZHMtUnVsZS9tYWluL0ZpbHRlcnMvQVdBdmVudWUtQWRzLVJ1bGUtU3VyZ2UubGlzdCIsCiAgICAgICAgICAiZm9ybWF0IjogImhvc3QiLAogICAgICAgICAgIm9rIjogdHJ1ZSwKICAgICAgICAgICJsaW5lcyI6IDk3MiwKICAgICAgICAgICJwYXJzZWQiOiA5NzIKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICJuYW1lIjogImh0dHBzOi8vcmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbS9iYWRtb2pyLzFIb3N0cy9tYXN0ZXIvTGl0ZS9kb21haW5zLnR4dCIsCiAgICAgICAgICAiZm9ybWF0IjogImhvc3QiLAogICAgICAgICAgIm9rIjogdHJ1ZSwKICAgICAgICAgICJsaW5lcyI6IDIwMjk1NCwKICAgICAgICAgICJwYXJzZWQiOiAyMDI5NTMsCiAgICAgICAgICAic2tpcHBlZCI6IHsKICAgICAgICAgICAgInVua25vd24iOiAxCiAgICAgICAgICB9CiAgICAgICAgfSwKICAgICAgICB7CiAgICAgICAgICAibmFtZSI6ICJodHRwczovL3Jhdy5naXRodWJ1c2VyY29udGVudC5jb20vZm16MjAwL3dvb2xfc2NyaXB0cy9tYWluL1F1YW50dW11bHRYL2ZpbHRlci9maWx0ZXIubGlzdCIsCiAgICAgICAgICAiZm9ybWF0IjogInF4IiwKICAgICAgICAgICJvayI6IHRydWUsCiAgICAgICAgICAibGluZXMiOiAyNzYxLAogICAgICAgICAgInBhcnNlZCI6IDI3NTMsCiAgICAgICAgICAic2tpcHBlZCI6IHsKICAgICAgICAgICAgInVua25vd24iOiA4CiAgICAgICAgICB9CiAgICAgICAgfQogICAgICBdCiAgICB9LAogICAgewogICAgICAidGFnIjogInJlamVjdC1pcCIsCiAgICAgICJydWxlcyI6IDEsCiAgICAgICJ2YWx1ZXMiOiAyOSwKICAgICAgInNraXBwZWQiOiB7CiAgICAgICAgInVua25vd24iOiA5CiAgICAgIH0sCiAgICAgICJ0ZXh0X3JlbW92ZWQiOiA0NzA5MiwKICAgICAgInNlbWFudGljX3JlbW92ZWQiOiAxMDg5MzIsCiAgICAgICJleGNsdWRlX3JlbW92ZWQiOiAwLAogICAgICAic291cmNlX2J5dGVzIjogOTI5LAogICAgICAiYmluYXJ5X2J5dGVzIjogMjAzLAogICAgICAic291cmNlcyI6IFsKICAgICAgICB7CiAgICAgICAgICAibmFtZSI6ICJodHRwczovL3Jhdy5naXRodWJ1c2VyY29udGVudC5jb20vZ2Vla2RhZGEvc3VyZ2UtbGlzdC9tYXN0ZXIvZG9tYWluLXNldC9kbnMtZmlsdGVyLnR4dCIsCiAgICAgICAgICAiZm9ybWF0IjogImhvc3QiLAogICAgICAgICAgIm9rIjogdHJ1ZSwKICAgICAgICAgICJsaW5lcyI6IDE4NDg0NCwKICAgICAgICAgICJwYXJzZWQiOiAxODQ4NDQKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICJuYW1lIjogImh0dHBzOi8vcmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbS9nZWVrZGFkYS9zdXJnZS1saXN0L21hc3Rlci9kb21haW4tc2V0L2NoaW5lc2UtZmlsdGVyLnR4dCIsCiAgICAgICAgICAiZm9ybWF0IjogImhvc3QiLAogICAgICAgICAgIm9rIjogdHJ1ZSwKICAgICAgICAgICJsaW5lcyI6IDYwNjksCiAgICAgICAgICAicGFyc2VkIjogNjA2OQogICAgICAgIH0sCiAgICAgICAgewogICAgICAgICAgIm5hbWUiOiAiaHR0cHM6Ly9yYXcuZ2l0aHVidXNlcmNvbnRlbnQuY29tL1RHLVR3aWxpZ2h0L0FXQXZlbnVlLUFkcy1SdWxlL21haW4vRmlsdGVycy9BV0F2ZW51ZS1BZHMtUnVsZS1TdXJnZS5saXN0IiwKICAgICAgICAgICJmb3JtYXQiOiAiaG9zdCIsCiAgICAgICAgICAib2siOiB0cnVlLAogICAgICAgICAgImxpbmVzIjogOTcyLAogICAgICAgICAgInBhcnNlZCI6IDk3MgogICAgICAgIH0sCiAgICAgICAgewogICAgICAgICAgIm5hbWUiOiAiaHR0cHM6Ly9yYXcuZ2l0aHVidXNlcmNvbnRlbnQuY29tL2JhZG1vanIvMUhvc3RzL21hc3Rlci9MaXRlL2RvbWFpbnMudHh0IiwKICAgICAgICAgICJmb3JtYXQiOiAiaG9zdCIsCiAgICAgICAgICAib2siOiB0cnVlLAogICAgICAgICAgImxpbmVzIjogMjAyOTU0LAogICAgICAgICAgInBhcnNlZCI6IDIwMjk1MywKICAgICAgICAgICJza2lwcGVkIjogewogICAgICAgICAgICAidW5rbm93biI6IDEKICAgICAgICAgIH0KICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICJuYW1lIjogImh0dHBzOi8vcmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbS9mbXoyMDAvd29vbF9zY3JpcHRzL21haW4vUXVhbnR1bXVsdFgvZmlsdGVyL2ZpbHRlci5saXN0IiwKICAgICAgICAgICJmb3JtYXQiOiAicXgiLAogICAgICAgICAgIm9rIjogdHJ1ZSwKICAgICAgICAgICJsaW5lcyI6IDI3NjEsCiAgICAgICAgICAicGFyc2VkIjogMjc1MywKICAgICAgICAgICJza2lwcGVkIjogewogICAgICAgICAgICAidW5rbm93biI6IDgKICAgICAgICAgIH0KICAgICAgICB9CiAgICAgIF0KICAgIH0KICBdCn0=-->