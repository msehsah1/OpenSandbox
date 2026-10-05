# 20 — Cancel process group

**Language: Go.** Killing only the shell leaves children running. execd sets
`Setpgid: true` and signals the group on cancel/timeout. This is a frequent
execd bugfix area (see recent PRs around bash-session timeout).

## Objective

Contribute timeout/interrupt handling that actually stops the
workload. Killing only the shell leaves children running — a frequent
execd bugfix (bash-session timeout PRs). This drill is `Setpgid` +
signal-the-group, which you must keep when you touch cancel or
timeout paths.

## Task

`RunCancellable(ctx context.Context, command string) error`

- Start `sh -c command` in its own process group (`syscall.SysProcAttr{Setpgid: true}`)
- On `ctx.Done()`, send `SIGKILL` to the group (`-pid`)
- Wait for the process

Linux-only is fine (`//go:build linux` or just document it).

## Verify

```bash
cd onboarding/exercises
go test ./20-cancel-process-group/solution -count=1 -timeout 5s
```
