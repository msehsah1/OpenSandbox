# 02 — Create-request XOR

**Language: Python.** `Sandbox.create` and the lifecycle body require exactly
one startup source: `image` or `snapshot_id`. Getting this rule wrong is a
common SDK/server mismatch.

## Task

Implement `CreateSandboxRequest` (Pydantic v2) with optional `image` and
`snapshot_id`. A model validator rejects both-set and both-missing with a
clear `ValueError` message containing `exactly one`.

## Verify

```bash
python3 -m pytest onboarding/exercises/02-create-request-xor/solution -q
```
