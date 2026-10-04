# Solution steps

1. Split on `,`, trim, skip blanks.
2. `strings.Cut(seg, "=")` — both sides required.
3. `key_id` is one byte in `[0-9a-z]`. Reject `A-Z` with a specific error.
4. `base64.StdEncoding.DecodeString` and reject a zero-length secret.

See `ParseKeys` in `components/ingress/pkg/signature/signature.go`.
