# Rule Set Changes

- generated: 2026-08-22 06:56 UTC
- sing-box pinned: v1.14.0-beta.14

## AI evaluation

## 本期评估

**本期要点**  
reject-domain 增长超两万条，驱动整体规模扩张。所有上游源健康无失败，全构建成功。增长主要来自现有源的内容更新，而非新源接入。

**质量与风险**  
无异常。上游全绿、skip rate 均低、所有源保持稳定。但 reject-domain 新增规则的语义去重率相对下降（总体 sem-rm 计数反减），暗示新增内容重复度或泛匹配度略高于历史基准——建议关注是否因源端更新策略变化引入了语义冗余。

**建议**  
1. 抽查 dns-filter.txt/domains.txt 等大源的近期新增条目，确认无泛匹配或过度覆盖。  
2. 观察下期 sem-rm 趋势，如持续下降可考虑调整去重参数或源顺序优化。

## Summary

total: +1914672 -1713051

## Per File

```

direct.json: +0 -0
global-domain.json: +6 -0
global-ip.json: +0 -0
reject-domain.json: +31861 -10714
reject-ip.json: +0 -0
direct.srs: +333 -336
global-domain.srs: +28471 -27294
global-ip.srs: +155 -159
reject-domain.srs: +1853643 -1674342
reject-ip.srs: +203 -206
```

## Rule Sets

| tag | rules | values | text-rm | sem-rm | excl-rm | skipped | src bytes | bin bytes | sources |
|---|---|---|---|---|---|---|---|---|---|
| direct | 1 | 24 | 0 | 0 | 0 |  | 847 | 333 | Direct.list(clash,ok,24/24) |
| global-domain | 1 | 4187 | 356 | 312 | 19 |  | 102693 | 28471 | Proxy.list(clash,ok,513/513); gfw.txt(host,ok,4379/4379) |
| global-ip | 1 | 18 | 356 | 312 | 19 |  | 563 | 155 | Proxy.list(clash,ok,513/513); gfw.txt(host,ok,4379/4379) |
| reject-domain | 1 | 222616 | 59197 | 110681 | 0 | unknown:9 | 7070459 | 1853643 | dns-filter.txt(host,ok,176930/176930); chinese-filter.txt(host,ok,6470/6470); AWAvenue-Ads-Rule-Surge.list(host,ok,897/897); domains.txt(host,ok,205459/205460); filter.list(qx,ok,2753/2761) |
| reject-ip | 1 | 29 | 59197 | 110681 | 0 | unknown:9 | 929 | 203 | dns-filter.txt(host,ok,176930/176930); chinese-filter.txt(host,ok,6470/6470); AWAvenue-Ads-Rule-Surge.list(host,ok,897/897); domains.txt(host,ok,205459/205460); filter.list(qx,ok,2753/2761) |

## Upstream Activity & Quality

### Per Rule Set (new vs previous)

| tag | values | text-rm | sem-rm | excl-rm | src bytes | bin bytes |
|---|---|---|---|---|---|---|
| direct | 24 (+0) | 0 (+0) | 0 (+0) | 0 (+0) | 847 (+0) | 333 (-3) |
| global-domain | 4187 (+6) | 356 (+0) | 312 (+0) | 19 (+0) | 102693 (+184) | 28471 (+1177) |
| global-ip | 18 (+0) | 356 (+0) | 312 (+0) | 19 (+0) | 563 (+0) | 155 (-4) |
| reject-domain | 222616 (+21147) | 59197 (+795) | 110681 (-305) | 0 (+0) | 7070459 (+710401) | 1853643 (+179301) |
| reject-ip | 29 (+0) | 59197 (+795) | 110681 (-305) | 0 (+0) | 929 (+0) | 203 (-3) |

### Upstream Sources (activity & quality)

