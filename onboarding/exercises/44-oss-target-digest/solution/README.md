# Solution steps

1. `url.Parse`; require `https`, non-empty host, no user/query/fragment, path empty or `/`.
2. Lowercase hostname. Append `net.JoinHostPort` only when the port is not `443`.
3. Length-prefix parts so `ab`+`c` ≠ `a`+`bc`.
4. Do not include a trailing slash on the origin.

See `CanonicalOSSEndpoint` / `digest` in `components/nodeagent/pkg/identity/identity.go`.
