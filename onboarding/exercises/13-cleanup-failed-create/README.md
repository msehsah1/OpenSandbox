# 13 — Cleanup after failed create

**Language: Python.** If create succeeds but readiness fails, `Sandbox._launch`
tries `kill_sandbox` so Docker does not leak containers. Forgetting this is a
real bug class.

## Task

`async def launch(create, health, kill) -> str`

- `id = (await create())["id"]`
- `await health(id)`; on success return `id`
- On any later exception, `await kill(id)` (ignore kill errors), then re-raise

## Verify

```bash
python3 -m pytest onboarding/exercises/13-cleanup-failed-create/solution -q
```
