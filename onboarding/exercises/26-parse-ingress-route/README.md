# 26 — Parse an ingress route

**Language: Go.** `components/ingress` routes by header
`OpenSandbox-Ingress-To: <sandbox-id>-<port>` or URI
`/<sandbox-id>/<port>/<path>`. Sandbox IDs themselves contain hyphens, so
the port is the **last** `-` segment (or the second URI segment).

## Objective

Contribute ingress routing without mis-splitting sandbox IDs. IDs
contain hyphens; the port is the **last** `-` segment (header) or the
second URI segment. A wrong split 404s or proxies to the wrong
workload. This drill is `parseHostRoute` / `parseURIRoute` in
`components/ingress` — required for any host-mode or URI-mode PR.

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