| rule set | source | fmt | status | lines | parsed | skip rate |
|---|---|---|---|---|---|---|
| direct | Direct.list | clash | ok | 24 | 24 | 0.0% |
| global-domain | Proxy.list | clash | ok | 513 | 513 | 0.0% |
| global-domain | gfw.txt | host | ok | 4379 | 4379 | 0.0% |
| global-ip | Proxy.list | clash | ok | 513 | 513 | 0.0% |
| global-ip | gfw.txt | host | ok | 4379 | 4379 | 0.0% |
| reject-domain | dns-filter.txt | host | ok | 176930 | 176930 | 0.0% |
| reject-domain | chinese-filter.txt | host | ok | 6470 | 6470 | 0.0% |
| reject-domain | AWAvenue-Ads-Rule-Surge.list | host | ok | 897 | 897 | 0.0% |
| reject-domain | domains.txt | host | ok | 205460 | 205459 | 0.0% |
| reject-domain | filter.list | qx | ok | 2761 | 2753 | 0.3% |
| reject-ip | dns-filter.txt | host | ok | 176930 | 176930 | 0.0% |
| reject-ip | chinese-filter.txt | host | ok | 6470 | 6470 | 0.0% |
| reject-ip | AWAvenue-Ads-Rule-Surge.list | host | ok | 897 | 897 | 0.0% |
| reject-ip | domains.txt | host | ok | 205460 | 205459 | 0.0% |
| reject-ip | filter.list | qx | ok | 2761 | 2753 | 0.3% |

## sing-box update check

- pinned v1.14.0-beta.14 (beta), latest stable v1.13.19, beta ahead of stable (no action)

