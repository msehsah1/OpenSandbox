# Solution steps

1. `*.base` vs exact: store `base` + a bool.
2. `validHost`: must contain `.`, not parse as an IP, labels 1–63, `[a-z0-9-]`, no leading/trailing `-`.
3. `Matches`: lowercase + strip one trailing dot; wildcard → `HasSuffix(host, "."+base)`.
4. Two wildcards overlap if equal or one base is a suffix of the other.

See `components/egress/pkg/hostselector/selector.go`.
