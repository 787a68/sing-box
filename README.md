# Rule Set Changes

- generated: 2026-08-12 07:15 UTC
- sing-box pinned: v1.14.0-beta.14

## AI evaluation

(AI evaluation unavailable)

## Summary

total: +1819292 -1817680

## Per File

```

direct.json: +4 -4
global-domain.json: +463 -463
global-ip.json: +0 -0
reject-domain.json: +150586 -150377
reject-ip.json: +0 -0
direct.srs: +0 -0
global-domain.srs: +0 -0
global-ip.srs: +0 -0
reject-domain.srs: +1668239 -1666836
reject-ip.srs: +0 -0
```

## Rule Sets

| tag | rules | values | text-rm | sem-rm | excl-rm | skipped | src bytes | bin bytes | sources |
|---|---|---|---|---|---|---|---|---|---|
| direct | 1 | 24 | 0 | 0 | 0 |  | 847 | 336 | Direct.list(clash,ok,24/24) |
| global-domain | 1 | 4181 | 356 | 312 | 19 |  | 102509 | 27294 | Proxy.list(clash,ok,513/513); gfw.txt(host,ok,4373/4373) |
| global-ip | 1 | 18 | 356 | 312 | 19 |  | 563 | 159 | Proxy.list(clash,ok,513/513); gfw.txt(host,ok,4373/4373) |
| reject-domain | 1 | 200695 | 57801 | 111107 | 0 | unknown:9 | 6336785 | 1668239 | dns-filter.txt(host,ok,153927/153927); chinese-filter.txt(host,ok,6349/6349); AWAvenue-Ads-Rule-Surge.list(host,ok,898/898); domains.txt(host,ok,205691/205692); filter.list(qx,ok,2753/2761) |
| reject-ip | 1 | 29 | 57801 | 111107 | 0 | unknown:9 | 929 | 206 | dns-filter.txt(host,ok,153927/153927); chinese-filter.txt(host,ok,6349/6349); AWAvenue-Ads-Rule-Surge.list(host,ok,898/898); domains.txt(host,ok,205691/205692); filter.list(qx,ok,2753/2761) |

## Upstream Activity & Quality

### Per Rule Set (new vs previous)

| tag | values | text-rm | sem-rm | excl-rm | src bytes | bin bytes |
|---|---|---|---|---|---|---|
| direct | 24 (+0) | 0 (+0) | 0 (+0) | 0 (+0) | 847 (+0) | 336 (+0) |
| global-domain | 4181 (+0) | 356 (+0) | 312 (+0) | 19 (+0) | 102509 (+0) | 27294 (+0) |
| global-ip | 18 (+0) | 356 (+0) | 312 (+0) | 19 (+0) | 563 (+0) | 159 (+0) |
| reject-domain | 200695 (+213) | 57801 (+149) | 111107 (+16) | 0 (+0) | 6336785 (+6370) | 1668239 (+1403) |
| reject-ip | 29 (+0) | 57801 (+149) | 111107 (+16) | 0 (+0) | 929 (+0) | 206 (+0) |

### Upstream Sources (activity & quality)

| rule set | source | fmt | status | lines | parsed | skip rate |
|---|---|---|---|---|---|---|
| direct | Direct.list | clash | ok | 24 | 24 | 0.0% |
| global-domain | Proxy.list | clash | ok | 513 | 513 | 0.0% |
| global-domain | gfw.txt | host | ok | 4373 | 4373 | 0.0% |
| global-ip | Proxy.list | clash | ok | 513 | 513 | 0.0% |
| global-ip | gfw.txt | host | ok | 4373 | 4373 | 0.0% |
| reject-domain | dns-filter.txt | host | ok | 153927 | 153927 | 0.0% |
| reject-domain | chinese-filter.txt | host | ok | 6349 | 6349 | 0.0% |
| reject-domain | AWAvenue-Ads-Rule-Surge.list | host | ok | 898 | 898 | 0.0% |
| reject-domain | domains.txt | host | ok | 205692 | 205691 | 0.0% |
| reject-domain | filter.list | qx | ok | 2761 | 2753 | 0.3% |
| reject-ip | dns-filter.txt | host | ok | 153927 | 153927 | 0.0% |
| reject-ip | chinese-filter.txt | host | ok | 6349 | 6349 | 0.0% |
| reject-ip | AWAvenue-Ads-Rule-Surge.list | host | ok | 898 | 898 | 0.0% |
| reject-ip | domains.txt | host | ok | 205692 | 205691 | 0.0% |
| reject-ip | filter.list | qx | ok | 2761 | 2753 | 0.3% |

## sing-box update check

- pinned v1.14.0-beta.14 (beta), latest stable v1.13.18, beta ahead of stable (no action)

