# 42 — Match an IP against IP/CIDR rules

**Language: Go.** Egress policy targets are a domain, a single IP, or a
CIDR (`netip`). Matching must not treat `10.1.2.3` as a string prefix of
`10.1.2.0/24` — use `netip.Prefix.Contains`.

## Objective

Contribute egress IP/CIDR targets without string-prefix bugs.
`10.1.2.3` is not “in” `10.1.2.0/24` because the text starts with
`10.1.2`. Use `netip`. This drill is `normalizePolicy` /
`StaticIPSets` in `components/egress/pkg/policy`. Policy PRs that
match IPs with `HasPrefix` are security defects.

## Task

```go
func MatchIP(addr, rule string) bool
func ClassifyTarget(rule string) string // "ip", "cidr", "domain", or "invalid"
```

`MatchIP` is true when `addr` equals an IP rule or falls inside a CIDR
rule. Domain rules never match an address. Unparseable inputs are false /
`"invalid"`. IPv4-mapped IPv6 must not silently match IPv4 rules — parse
both sides with `netip` and compare as parsed.

## Verify

```bash
cd onboarding/exercises
go test ./42-match-ip-cidr/solution -count=1
```
