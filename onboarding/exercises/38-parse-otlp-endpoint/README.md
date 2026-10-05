# 38 — Parse an OTLP endpoint URL

**Language: Go.** Shared telemetry (`components/internal/telemetry`)
reads `OTEL_EXPORTER_OTLP_*_ENDPOINT`. The exporter only accepts a URL
with a scheme (`https://host[:port][/path]`). Bare `host:port` is
invalid. A missing port uses the scheme default (`https`→443, `http`→80).
A trailing DNS root dot is stripped.

## Objective

Contribute telemetry and egress auto-allow for OTLP. The exporter
only accepts a URL with a scheme; bare `host:port` is an opaque URL
with an empty host. Wrong parse either blocks metrics or allowlists
the wrong destination. This drill is
`components/internal/telemetry/endpoint.go` — shared by execd,
ingress, and egress.

## Task

```go
func ParseOTLPEndpoint(raw string) (host, port string, ok bool)
```

Return `ok=false` for empty input, missing `://`, unparseable URLs, or
an empty host.

## Verify

```bash
cd onboarding/exercises
go test ./38-parse-otlp-endpoint/solution -count=1
```
