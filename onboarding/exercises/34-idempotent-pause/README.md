# 34 — Idempotent pause dispatch

**Language: Go.** The controller must not dispatch pause again when the
BatchSandbox is already `Paused` (merged PR #2082). Reconcile can run twice.

## Objective

Contribute controller pause/resume without double-dispatch.
Reconcile can run twice; a BatchSandbox already `Paused` must not get
another pause job (merged PR #2082). This drill is that state
machine. Any pause, snapshot, or phase-transition PR needs
`changed=false` no-ops for already-terminal phases.

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
