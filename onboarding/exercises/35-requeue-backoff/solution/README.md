# Solution steps

1. Inject `rng` — do not call `rand` in the pure function so tests are deterministic.
2. Double, then clamp, then jitter, then clamp again.
3. `rng()*2 - 1` maps `[0,1)` to `[-1,1)`.
4. Midpoint `rng=0.5` means zero jitter — use that in tests.

See `components/internal/supervisor/backoff.go`.
