# 26 — Parse an ingress route

**Language: Go.** `components/ingress` routes by header
`OpenSandbox-Ingress-To: <sandbox-id>-<port>` or URI
`/<sandbox-id>/<port>/<path>`. Sandbox IDs themselves contain hyphens, so
the port is the **last** `-` segment (or the second URI segment).

## Task

```go
type Route struct {
    SandboxID string
    Port      int
    Path      string // remaining URI path, "/" if none; empty for header mode
}

func ParseIngressHeader(value string) (Route, error)
func ParseIngressURI(path string) (Route, error)
```

Reject missing port, non-numeric port, empty id. Tolerate a leading scheme on
hosts (`https://id-44772.example`).

## Verify

```bash
cd onboarding/exercises
go test ./26-parse-ingress-route/solution -count=1
```
