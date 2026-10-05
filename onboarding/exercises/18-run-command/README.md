# 18 — Run a shell command

**Language: Go.** Foreground `/command` is `exec.CommandContext` plus captured
stdout/stderr/exit. This is the core of `runtime.runCommand` without SSE.

## Objective

Contribute foreground exec without confusing process failure with
daemon failure. Non-zero exit is a `Result`, not a Go error; only
start failures are errors. This drill is `runtime.runCommand`. You
need this shape for any execd PR that changes how stdout/stderr/exit
are captured.

## Task

```go
type Result struct {
    Stdout   string
    Stderr   string
    ExitCode int
}

func Run(command string) (Result, error)
```

Use the platform shell (`sh -c` on Unix). A missing binary or non-zero exit
is a `Result` with `ExitCode`, not a Go error. Only fail the Go error for
start problems you cannot represent as an exit code.

## Verify

```bash
cd onboarding/exercises
go test ./18-run-command/solution -count=1
```