<!--report:ewogICJnZW5lcmF0ZWRfYXQiOiAiMjAyNi0wOC0yMlQwNjo1NjowNC4xNTE3MzA3M1oiLAogICJzaW5nX2JveF92ZXJzaW9uIjogInYxLjE0LjAtYmV0YS4xNCIsCiAgInJ1bGVfc2V0cyI6IFsKICAgIHsKICAgICAgInRhZyI6ICJkaXJlY3QiLAogICAgICAicnVsZXMiOiAxLAogICAgICAidmFsdWVzIjogMjQsCiAgICAgICJ0ZXh0X3JlbW92ZWQiOiAwLAogICAgICAic2VtYW50aWNfcmVtb3ZlZCI6IDAsCiAgICAgICJleGNsdWRlX3JlbW92ZWQiOiAwLAogICAgICAic291cmNlX2J5dGVzIjogODQ3LAogICAgICAiYmluYXJ5X2J5dGVzIjogMzMzLAogICAgICAic291cmNlcyI6IFsKICAgICAgICB7CiAgICAgICAgICAibmFtZSI6ICJodHRwczovL3Jhdy5naXRodWJ1c2VyY29udGVudC5jb20vQ29ubmVyc0h1YS9SdWxlR28vbWFzdGVyL1N1cmdlL1J1bGVzZXQvRGlyZWN0Lmxpc3QiLAogICAgICAgICAgImZvcm1hdCI6ICJjbGFzaCIsCiAgICAgICAgICAib2siOiB0cnVlLAogICAgICAgICAgImxpbmVzIjogMjQsCiAgICAgICAgICAicGFyc2VkIjogMjQKICAgICAgICB9CiAgICAgIF0KICAgIH0sCiAgICB7CiAgICAgICJ0YWciOiAiZ2xvYmFsLWRvbWFpbiIsCiAgICAgICJydWxlcyI6IDEsCiAgICAgICJ2YWx1ZXMiOiA0MTg3LAogICAgICAidGV4dF9yZW1vdmVkIjogMzU2LAogICAgICAic2VtYW50aWNfcmVtb3ZlZCI6IDMxMiwKICAgICAgImV4Y2x1ZGVfcmVtb3ZlZCI6IDE5LAogICAgICAic291cmNlX2J5dGVzIjogMTAyNjkzLAogICAgICAiYmluYXJ5X2J5dGVzIjogMjg0NzEsCiAgICAgICJzb3VyY2VzIjogWwogICAgICAgIHsKICAgICAgICAgICJuYW1lIjogImh0dHBzOi8vcmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbS9Db25uZXJzSHVhL1J1bGVHby9tYXN0ZXIvU3VyZ2UvUnVsZXNldC9Qcm94eS5saXN0IiwKICAgICAgICAgICJmb3JtYXQiOiAiY2xhc2giLAogICAgICAgICAgIm9rIjogdHJ1ZSwKICAgICAgICAgICJsaW5lcyI6IDUxMywKICAgICAgICAgICJwYXJzZWQiOiA1MTMKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICJuYW1lIjogImh0dHBzOi8vcmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbS9Mb3lhbHNvbGRpZXIvc3VyZ2UtcnVsZXMvcmVsZWFzZS9nZncudHh0IiwKICAgICAgICAgICJmb3JtYXQiOiAiaG9zdCIsCiAgICAgICAgICAib2siOiB0cnVlLAogICAgICAgICAgImxpbmVzIjogNDM3OSwKICAgICAgICAgICJwYXJzZWQiOiA0Mzc5CiAgICAgICAgfQogICAgICBdCiAgICB9LAogICAgewogICAgICAidGFnIjogImdsb2JhbC1pcCIsCiAgICAgICJydWxlcyI6IDEsCiAgICAgICJ2YWx1ZXMiOiAxOCwKICAgICAgInRleHRfcmVtb3ZlZCI6IDM1NiwKICAgICAgInNlbWFudGljX3JlbW92ZWQiOiAzMTIsCiAgICAgICJleGNsdWRlX3JlbW92ZWQiOiAxOSwKICAgICAgInNvdXJjZV9ieXRlcyI6IDU2MywKICAgICAgImJpbmFyeV9ieXRlcyI6IDE1NSwKICAgICAgInNvdXJjZXMiOiBbCiAgICAgICAgewogICAgICAgICAgIm5hbWUiOiAiaHR0cHM6Ly9yYXcuZ2l0aHVidXNlcmNvbnRlbnQuY29tL0Nvbm5lcnNIdWEvUnVsZUdvL21hc3Rlci9TdXJnZS9SdWxlc2V0L1Byb3h5Lmxpc3QiLAogICAgICAgICAgImZvcm1hdCI6ICJjbGFzaCIsCiAgICAgICAgICAib2siOiB0cnVlLAogICAgICAgICAgImxpbmVzIjogNTEzLAogICAgICAgICAgInBhcnNlZCI6IDUxMwogICAgICAgIH0sCiAgICAgICAgewogICAgICAgICAgIm5hbWUiOiAiaHR0cHM6Ly9yYXcuZ2l0aHVidXNlcmNvbnRlbnQuY29tL0xveWFsc29sZGllci9zdXJnZS1ydWxlcy9yZWxlYXNlL2dmdy50eHQiLAogICAgICAgICAgImZvcm1hdCI6ICJob3N0IiwKICAgICAgICAgICJvayI6IHRydWUsCiAgICAgICAgICAibGluZXMiOiA0Mzc5LAogICAgICAgICAgInBhcnNlZCI6IDQzNzkKICAgICAgICB9CiAgICAgIF0KICAgIH0sCiAgICB7CiAgICAgICJ0YWciOiAicmVqZWN0LWRvbWFpbiIsCiAgICAgICJydWxlcyI6IDEsCiAgICAgICJ2YWx1ZXMiOiAyMjI2MTYsCiAgICAgICJza2lwcGVkIjogewogICAgICAgICJ1bmtub3duIjogOQogICAgICB9LAogICAgICAidGV4dF9yZW1vdmVkIjogNTkxOTcsCiAgICAgICJzZW1hbnRpY19yZW1vdmVkIjogMTEwNjgxLAogICAgICAiZXhjbHVkZV9yZW1vdmVkIjogMCwKICAgICAgInNvdXJjZV9ieXRlcyI6IDcwNzA0NTksCiAgICAgICJiaW5hcnlfYnl0ZXMiOiAxODUzNjQzLAogICAgICAic291cmNlcyI6IFsKICAgICAgICB7CiAgICAgICAgICAibmFtZSI6ICJodHRwczovL3Jhdy5naXRodWJ1c2VyY29udGVudC5jb20vZ2Vla2RhZGEvc3VyZ2UtbGlzdC9tYXN0ZXIvZG9tYWluLXNldC9kbnMtZmlsdGVyLnR4dCIsCiAgICAgICAgICAiZm9ybWF0IjogImhvc3QiLAogICAgICAgICAgIm9rIjogdHJ1ZSwKICAgICAgICAgICJsaW5lcyI6IDE3NjkzMCwKICAgICAgICAgICJwYXJzZWQiOiAxNzY5MzAKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICJuYW1lIjogImh0dHBzOi8vcmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbS9nZWVrZGFkYS9zdXJnZS1saXN0L21hc3Rlci9kb21haW4tc2V0L2NoaW5lc2UtZmlsdGVyLnR4dCIsCiAgICAgICAgICAiZm9ybWF0IjogImhvc3QiLAogICAgICAgICAgIm9rIjogdHJ1ZSwKICAgICAgICAgICJsaW5lcyI6IDY0NzAsCiAgICAgICAgICAicGFyc2VkIjogNjQ3MAogICAgICAgIH0sCiAgICAgICAgewogICAgICAgICAgIm5hbWUiOiAiaHR0cHM6Ly9yYXcuZ2l0aHVidXNlcmNvbnRlbnQuY29tL1RHLVR3aWxpZ2h0L0FXQXZlbnVlLUFkcy1SdWxlL21haW4vRmlsdGVycy9BV0F2ZW51ZS1BZHMtUnVsZS1TdXJnZS5saXN0IiwKICAgICAgICAgICJmb3JtYXQiOiAiaG9zdCIsCiAgICAgICAgICAib2siOiB0cnVlLAogICAgICAgICAgImxpbmVzIjogODk3LAogICAgICAgICAgInBhcnNlZCI6IDg5NwogICAgICAgIH0sCiAgICAgICAgewogICAgICAgICAgIm5hbWUiOiAiaHR0cHM6Ly9yYXcuZ2l0aHVidXNlcmNvbnRlbnQuY29tL2JhZG1vanIvMUhvc3RzL21hc3Rlci9MaXRlL2RvbWFpbnMudHh0IiwKICAgICAgICAgICJmb3JtYXQiOiAiaG9zdCIsCiAgICAgICAgICAib2siOiB0cnVlLAogICAgICAgICAgImxpbmVzIjogMjA1NDYwLAogICAgICAgICAgInBhcnNlZCI6IDIwNTQ1OSwKICAgICAgICAgICJza2lwcGVkIjogewogICAgICAgICAgICAidW5rbm93biI6IDEKICAgICAgICAgIH0KICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICJuYW1lIjogImh0dHBzOi8vcmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbS9mbXoyMDAvd29vbF9zY3JpcHRzL21haW4vUXVhbnR1bXVsdFgvZmlsdGVyL2ZpbHRlci5saXN0IiwKICAgICAgICAgICJmb3JtYXQiOiAicXgiLAogICAgICAgICAgIm9rIjogdHJ1ZSwKICAgICAgICAgICJsaW5lcyI6IDI3NjEsCiAgICAgICAgICAicGFyc2VkIjogMjc1MywKICAgICAgICAgICJza2lwcGVkIjogewogICAgICAgICAgICAidW5rbm93biI6IDgKICAgICAgICAgIH0KICAgICAgICB9CiAgICAgIF0KICAgIH0sCiAgICB7CiAgICAgICJ0YWciOiAicmVqZWN0LWlwIiwKICAgICAgInJ1bGVzIjogMSwKICAgICAgInZhbHVlcyI6IDI5LAogICAgICAic2tpcHBlZCI6IHsKICAgICAgICAidW5rbm93biI6IDkKICAgICAgfSwKICAgICAgInRleHRfcmVtb3ZlZCI6IDU5MTk3LAogICAgICAic2VtYW50aWNfcmVtb3ZlZCI6IDExMDY4MSwKICAgICAgImV4Y2x1ZGVfcmVtb3ZlZCI6IDAsCiAgICAgICJzb3VyY2VfYnl0ZXMiOiA5MjksCiAgICAgICJiaW5hcnlfYnl0ZXMiOiAyMDMsCiAgICAgICJzb3VyY2VzIjogWwogICAgICAgIHsKICAgICAgICAgICJuYW1lIjogImh0dHBzOi8vcmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbS9nZWVrZGFkYS9zdXJnZS1saXN0L21hc3Rlci9kb21haW4tc2V0L2Rucy1maWx0ZXIudHh0IiwKICAgICAgICAgICJmb3JtYXQiOiAiaG9zdCIsCiAgICAgICAgICAib2siOiB0cnVlLAogICAgICAgICAgImxpbmVzIjogMTc2OTMwLAogICAgICAgICAgInBhcnNlZCI6IDE3NjkzMAogICAgICAgIH0sCiAgICAgICAgewogICAgICAgICAgIm5hbWUiOiAiaHR0cHM6Ly9yYXcuZ2l0aHVidXNlcmNvbnRlbnQuY29tL2dlZWtkYWRhL3N1cmdlLWxpc3QvbWFzdGVyL2RvbWFpbi1zZXQvY2hpbmVzZS1maWx0ZXIudHh0IiwKICAgICAgICAgICJmb3JtYXQiOiAiaG9zdCIsCiAgICAgICAgICAib2siOiB0cnVlLAogICAgICAgICAgImxpbmVzIjogNjQ3MCwKICAgICAgICAgICJwYXJzZWQiOiA2NDcwCiAgICAgICAgfSwKICAgICAgICB7CiAgICAgICAgICAibmFtZSI6ICJodHRwczovL3Jhdy5naXRodWJ1c2VyY29udGVudC5jb20vVEctVHdpbGlnaHQvQVdBdmVudWUtQWRzLVJ1bGUvbWFpbi9GaWx0ZXJzL0FXQXZlbnVlLUFkcy1SdWxlLVN1cmdlLmxpc3QiLAogICAgICAgICAgImZvcm1hdCI6ICJob3N0IiwKICAgICAgICAgICJvayI6IHRydWUsCiAgICAgICAgICAibGluZXMiOiA4OTcsCiAgICAgICAgICAicGFyc2VkIjogODk3CiAgICAgICAgfSwKICAgICAgICB7CiAgICAgICAgICAibmFtZSI6ICJodHRwczovL3Jhdy5naXRodWJ1c2VyY29udGVudC5jb20vYmFkbW9qci8xSG9zdHMvbWFzdGVyL0xpdGUvZG9tYWlucy50eHQiLAogICAgICAgICAgImZvcm1hdCI6ICJob3N0IiwKICAgICAgICAgICJvayI6IHRydWUsCiAgICAgICAgICAibGluZXMiOiAyMDU0NjAsCiAgICAgICAgICAicGFyc2VkIjogMjA1NDU5LAogICAgICAgICAgInNraXBwZWQiOiB7CiAgICAgICAgICAgICJ1bmtub3duIjogMQogICAgICAgICAgfQogICAgICAgIH0sCiAgICAgICAgewogICAgICAgICAgIm5hbWUiOiAiaHR0cHM6Ly9yYXcuZ2l0aHVidXNlcmNvbnRlbnQuY29tL2ZtejIwMC93b29sX3NjcmlwdHMvbWFpbi9RdWFudHVtdWx0WC9maWx0ZXIvZmlsdGVyLmxpc3QiLAogICAgICAgICAgImZvcm1hdCI6ICJxeCIsCiAgICAgICAgICAib2siOiB0cnVlLAogICAgICAgICAgImxpbmVzIjogMjc2MSwKICAgICAgICAgICJwYXJzZWQiOiAyNzUzLAogICAgICAgICAgInNraXBwZWQiOiB7CiAgICAgICAgICAgICJ1bmtub3duIjogOAogICAgICAgICAgfQogICAgICAgIH0KICAgICAgXQogICAgfQogIF0KfQ==-->