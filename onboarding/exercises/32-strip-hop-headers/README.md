# 32 — Strip hop-by-hop and routing headers

**Language: Go.** Ingress is a reverse proxy. RFC 7230 hop-by-hop headers
and `OpenSandbox-Ingress-To` must not be forwarded to the sandbox workload.

## Objective

Contribute ingress proxy header hygiene. RFC 7230 hop-by-hop headers
and `OpenSandbox-Ingress-To` must not reach the sandbox workload —
that would leak the routing target. Copy, do not mutate, the request
header map. This drill is `components/ingress/pkg/proxy/header.go`,
required for any forward/WebSocket header PR.

## Task

```go
func StripForwardHeaders(h map[string][]string) map[string][]string
```

Return a **copy** with these keys removed (canonical or any case):
`Connection`, `Keep-Alive`, `Proxy-Authenticate`, `Proxy-Authorization`,
`TE`, `Trailer`, `Transfer-Encoding`, `Upgrade`, `Proxy-Connection`,
`OpenSandbox-Ingress-To`, `OPEN-SANDBOX-INGRESS`.

Preserve other headers including `Authorization` and `X-Request-ID`.

## Verify

```bash
cd onboarding/exercises
go test ./32-strip-hop-headers/solution -count=1
```
