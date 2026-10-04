# 09 — Host-path allowlist

**Language: Python.** Host bind mounts are checked against
`[storage].allowed_host_paths`. Empty allowlist = allow all (the example
config warns this is not for production). `..` must not escape a prefix.

## Task

`is_allowed_host_path(path: str, allowed: list[str]) -> bool`

- Empty `allowed` → `True`
- Resolve both sides with `Path.resolve()`
- Allowed when the path equals a prefix or is under it (`relative_to` works)

## Verify

```bash
python3 -m pytest onboarding/exercises/09-host-path-allowlist/solution -q
```
