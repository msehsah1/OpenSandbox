# 45 — Classify a TCP connect error

**Language: Go.** Ingress network-readiness shadow mode classifies
`DialContext` errors into a small enum. HTTP/TLS outcomes are out of
scope — this is connect only. Use `errors.Is` / `errors.As` so wrapped
errors still match.

## Objective

Contribute ingress network-readiness shadow mode. Dial errors are
classified (`timeout`, `unreachable`, `dns_error`, …); HTTP/TLS are
out of scope. Misclassifying DNS as timeout changes the degraded
signal. This drill is `ClassifyConnectError` in
`components/ingress/pkg/proxy/connectivity`. Use `errors.Is` / `As`
on any new connect-error path.

## Task

```go
type Result string
const (
    ResultSuccess     Result = "success"
    ResultTimeout     Result = "timeout"
    ResultUnreachable Result = "unreachable"
    ResultRefused     Result = "refused"
    ResultDNS         Result = "dns_error"
    ResultCanceled    Result = "canceled"
    ResultOther       Result = "other"
)

func ClassifyConnectError(err error) Result
```

Order: syscall `ETIMEDOUT` → unreachable family → `ECONNREFUSED` →
`*net.DNSError` → `context.DeadlineExceeded` → `context.Canceled` →
`net.Error.Timeout()` → other. `nil` is success.

## Verify

```bash
cd onboarding/exercises
go test ./45-classify-connect-error/solution -count=1
```
