# Solution steps

1. Allocate `ring` of length `max`. `idx` is the next write; `filled` counts used slots.
2. `Record` writes `now()` then advances `idx` modulo `max`.
3. `Exceeded` is false until `filled == max`. Then the oldest slot is `ring[idx]`.
4. Compare `now().Sub(oldest) <= window`.

See `components/internal/supervisor/burst.go`.
