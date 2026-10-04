# Solution steps

1. Check `errors.Is` against `syscall` errnos before context errors.
2. `errors.As` into `*net.DNSError` next — a DNS timeout is DNS, not timeout.
3. Then context deadline / canceled, then `net.Error.Timeout()`.
4. Never stringify the error to classify it.

See `ClassifyConnectError` in
`components/ingress/pkg/proxy/connectivity/result.go`.
