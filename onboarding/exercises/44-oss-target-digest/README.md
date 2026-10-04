# 44 — Canonical OSS endpoint and target digest

**Language: Go.** Node-agent stream identities are `sha256:` hex of a
domain string plus length-prefixed parts. OSS endpoints must be an
HTTPS **origin** (no userinfo, query, fragment, or path other than `/`).
Host is lowercased; port 443 is omitted.

## Task

```go
func CanonicalOSSEndpoint(raw string) (string, error)
func TargetDigest(domain string, parts ...string) string
```

`TargetDigest` writes `domain`, then for each part `fmt.Fprintf(h, "%d:%s", len(part), part)`
(use `len([]byte(part))`), and returns `"sha256:" + hex`.

IPv6 hosts stay bracketed (`https://[2001:db8::1]`).

## Verify

```bash
cd onboarding/exercises
go test ./44-oss-target-digest/solution -count=1
```
