# Solution steps

1. Require `://` before `url.Parse` — otherwise `host:port` is an opaque URL.
2. `u.Hostname()` then `TrimRight(..., ".")`.
3. Empty host → not ok. Empty port → `443`/`80` from the scheme.
4. Do not invent a host from env or node IP here; that is a different helper.

See `parseOTLPEndpoint` in `components/internal/telemetry/endpoint.go`.
