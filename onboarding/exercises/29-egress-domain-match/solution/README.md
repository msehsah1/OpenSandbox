# Solution steps

1. Lowercase both sides.
2. Exact match, or `*` rule, or `*.suffix` where host ends with `"."+suffix`.
   The bare suffix (`example.com` vs `*.example.com`) does not match.
3. Deny list first — a host on both lists is denied.
4. Do not use naive `strings.Contains`.

This is the policy core of `components/egress` without nftables.
