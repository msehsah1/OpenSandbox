# 34 — Idempotent pause dispatch

**Language: Go.** The controller must not dispatch pause again when the
BatchSandbox is already `Paused` (merged PR #2082). Reconcile can run twice.

## Task

```go
type Phase string
const (
    PhaseRunning Phase = "Running"
    PhasePaused  Phase = "Paused"
    PhasePausing Phase = "Pausing"
)

func ShouldDispatchPause(current Phase) bool
func NextPausePhase(current Phase) (Phase, bool) // next, changed
```

`ShouldDispatchPause` is true only from `Running`. `NextPausePhase` returns
`Pausing` from `Running`, otherwise the same phase and `changed=false`.

## Verify

```bash
cd onboarding/exercises
go test ./34-idempotent-pause/solution -count=1
```
