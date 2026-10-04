# Solution steps

1. Walk runes: only `0-9` and `a-z`. Reject `A-Z` explicitly so the error is clear.
2. Reject `len > 1 && s[0] == '0'`. Then `strconv.ParseUint(s, 36, 64)`.
3. Format with `strconv.FormatUint(sec, 36)` — it already omits leading zeros.
4. Signature is exactly 9 bytes: first 8 lowercase hex, last `[0-9a-z]`.

See `ParseExpiresB36` / `ValidateSignatureFormat` in
`components/ingress/pkg/signature/signature.go`.
