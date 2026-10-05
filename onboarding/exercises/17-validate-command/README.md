# 17 — Validate run-command

**Language: Go.** execd binds JSON then calls `RunCommandRequest.Validate()`
before starting a process. Bad cwd/empty command must be 400, not a panic
inside `exec.Command`.

## Objective

Contribute execd request validation so bad input is 400, not a panic
in `exec.Command`. Empty command, negative timeout, and relative cwd
are rejected in `RunCommandRequest.Validate()` before a process
starts. This drill is that gate — skip it and a command-API PR
becomes a 500 under fuzz.

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