<!--report:ewogICJnZW5lcmF0ZWRfYXQiOiAiMjAyNi0wOC0xMlQwNzoxNToyNy44NTM4MDE4NzZaIiwKICAic2luZ19ib3hfdmVyc2lvbiI6ICJ2MS4xNC4wLWJldGEuMTQiLAogICJydWxlX3NldHMiOiBbCiAgICB7CiAgICAgICJ0YWciOiAiZGlyZWN0IiwKICAgICAgInJ1bGVzIjogMSwKICAgICAgInZhbHVlcyI6IDI0LAogICAgICAidGV4dF9yZW1vdmVkIjogMCwKICAgICAgInNlbWFudGljX3JlbW92ZWQiOiAwLAogICAgICAiZXhjbHVkZV9yZW1vdmVkIjogMCwKICAgICAgInNvdXJjZV9ieXRlcyI6IDg0NywKICAgICAgImJpbmFyeV9ieXRlcyI6IDMzNiwKICAgICAgInNvdXJjZXMiOiBbCiAgICAgICAgewogICAgICAgICAgIm5hbWUiOiAiaHR0cHM6Ly9yYXcuZ2l0aHVidXNlcmNvbnRlbnQuY29tL0Nvbm5lcnNIdWEvUnVsZUdvL21hc3Rlci9TdXJnZS9SdWxlc2V0L0RpcmVjdC5saXN0IiwKICAgICAgICAgICJmb3JtYXQiOiAiY2xhc2giLAogICAgICAgICAgIm9rIjogdHJ1ZSwKICAgICAgICAgICJsaW5lcyI6IDI0LAogICAgICAgICAgInBhcnNlZCI6IDI0CiAgICAgICAgfQogICAgICBdCiAgICB9LAogICAgewogICAgICAidGFnIjogImdsb2JhbC1kb21haW4iLAogICAgICAicnVsZXMiOiAxLAogICAgICAidmFsdWVzIjogNDE4MSwKICAgICAgInRleHRfcmVtb3ZlZCI6IDM1NiwKICAgICAgInNlbWFudGljX3JlbW92ZWQiOiAzMTIsCiAgICAgICJleGNsdWRlX3JlbW92ZWQiOiAxOSwKICAgICAgInNvdXJjZV9ieXRlcyI6IDEwMjUwOSwKICAgICAgImJpbmFyeV9ieXRlcyI6IDI3Mjk0LAogICAgICAic291cmNlcyI6IFsKICAgICAgICB7CiAgICAgICAgICAibmFtZSI6ICJodHRwczovL3Jhdy5naXRodWJ1c2VyY29udGVudC5jb20vQ29ubmVyc0h1YS9SdWxlR28vbWFzdGVyL1N1cmdlL1J1bGVzZXQvUHJveHkubGlzdCIsCiAgICAgICAgICAiZm9ybWF0IjogImNsYXNoIiwKICAgICAgICAgICJvayI6IHRydWUsCiAgICAgICAgICAibGluZXMiOiA1MTMsCiAgICAgICAgICAicGFyc2VkIjogNTEzCiAgICAgICAgfSwKICAgICAgICB7CiAgICAgICAgICAibmFtZSI6ICJodHRwczovL3Jhdy5naXRodWJ1c2VyY29udGVudC5jb20vTG95YWxzb2xkaWVyL3N1cmdlLXJ1bGVzL3JlbGVhc2UvZ2Z3LnR4dCIsCiAgICAgICAgICAiZm9ybWF0IjogImhvc3QiLAogICAgICAgICAgIm9rIjogdHJ1ZSwKICAgICAgICAgICJsaW5lcyI6IDQzNzMsCiAgICAgICAgICAicGFyc2VkIjogNDM3MwogICAgICAgIH0KICAgICAgXQogICAgfSwKICAgIHsKICAgICAgInRhZyI6ICJnbG9iYWwtaXAiLAogICAgICAicnVsZXMiOiAxLAogICAgICAidmFsdWVzIjogMTgsCiAgICAgICJ0ZXh0X3JlbW92ZWQiOiAzNTYsCiAgICAgICJzZW1hbnRpY19yZW1vdmVkIjogMzEyLAogICAgICAiZXhjbHVkZV9yZW1vdmVkIjogMTksCiAgICAgICJzb3VyY2VfYnl0ZXMiOiA1NjMsCiAgICAgICJiaW5hcnlfYnl0ZXMiOiAxNTksCiAgICAgICJzb3VyY2VzIjogWwogICAgICAgIHsKICAgICAgICAgICJuYW1lIjogImh0dHBzOi8vcmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbS9Db25uZXJzSHVhL1J1bGVHby9tYXN0ZXIvU3VyZ2UvUnVsZXNldC9Qcm94eS5saXN0IiwKICAgICAgICAgICJmb3JtYXQiOiAiY2xhc2giLAogICAgICAgICAgIm9rIjogdHJ1ZSwKICAgICAgICAgICJsaW5lcyI6IDUxMywKICAgICAgICAgICJwYXJzZWQiOiA1MTMKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICJuYW1lIjogImh0dHBzOi8vcmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbS9Mb3lhbHNvbGRpZXIvc3VyZ2UtcnVsZXMvcmVsZWFzZS9nZncudHh0IiwKICAgICAgICAgICJmb3JtYXQiOiAiaG9zdCIsCiAgICAgICAgICAib2siOiB0cnVlLAogICAgICAgICAgImxpbmVzIjogNDM3MywKICAgICAgICAgICJwYXJzZWQiOiA0MzczCiAgICAgICAgfQogICAgICBdCiAgICB9LAogICAgewogICAgICAidGFnIjogInJlamVjdC1kb21haW4iLAogICAgICAicnVsZXMiOiAxLAogICAgICAidmFsdWVzIjogMjAwNjk1LAogICAgICAic2tpcHBlZCI6IHsKICAgICAgICAidW5rbm93biI6IDkKICAgICAgfSwKICAgICAgInRleHRfcmVtb3ZlZCI6IDU3ODAxLAogICAgICAic2VtYW50aWNfcmVtb3ZlZCI6IDExMTEwNywKICAgICAgImV4Y2x1ZGVfcmVtb3ZlZCI6IDAsCiAgICAgICJzb3VyY2VfYnl0ZXMiOiA2MzM2Nzg1LAogICAgICAiYmluYXJ5X2J5dGVzIjogMTY2ODIzOSwKICAgICAgInNvdXJjZXMiOiBbCiAgICAgICAgewogICAgICAgICAgIm5hbWUiOiAiaHR0cHM6Ly9yYXcuZ2l0aHVidXNlcmNvbnRlbnQuY29tL2dlZWtkYWRhL3N1cmdlLWxpc3QvbWFzdGVyL2RvbWFpbi1zZXQvZG5zLWZpbHRlci50eHQiLAogICAgICAgICAgImZvcm1hdCI6ICJob3N0IiwKICAgICAgICAgICJvayI6IHRydWUsCiAgICAgICAgICAibGluZXMiOiAxNTM5MjcsCiAgICAgICAgICAicGFyc2VkIjogMTUzOTI3CiAgICAgICAgfSwKICAgICAgICB7CiAgICAgICAgICAibmFtZSI6ICJodHRwczovL3Jhdy5naXRodWJ1c2VyY29udGVudC5jb20vZ2Vla2RhZGEvc3VyZ2UtbGlzdC9tYXN0ZXIvZG9tYWluLXNldC9jaGluZXNlLWZpbHRlci50eHQiLAogICAgICAgICAgImZvcm1hdCI6ICJob3N0IiwKICAgICAgICAgICJvayI6IHRydWUsCiAgICAgICAgICAibGluZXMiOiA2MzQ5LAogICAgICAgICAgInBhcnNlZCI6IDYzNDkKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICJuYW1lIjogImh0dHBzOi8vcmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbS9URy1Ud2lsaWdodC9BV0F2ZW51ZS1BZHMtUnVsZS9tYWluL0ZpbHRlcnMvQVdBdmVudWUtQWRzLVJ1bGUtU3VyZ2UubGlzdCIsCiAgICAgICAgICAiZm9ybWF0IjogImhvc3QiLAogICAgICAgICAgIm9rIjogdHJ1ZSwKICAgICAgICAgICJsaW5lcyI6IDg5OCwKICAgICAgICAgICJwYXJzZWQiOiA4OTgKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICJuYW1lIjogImh0dHBzOi8vcmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbS9iYWRtb2pyLzFIb3N0cy9tYXN0ZXIvTGl0ZS9kb21haW5zLnR4dCIsCiAgICAgICAgICAiZm9ybWF0IjogImhvc3QiLAogICAgICAgICAgIm9rIjogdHJ1ZSwKICAgICAgICAgICJsaW5lcyI6IDIwNTY5MiwKICAgICAgICAgICJwYXJzZWQiOiAyMDU2OTEsCiAgICAgICAgICAic2tpcHBlZCI6IHsKICAgICAgICAgICAgInVua25vd24iOiAxCiAgICAgICAgICB9CiAgICAgICAgfSwKICAgICAgICB7CiAgICAgICAgICAibmFtZSI6ICJodHRwczovL3Jhdy5naXRodWJ1c2VyY29udGVudC5jb20vZm16MjAwL3dvb2xfc2NyaXB0cy9tYWluL1F1YW50dW11bHRYL2ZpbHRlci9maWx0ZXIubGlzdCIsCiAgICAgICAgICAiZm9ybWF0IjogInF4IiwKICAgICAgICAgICJvayI6IHRydWUsCiAgICAgICAgICAibGluZXMiOiAyNzYxLAogICAgICAgICAgInBhcnNlZCI6IDI3NTMsCiAgICAgICAgICAic2tpcHBlZCI6IHsKICAgICAgICAgICAgInVua25vd24iOiA4CiAgICAgICAgICB9CiAgICAgICAgfQogICAgICBdCiAgICB9LAogICAgewogICAgICAidGFnIjogInJlamVjdC1pcCIsCiAgICAgICJydWxlcyI6IDEsCiAgICAgICJ2YWx1ZXMiOiAyOSwKICAgICAgInNraXBwZWQiOiB7CiAgICAgICAgInVua25vd24iOiA5CiAgICAgIH0sCiAgICAgICJ0ZXh0X3JlbW92ZWQiOiA1NzgwMSwKICAgICAgInNlbWFudGljX3JlbW92ZWQiOiAxMTExMDcsCiAgICAgICJleGNsdWRlX3JlbW92ZWQiOiAwLAogICAgICAic291cmNlX2J5dGVzIjogOTI5LAogICAgICAiYmluYXJ5X2J5dGVzIjogMjA2LAogICAgICAic291cmNlcyI6IFsKICAgICAgICB7CiAgICAgICAgICAibmFtZSI6ICJodHRwczovL3Jhdy5naXRodWJ1c2VyY29udGVudC5jb20vZ2Vla2RhZGEvc3VyZ2UtbGlzdC9tYXN0ZXIvZG9tYWluLXNldC9kbnMtZmlsdGVyLnR4dCIsCiAgICAgICAgICAiZm9ybWF0IjogImhvc3QiLAogICAgICAgICAgIm9rIjogdHJ1ZSwKICAgICAgICAgICJsaW5lcyI6IDE1MzkyNywKICAgICAgICAgICJwYXJzZWQiOiAxNTM5MjcKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICJuYW1lIjogImh0dHBzOi8vcmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbS9nZWVrZGFkYS9zdXJnZS1saXN0L21hc3Rlci9kb21haW4tc2V0L2NoaW5lc2UtZmlsdGVyLnR4dCIsCiAgICAgICAgICAiZm9ybWF0IjogImhvc3QiLAogICAgICAgICAgIm9rIjogdHJ1ZSwKICAgICAgICAgICJsaW5lcyI6IDYzNDksCiAgICAgICAgICAicGFyc2VkIjogNjM0OQogICAgICAgIH0sCiAgICAgICAgewogICAgICAgICAgIm5hbWUiOiAiaHR0cHM6Ly9yYXcuZ2l0aHVidXNlcmNvbnRlbnQuY29tL1RHLVR3aWxpZ2h0L0FXQXZlbnVlLUFkcy1SdWxlL21haW4vRmlsdGVycy9BV0F2ZW51ZS1BZHMtUnVsZS1TdXJnZS5saXN0IiwKICAgICAgICAgICJmb3JtYXQiOiAiaG9zdCIsCiAgICAgICAgICAib2siOiB0cnVlLAogICAgICAgICAgImxpbmVzIjogODk4LAogICAgICAgICAgInBhcnNlZCI6IDg5OAogICAgICAgIH0sCiAgICAgICAgewogICAgICAgICAgIm5hbWUiOiAiaHR0cHM6Ly9yYXcuZ2l0aHVidXNlcmNvbnRlbnQuY29tL2JhZG1vanIvMUhvc3RzL21hc3Rlci9MaXRlL2RvbWFpbnMudHh0IiwKICAgICAgICAgICJmb3JtYXQiOiAiaG9zdCIsCiAgICAgICAgICAib2siOiB0cnVlLAogICAgICAgICAgImxpbmVzIjogMjA1NjkyLAogICAgICAgICAgInBhcnNlZCI6IDIwNTY5MSwKICAgICAgICAgICJza2lwcGVkIjogewogICAgICAgICAgICAidW5rbm93biI6IDEKICAgICAgICAgIH0KICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICJuYW1lIjogImh0dHBzOi8vcmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbS9mbXoyMDAvd29vbF9zY3JpcHRzL21haW4vUXVhbnR1bXVsdFgvZmlsdGVyL2ZpbHRlci5saXN0IiwKICAgICAgICAgICJmb3JtYXQiOiAicXgiLAogICAgICAgICAgIm9rIjogdHJ1ZSwKICAgICAgICAgICJsaW5lcyI6IDI3NjEsCiAgICAgICAgICAicGFyc2VkIjogMjc1MywKICAgICAgICAgICJza2lwcGVkIjogewogICAgICAgICAgICAidW5rbm93biI6IDgKICAgICAgICAgIH0KICAgICAgICB9CiAgICAgIF0KICAgIH0KICBdCn0=-->