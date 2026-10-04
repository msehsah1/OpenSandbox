# 29 — Egress FQDN / wildcard match

**Language: Go.** Egress allow/deny rules are FQDNs or `*.example.com`
wildcards (OSEP-0001). Matching must be label-aware: `notexample.com` must
not match `*.example.com`. A wildcard also does **not** match the bare
suffix (`example.com` vs `*.example.com`) — same as
`components/egress/pkg/policy/policy.go`.

## Task

```go
func MatchDomain(host, rule string) bool
func Decide(host string, deny, allow []string) string // "deny", "allow", or "default"
```

`Decide` applies deny first, then allow (deny-wins). Host compare is
case-insensitive. A rule `*` matches everything.

## Verify

```bash
cd onboarding/exercises
go test ./29-egress-domain-match/solution -count=1
```
