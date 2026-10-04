# 36 — Parse expires_b36 and signature format

**Language: Go.** Signed ingress routes encode expiry as a lowercase
base-36 Unix-seconds segment and a 9-character signature
(`8` lowercase hex + `1` key id `[0-9a-z]`). Bad format must fail
before any HMAC work.

## Task

```go
func ParseExpiresB36(s string) (uint64, error)
func FormatExpiresB36(sec uint64) string
func ValidateSignatureFormat(signature string) error
```

`ParseExpiresB36`: reject empty, uppercase, non `[0-9a-z]`, leading zeros
(`00`), and length > 13. `"0"` is valid.

## Verify

```bash
cd onboarding/exercises
go test ./36-parse-expires-b36/solution -count=1
```
