# 33 — Constant-time access token compare

**Language: Go.** execd and egress compare `X-EXECD-ACCESS-TOKEN` (and similar)
to a configured secret. `==` on strings is not constant-time and leaks length
via early return. Use `crypto/subtle`.

## Task

```go
func TokenValid(provided, expected string) bool
```

- Empty `expected` → allow (auth disabled), same as exercise 16.
- Otherwise compare in constant time. Different lengths must still be
  constant-time (hash both, or compare SHA-256 digests).

## Verify

```bash
cd onboarding/exercises
go test ./33-constant-time-token/solution -count=1
```
