# 05 — Runtime factory

**Language: Python.** `create_sandbox_service()` in
`server/opensandbox_server/services/factory.py` only accepts `docker` and
`kubernetes`. FastSandbox is composed under Kubernetes, not a third
`runtime.type`.

## Task

`create_sandbox_service(runtime_type: str) -> str` (return a label, not a
real client):

- `docker` → `"docker"`
- `kubernetes` → `"composite"`
- anything else → `ValueError` listing supported types

Case-insensitive.

## Verify

```bash
python3 -m pytest onboarding/exercises/05-runtime-factory/solution -q
```
