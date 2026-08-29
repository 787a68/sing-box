# Rule Set Changes

- generated: 2026-08-29 12:39 UTC
- sing-box pinned: v1.14.0-beta.14

## AI evaluation

本期要点：本次更新以拒绝类规则为主，拒绝集发生了大规模的重构和行数变动；全局域名组有小幅增加，直连与全局 IP 基本保持稳定；上游源总体可用，构建完成且状态为 OK。

质量与风险：语义去重在拒绝集中仍然占比偏高但较上次有所下降，文本去重略有上升——提示去重策略影响了大量条目；有一个源出现少量跳过/未完全解析，需要关注，但未见高跳率或上游离线；总体没有直接证据显示误删，但语义去重带来的覆盖变化值得注意。

建议：1) 抽样审查被语义去重移除或合并的拒绝条目（优先常见或可疑域名），确认是否产生误杀并调整去重阈值；2) 跟进出现跳过/未解析的源（filter.list），修复源端或解析器并验证源文件与构建产物的一致性。

## Summary

total: +1885800 -1884669

## Per File

```

direct.json: +0 -0
global-domain.json: +7 -0
global-ip.json: +0 -0
reject-domain.json: +2738 -2555
reject-ip.json: +0 -0
direct.srs: +0 -0
global-domain.srs: +28488 -28471
global-ip.srs: +0 -0
reject-domain.srs: +1854567 -1853643
reject-ip.srs: +0 -0
```

## Rule Sets

| tag | rules | values | text-rm | sem-rm | excl-rm | skipped | src bytes | bin bytes | sources |
|---|---|---|---|---|---|---|---|---|---|
| direct | 1 | 24 | 0 | 0 | 0 |  | 847 | 333 | Direct.list(clash,ok,24/24) |
| global-domain | 1 | 4194 | 356 | 312 | 19 |  | 102849 | 28488 | Proxy.list(clash,ok,513/513); gfw.txt(host,ok,4386/4386) |
| global-ip | 1 | 18 | 356 | 312 | 19 |  | 563 | 155 | Proxy.list(clash,ok,513/513); gfw.txt(host,ok,4386/4386) |
| reject-domain | 1 | 222799 | 59515 | 109740 | 0 | unknown:9 | 7076067 | 1854567 | dns-filter.txt(host,ok,177738/177738); chinese-filter.txt(host,ok,6267/6267); AWAvenue-Ads-Rule-Surge.list(host,ok,897/897); domains.txt(host,ok,204414/204415); filter.list(qx,ok,2753/2761) |
| reject-ip | 1 | 29 | 59515 | 109740 | 0 | unknown:9 | 929 | 203 | dns-filter.txt(host,ok,177738/177738); chinese-filter.txt(host,ok,6267/6267); AWAvenue-Ads-Rule-Surge.list(host,ok,897/897); domains.txt(host,ok,204414/204415); filter.list(qx,ok,2753/2761) |

## Upstream Activity & Quality

### Per Rule Set (new vs previous)

| tag | values | text-rm | sem-rm | excl-rm | src bytes | bin bytes |
|---|---|---|---|---|---|---|
| direct | 24 (+0) | 0 (+0) | 0 (+0) | 0 (+0) | 847 (+0) | 333 (+0) |
| global-domain | 4194 (+7) | 356 (+0) | 312 (+0) | 19 (+0) | 102849 (+156) | 28488 (+17) |
| global-ip | 18 (+0) | 356 (+0) | 312 (+0) | 19 (+0) | 563 (+0) | 155 (+0) |
| reject-domain | 222799 (+183) | 59515 (+318) | 109740 (-941) | 0 (+0) | 7076067 (+5608) | 1854567 (+924) |
| reject-ip | 29 (+0) | 59515 (+318) | 109740 (-941) | 0 (+0) | 929 (+0) | 203 (+0) |

### Upstream Sources (activity & quality)

