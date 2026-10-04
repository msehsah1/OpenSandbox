# Solution steps

1. `strings.TrimSpace` the command. The Python SDK also rejects empty strings.
2. Absolute cwd only — relative cwd is host-dependent inside a sandbox.
3. Return `fmt.Errorf` with a stable field name so API messages stay greppable.

See `pkg/web/model/codeinterpreting.go` `Validate`.
