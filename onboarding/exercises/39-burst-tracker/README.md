# 39 — Sliding-window burst tracker

**Language: Go.** The shared supervisor refuses to keep restarting a
child that crash-loops. A ring of size `BurstMax` records launch times;
the burst is exceeded when the ring is full **and** the oldest of those
launches is still inside `BurstWindow`.

## Task

```go
func NewBurstTracker(max int, window time.Duration, now func() time.Time) *BurstTracker
func (b *BurstTracker) Record()
func (b *BurstTracker) Exceeded() bool
```

`max < 1` becomes 1. Inject `now` so tests are deterministic. Do not
call `time.Now` inside the tracker.

## Verify

```bash
cd onboarding/exercises
go test ./39-burst-tracker/solution -count=1
```
