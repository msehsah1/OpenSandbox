# 37 — Parse secure-access keys

**Language: Go.** Ingress `--secure-access-keys` is
`a=BASE64,b=BASE64`. Each `key_id` is exactly one character `[0-9a-z]`
(uppercase rejected). Empty segments are skipped; an empty secret or
bad base64 is an error.

## Objective

Contribute ingress `--secure-access-keys` parsing without dropping a
key or accepting `A=` (uppercase). A bad parse can disable signed
routes or leave one kid unused. This drill is `ParseKeys`. Flag and
key-ring PRs must keep the one-char `[0-9a-z]` + standard-base64
contract.

## Task

```go
func ParseKeys(s string) (map[string][]byte, error)
```

Reject empty input, missing `=`, empty key or value, `key_id` length ≠ 1,
uppercase key ids, invalid base64, and a string that yields no keys.

## Verify

```bash
cd onboarding/exercises
go test ./37-parse-access-keys/solution -count=1
```
