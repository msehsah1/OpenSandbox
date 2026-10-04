# 21 — Reject path traversal

**Language: Go.** execd file APIs join a sandbox root with a user path.
`filepath.Join(root, "../etc/passwd")` can escape if you do not check the
result stays under root.

## Task

`SafeJoin(root, userPath string) (string, error)`

- Reject empty `userPath`
- If `userPath` is absolute, treat it as relative to `root` after stripping
  the volume (`filepath.Rel` / clean) **or** simply reject absolute paths
- After `filepath.Join(root, userPath)` + `Clean`, the result must equal
  `root` or be a child (`rel` does not start with `..`)

## Verify

```bash
cd onboarding/exercises
go test ./21-safe-join/solution -count=1
```