| rule set | source | fmt | status | lines | parsed | skip rate |
|---|---|---|---|---|---|---|
| direct | Direct.list | clash | ok | 24 | 24 | 0.0% |
| global-domain | Proxy.list | clash | ok | 513 | 513 | 0.0% |
| global-domain | gfw.txt | host | ok | 4386 | 4386 | 0.0% |
| global-ip | Proxy.list | clash | ok | 513 | 513 | 0.0% |
| global-ip | gfw.txt | host | ok | 4386 | 4386 | 0.0% |
| reject-domain | dns-filter.txt | host | ok | 177738 | 177738 | 0.0% |
| reject-domain | chinese-filter.txt | host | ok | 6267 | 6267 | 0.0% |
| reject-domain | AWAvenue-Ads-Rule-Surge.list | host | ok | 897 | 897 | 0.0% |
| reject-domain | domains.txt | host | ok | 204415 | 204414 | 0.0% |
| reject-domain | filter.list | qx | ok | 2761 | 2753 | 0.3% |
| reject-ip | dns-filter.txt | host | ok | 177738 | 177738 | 0.0% |
| reject-ip | chinese-filter.txt | host | ok | 6267 | 6267 | 0.0% |
| reject-ip | AWAvenue-Ads-Rule-Surge.list | host | ok | 897 | 897 | 0.0% |
| reject-ip | domains.txt | host | ok | 204415 | 204414 | 0.0% |
| reject-ip | filter.list | qx | ok | 2761 | 2753 | 0.3% |

## sing-box update check

- pinned v1.14.0-beta.14 (beta), latest stable v1.13.20, beta ahead of stable (no action)

