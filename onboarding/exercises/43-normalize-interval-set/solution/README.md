# Solution steps

1. `netip.ParsePrefix`, else `ParseAddr` then `.Prefix(32|128)`, then `.Masked()`.
2. Dedupe on `prefix.String()`.
3. Strict supernet: same family, `super.Bits() < sub.Bits()`, and `super.Contains(sub.Addr())`.
4. `/32` and `/128` print as the address only.

See `components/egress/pkg/nftables/interval.go`.
