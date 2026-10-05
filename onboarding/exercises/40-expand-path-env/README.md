# 40 — Expand a path with env vars

**Language: Go.** execd working directories accept `$VAR`, `${VAR}`, and
a leading `~`. Undefined variables must fail (do not silently expand to
empty). Inject the env map — do not read `os.Environ` — so the drill is
deterministic.

## Objective

Contribute execd working-directory handling without turning a missing
`$VAR` into `""` (which can resolve to `/`). `${VAR}`, `$VAR`, and
`~` are supported; undefined names must error. This drill is
`pathutil.ExpandPathWithEnv`. Required for cwd, volume, or path-flag
PRs in execd.

## Task

```go
func ExpandPath(path string, env map[string]string) (string, error)
```

- Empty path → `("", nil)`
- Names match `[A-Za-z_][A-Za-z0-9_]*`
- Missing names → error listing them (sorted, comma-separated)
- `~` and `~/...` use `env["HOME"]`; missing HOME is an error
- After expansion, return the string (do not require the path to exist)

## Verify

```bash
cd onboarding/exercises
go test ./40-expand-path-env/solution -count=1
```