<!--report:ewogICJnZW5lcmF0ZWRfYXQiOiAiMjAyNi0wOC0yOVQxMjozODo1Ni45NTk2NDU5NjlaIiwKICAic2luZ19ib3hfdmVyc2lvbiI6ICJ2MS4xNC4wLWJldGEuMTQiLAogICJydWxlX3NldHMiOiBbCiAgICB7CiAgICAgICJ0YWciOiAiZGlyZWN0IiwKICAgICAgInJ1bGVzIjogMSwKICAgICAgInZhbHVlcyI6IDI0LAogICAgICAidGV4dF9yZW1vdmVkIjogMCwKICAgICAgInNlbWFudGljX3JlbW92ZWQiOiAwLAogICAgICAiZXhjbHVkZV9yZW1vdmVkIjogMCwKICAgICAgInNvdXJjZV9ieXRlcyI6IDg0NywKICAgICAgImJpbmFyeV9ieXRlcyI6IDMzMywKICAgICAgInNvdXJjZXMiOiBbCiAgICAgICAgewogICAgICAgICAgIm5hbWUiOiAiaHR0cHM6Ly9yYXcuZ2l0aHVidXNlcmNvbnRlbnQuY29tL0Nvbm5lcnNIdWEvUnVsZUdvL21hc3Rlci9TdXJnZS9SdWxlc2V0L0RpcmVjdC5saXN0IiwKICAgICAgICAgICJmb3JtYXQiOiAiY2xhc2giLAogICAgICAgICAgIm9rIjogdHJ1ZSwKICAgICAgICAgICJsaW5lcyI6IDI0LAogICAgICAgICAgInBhcnNlZCI6IDI0CiAgICAgICAgfQogICAgICBdCiAgICB9LAogICAgewogICAgICAidGFnIjogImdsb2JhbC1kb21haW4iLAogICAgICAicnVsZXMiOiAxLAogICAgICAidmFsdWVzIjogNDE5NCwKICAgICAgInRleHRfcmVtb3ZlZCI6IDM1NiwKICAgICAgInNlbWFudGljX3JlbW92ZWQiOiAzMTIsCiAgICAgICJleGNsdWRlX3JlbW92ZWQiOiAxOSwKICAgICAgInNvdXJjZV9ieXRlcyI6IDEwMjg0OSwKICAgICAgImJpbmFyeV9ieXRlcyI6IDI4NDg4LAogICAgICAic291cmNlcyI6IFsKICAgICAgICB7CiAgICAgICAgICAibmFtZSI6ICJodHRwczovL3Jhdy5naXRodWJ1c2VyY29udGVudC5jb20vQ29ubmVyc0h1YS9SdWxlR28vbWFzdGVyL1N1cmdlL1J1bGVzZXQvUHJveHkubGlzdCIsCiAgICAgICAgICAiZm9ybWF0IjogImNsYXNoIiwKICAgICAgICAgICJvayI6IHRydWUsCiAgICAgICAgICAibGluZXMiOiA1MTMsCiAgICAgICAgICAicGFyc2VkIjogNTEzCiAgICAgICAgfSwKICAgICAgICB7CiAgICAgICAgICAibmFtZSI6ICJodHRwczovL3Jhdy5naXRodWJ1c2VyY29udGVudC5jb20vTG95YWxzb2xkaWVyL3N1cmdlLXJ1bGVzL3JlbGVhc2UvZ2Z3LnR4dCIsCiAgICAgICAgICAiZm9ybWF0IjogImhvc3QiLAogICAgICAgICAgIm9rIjogdHJ1ZSwKICAgICAgICAgICJsaW5lcyI6IDQzODYsCiAgICAgICAgICAicGFyc2VkIjogNDM4NgogICAgICAgIH0KICAgICAgXQogICAgfSwKICAgIHsKICAgICAgInRhZyI6ICJnbG9iYWwtaXAiLAogICAgICAicnVsZXMiOiAxLAogICAgICAidmFsdWVzIjogMTgsCiAgICAgICJ0ZXh0X3JlbW92ZWQiOiAzNTYsCiAgICAgICJzZW1hbnRpY19yZW1vdmVkIjogMzEyLAogICAgICAiZXhjbHVkZV9yZW1vdmVkIjogMTksCiAgICAgICJzb3VyY2VfYnl0ZXMiOiA1NjMsCiAgICAgICJiaW5hcnlfYnl0ZXMiOiAxNTUsCiAgICAgICJzb3VyY2VzIjogWwogICAgICAgIHsKICAgICAgICAgICJuYW1lIjogImh0dHBzOi8vcmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbS9Db25uZXJzSHVhL1J1bGVHby9tYXN0ZXIvU3VyZ2UvUnVsZXNldC9Qcm94eS5saXN0IiwKICAgICAgICAgICJmb3JtYXQiOiAiY2xhc2giLAogICAgICAgICAgIm9rIjogdHJ1ZSwKICAgICAgICAgICJsaW5lcyI6IDUxMywKICAgICAgICAgICJwYXJzZWQiOiA1MTMKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICJuYW1lIjogImh0dHBzOi8vcmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbS9Mb3lhbHNvbGRpZXIvc3VyZ2UtcnVsZXMvcmVsZWFzZS9nZncudHh0IiwKICAgICAgICAgICJmb3JtYXQiOiAiaG9zdCIsCiAgICAgICAgICAib2siOiB0cnVlLAogICAgICAgICAgImxpbmVzIjogNDM4NiwKICAgICAgICAgICJwYXJzZWQiOiA0Mzg2CiAgICAgICAgfQogICAgICBdCiAgICB9LAogICAgewogICAgICAidGFnIjogInJlamVjdC1kb21haW4iLAogICAgICAicnVsZXMiOiAxLAogICAgICAidmFsdWVzIjogMjIyNzk5LAogICAgICAic2tpcHBlZCI6IHsKICAgICAgICAidW5rbm93biI6IDkKICAgICAgfSwKICAgICAgInRleHRfcmVtb3ZlZCI6IDU5NTE1LAogICAgICAic2VtYW50aWNfcmVtb3ZlZCI6IDEwOTc0MCwKICAgICAgImV4Y2x1ZGVfcmVtb3ZlZCI6IDAsCiAgICAgICJzb3VyY2VfYnl0ZXMiOiA3MDc2MDY3LAogICAgICAiYmluYXJ5X2J5dGVzIjogMTg1NDU2NywKICAgICAgInNvdXJjZXMiOiBbCiAgICAgICAgewogICAgICAgICAgIm5hbWUiOiAiaHR0cHM6Ly9yYXcuZ2l0aHVidXNlcmNvbnRlbnQuY29tL2dlZWtkYWRhL3N1cmdlLWxpc3QvbWFzdGVyL2RvbWFpbi1zZXQvZG5zLWZpbHRlci50eHQiLAogICAgICAgICAgImZvcm1hdCI6ICJob3N0IiwKICAgICAgICAgICJvayI6IHRydWUsCiAgICAgICAgICAibGluZXMiOiAxNzc3MzgsCiAgICAgICAgICAicGFyc2VkIjogMTc3NzM4CiAgICAgICAgfSwKICAgICAgICB7CiAgICAgICAgICAibmFtZSI6ICJodHRwczovL3Jhdy5naXRodWJ1c2VyY29udGVudC5jb20vZ2Vla2RhZGEvc3VyZ2UtbGlzdC9tYXN0ZXIvZG9tYWluLXNldC9jaGluZXNlLWZpbHRlci50eHQiLAogICAgICAgICAgImZvcm1hdCI6ICJob3N0IiwKICAgICAgICAgICJvayI6IHRydWUsCiAgICAgICAgICAibGluZXMiOiA2MjY3LAogICAgICAgICAgInBhcnNlZCI6IDYyNjcKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICJuYW1lIjogImh0dHBzOi8vcmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbS9URy1Ud2lsaWdodC9BV0F2ZW51ZS1BZHMtUnVsZS9tYWluL0ZpbHRlcnMvQVdBdmVudWUtQWRzLVJ1bGUtU3VyZ2UubGlzdCIsCiAgICAgICAgICAiZm9ybWF0IjogImhvc3QiLAogICAgICAgICAgIm9rIjogdHJ1ZSwKICAgICAgICAgICJsaW5lcyI6IDg5NywKICAgICAgICAgICJwYXJzZWQiOiA4OTcKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICJuYW1lIjogImh0dHBzOi8vcmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbS9iYWRtb2pyLzFIb3N0cy9tYXN0ZXIvTGl0ZS9kb21haW5zLnR4dCIsCiAgICAgICAgICAiZm9ybWF0IjogImhvc3QiLAogICAgICAgICAgIm9rIjogdHJ1ZSwKICAgICAgICAgICJsaW5lcyI6IDIwNDQxNSwKICAgICAgICAgICJwYXJzZWQiOiAyMDQ0MTQsCiAgICAgICAgICAic2tpcHBlZCI6IHsKICAgICAgICAgICAgInVua25vd24iOiAxCiAgICAgICAgICB9CiAgICAgICAgfSwKICAgICAgICB7CiAgICAgICAgICAibmFtZSI6ICJodHRwczovL3Jhdy5naXRodWJ1c2VyY29udGVudC5jb20vZm16MjAwL3dvb2xfc2NyaXB0cy9tYWluL1F1YW50dW11bHRYL2ZpbHRlci9maWx0ZXIubGlzdCIsCiAgICAgICAgICAiZm9ybWF0IjogInF4IiwKICAgICAgICAgICJvayI6IHRydWUsCiAgICAgICAgICAibGluZXMiOiAyNzYxLAogICAgICAgICAgInBhcnNlZCI6IDI3NTMsCiAgICAgICAgICAic2tpcHBlZCI6IHsKICAgICAgICAgICAgInVua25vd24iOiA4CiAgICAgICAgICB9CiAgICAgICAgfQogICAgICBdCiAgICB9LAogICAgewogICAgICAidGFnIjogInJlamVjdC1pcCIsCiAgICAgICJydWxlcyI6IDEsCiAgICAgICJ2YWx1ZXMiOiAyOSwKICAgICAgInNraXBwZWQiOiB7CiAgICAgICAgInVua25vd24iOiA5CiAgICAgIH0sCiAgICAgICJ0ZXh0X3JlbW92ZWQiOiA1OTUxNSwKICAgICAgInNlbWFudGljX3JlbW92ZWQiOiAxMDk3NDAsCiAgICAgICJleGNsdWRlX3JlbW92ZWQiOiAwLAogICAgICAic291cmNlX2J5dGVzIjogOTI5LAogICAgICAiYmluYXJ5X2J5dGVzIjogMjAzLAogICAgICAic291cmNlcyI6IFsKICAgICAgICB7CiAgICAgICAgICAibmFtZSI6ICJodHRwczovL3Jhdy5naXRodWJ1c2VyY29udGVudC5jb20vZ2Vla2RhZGEvc3VyZ2UtbGlzdC9tYXN0ZXIvZG9tYWluLXNldC9kbnMtZmlsdGVyLnR4dCIsCiAgICAgICAgICAiZm9ybWF0IjogImhvc3QiLAogICAgICAgICAgIm9rIjogdHJ1ZSwKICAgICAgICAgICJsaW5lcyI6IDE3NzczOCwKICAgICAgICAgICJwYXJzZWQiOiAxNzc3MzgKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICJuYW1lIjogImh0dHBzOi8vcmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbS9nZWVrZGFkYS9zdXJnZS1saXN0L21hc3Rlci9kb21haW4tc2V0L2NoaW5lc2UtZmlsdGVyLnR4dCIsCiAgICAgICAgICAiZm9ybWF0IjogImhvc3QiLAogICAgICAgICAgIm9rIjogdHJ1ZSwKICAgICAgICAgICJsaW5lcyI6IDYyNjcsCiAgICAgICAgICAicGFyc2VkIjogNjI2NwogICAgICAgIH0sCiAgICAgICAgewogICAgICAgICAgIm5hbWUiOiAiaHR0cHM6Ly9yYXcuZ2l0aHVidXNlcmNvbnRlbnQuY29tL1RHLVR3aWxpZ2h0L0FXQXZlbnVlLUFkcy1SdWxlL21haW4vRmlsdGVycy9BV0F2ZW51ZS1BZHMtUnVsZS1TdXJnZS5saXN0IiwKICAgICAgICAgICJmb3JtYXQiOiAiaG9zdCIsCiAgICAgICAgICAib2siOiB0cnVlLAogICAgICAgICAgImxpbmVzIjogODk3LAogICAgICAgICAgInBhcnNlZCI6IDg5NwogICAgICAgIH0sCiAgICAgICAgewogICAgICAgICAgIm5hbWUiOiAiaHR0cHM6Ly9yYXcuZ2l0aHVidXNlcmNvbnRlbnQuY29tL2JhZG1vanIvMUhvc3RzL21hc3Rlci9MaXRlL2RvbWFpbnMudHh0IiwKICAgICAgICAgICJmb3JtYXQiOiAiaG9zdCIsCiAgICAgICAgICAib2siOiB0cnVlLAogICAgICAgICAgImxpbmVzIjogMjA0NDE1LAogICAgICAgICAgInBhcnNlZCI6IDIwNDQxNCwKICAgICAgICAgICJza2lwcGVkIjogewogICAgICAgICAgICAidW5rbm93biI6IDEKICAgICAgICAgIH0KICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICJuYW1lIjogImh0dHBzOi8vcmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbS9mbXoyMDAvd29vbF9zY3JpcHRzL21haW4vUXVhbnR1bXVsdFgvZmlsdGVyL2ZpbHRlci5saXN0IiwKICAgICAgICAgICJmb3JtYXQiOiAicXgiLAogICAgICAgICAgIm9rIjogdHJ1ZSwKICAgICAgICAgICJsaW5lcyI6IDI3NjEsCiAgICAgICAgICAicGFyc2VkIjogMjc1MywKICAgICAgICAgICJza2lwcGVkIjogewogICAgICAgICAgICAidW5rbm93biI6IDgKICAgICAgICAgIH0KICAgICAgICB9CiAgICAgIF0KICAgIH0KICBdCn0=-->