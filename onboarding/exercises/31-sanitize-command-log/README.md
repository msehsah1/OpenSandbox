# 31 — Sanitize command logs

**Language: Go.** execd logs the command line. Tokens and `password=` values
must not appear in logs. `SanitizeCommand` / `MaskToken` live in
`components/execd/pkg/log/sanitize.go`.

## Objective

Contribute execd logging without leaking secrets. Command lines
carry `password=`, `token=`, and `Authorization: Bearer`. Those
values must not appear in logs. This drill is
`components/execd/pkg/log/sanitize.go`. Reviewers will reject a
logging PR that prints the raw command.

## Task

```go
func MaskToken(token string) string
func SanitizeCommand(cmd string) string
```

`MaskToken`: length ≤ 8 → `"****"`; otherwise first 4 + `****` + last 4.

`SanitizeCommand`: case-insensitive, replace
`password=VALUE`, `token=VALUE`, `Authorization: Bearer VALUE` with the key
kept and the value masked via `MaskToken`. `VALUE` is a run of non-space
characters.

## Verify

```bash
cd onboarding/exercises
go test ./31-sanitize-command-log/solution -count=1
```
