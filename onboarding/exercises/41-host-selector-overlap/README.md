# 41 — Host selector match and overlap

**Language: Go.** OSEP-0023 host selectors are an exact FQDN or a
leftmost wildcard (`*.example.com`). Wildcards match proper subdomains
only (not the apex). Overlap asks whether any valid hostname belongs to
both selectors — nested wildcards overlap.

This drill uses the **canonical ASCII** form (no IDNA). Reject IP
literals, single-label names, empty labels, and uppercase.

## Objective

Contribute OSEP-0023 host selectors for credential-vault / TLS
policy. Wildcards match proper subdomains only; overlap decides
whether two rules can claim the same hostname. A bad overlap lets
conflicting credentials apply. This drill is
`components/egress/pkg/hostselector` — the algebra, not IDNA.

## Task

```go
type Selector struct { /* unexported fields */ }

func ParseCanonical(text string) (Selector, error)
func (s Selector) String() string
func (s Selector) Matches(host string) bool
func (s Selector) Overlaps(other Selector) bool
```

Zero-value selector matches nothing.

## Verify

```bash
cd onboarding/exercises
go test ./41-host-selector-overlap/solution -count=1
```
