# Solution steps

1. `cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}`.
2. Start, then `select` on `ctx.Done()` vs a wait goroutine.
3. `syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)`.
4. Always `Wait()` so you do not leak zombies.

See `runCommand` + `subscribeCommandSignals` and the cancel tests in
`command_signal_test.go`.
