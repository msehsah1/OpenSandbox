# 11 — SDK create defaults

**Language: Python.** The Python SDK is the client facade maintainers review
for parity. `Sandbox.create` defaults entrypoint to `["tail", "-f", "/dev/null"]`
and resources to `{"cpu": "1", "memory": "2Gi"}`.

## Objective

Contribute Python SDK facade changes with cross-language parity in
mind. `Sandbox.create` defaults (entrypoint `tail -f /dev/null`,
cpu/memory) are a client contract. Changing them only in Python
desyncs JS/Go/Java and the docs. This drill is the defaulting and XOR
logic you will edit — and must update everywhere else — in an SDK PR.

## Task

`prepare_create(*, image=None, snapshot_id=None, entrypoint=None, resource=None) -> dict`

- XOR image/snapshot (`ValueError` with `Exactly one`)
- Apply those defaults when arguments are `None`
- String `image` becomes `{"image": image}`

## Verify

```bash
python3 -m pytest onboarding/exercises/11-sdk-create-defaults/solution -q
```
