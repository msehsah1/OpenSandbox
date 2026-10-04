# 17 — Validate run-command

**Language: Go.** execd binds JSON then calls `RunCommandRequest.Validate()`
before starting a process. Bad cwd/empty command must be 400, not a panic
inside `exec.Command`.

## Task

```go
type RunCommandRequest struct {
    Command   string
    Cwd       string
    TimeoutMs int
}

func (r RunCommandRequest) Validate() error
```

Reject: empty/whitespace `Command`; `TimeoutMs < 0`; `Cwd` that is non-empty
and not absolute (`path.IsAbs`).

## Verify

```bash
cd onboarding/exercises
go test ./17-validate-command/solution -count=1
```
