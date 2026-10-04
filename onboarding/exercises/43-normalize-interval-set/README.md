# 43 — Normalize an nft interval set

**Language: Go.** nftables sets with `flags interval` reject a host or
smaller CIDR that sits strictly inside another listed CIDR
("conflicting intervals specified"). Egress drops the redundant
entries before `nft add element`.

## Task

```go
func NormalizeIntervalSet(elems []string) ([]string, error)
```

- Parse each element as a CIDR, or as an IP (`/32` or `/128`)
- Mask prefixes (`10.1.2.3/16` → `10.1.0.0/16`)
- Drop duplicates (keep first)
- Drop an entry that is a **strict** subnet of another kept entry
- Format hosts as bare IPs; keep CIDR notation otherwise
- Invalid element → error

## Verify

```bash
cd onboarding/exercises
go test ./43-normalize-interval-set/solution -count=1
```
