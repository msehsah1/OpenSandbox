# Solution steps

1. Try `netip.ParseAddr` then `netip.ParsePrefix` on the rule.
2. Parse `addr` with `netip.ParseAddr`. Compare with `==` or `prefix.Contains(ip)`.
3. Do not use `strings.HasPrefix` on address text.
4. `ClassifyTarget` is the same dispatch `normalizePolicy` uses.

See `normalizePolicy` / `StaticIPSets` in `components/egress/pkg/policy/policy.go`.
