# Solution steps

1. `exec.Command("sh", "-c", command)`.
2. `Combined` is wrong — keep stdout and stderr separate (`bytes.Buffer` on each).
3. `cmd.Run()` then `exitError, ok := err.(*exec.ExitError)`.
4. The real execd also sets `Setpgid` and tails files. You only need capture.

See `components/execd/pkg/runtime/command.go` `runCommand`.
