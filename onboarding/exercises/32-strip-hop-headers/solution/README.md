# Solution steps

1. Build a deny set of `http.CanonicalHeaderKey` names.
2. Copy allowed headers into a new map — do not mutate the input.
3. Compare keys after canonicalization so `transfer-encoding` is stripped.

See `components/ingress/pkg/proxy/header.go`.
