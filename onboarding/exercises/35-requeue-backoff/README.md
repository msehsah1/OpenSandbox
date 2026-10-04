# 35 — Requeue backoff

**Language: Go.** Controllers and the shared supervisor retry with doubled
delay, min/max clamps, and jitter. A tight retry loop without backoff will
fail Kind e2e and production apiservers.

## Task

```go
func NextBackoff(prev, min, max time.Duration, jitter float64, rng func() float64) time.Duration
```

- `prev <= 0` → start at `min` (then apply jitter)
- otherwise `prev * 2`, clamp to `[min, max]`, then jitter `±jitter*value`
- clamp again so jitter cannot exceed max or go below `min` (or 1ns)
- `rng` returns `[0,1)` so tests can pass a stub (`func() float64 { return 0.5 }`)

## Verify

```bash
cd onboarding/exercises
go test ./35-requeue-backoff/solution -count=1
```
