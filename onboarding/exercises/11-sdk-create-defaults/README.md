# 11 — SDK create defaults

**Language: Python.** The Python SDK is the client facade maintainers review
for parity. `Sandbox.create` defaults entrypoint to `["tail", "-f", "/dev/null"]`
and resources to `{"cpu": "1", "memory": "2Gi"}`.

## Task

`prepare_create(*, image=None, snapshot_id=None, entrypoint=None, resource=None) -> dict`

- XOR image/snapshot (`ValueError` with `Exactly one`)
- Apply those defaults when arguments are `None`
- String `image` becomes `{"image": image}`

## Verify

```bash
python3 -m pytest onboarding/exercises/11-sdk-create-defaults/solution -q
```
