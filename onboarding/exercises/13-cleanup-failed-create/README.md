# 13 — Cleanup after failed create

**Language: Python.** If create succeeds but readiness fails, `Sandbox._launch`
tries `kill_sandbox` so Docker does not leak containers. Forgetting this is a
real bug class.

## Objective

Contribute create-path changes that do not leak sandboxes. If
provisioning succeeds but readiness fails, `Sandbox._launch` must
`kill_sandbox`. Forgetting that is a recurring bug (zombie Docker
containers / K8s objects). This drill is the try/health/kill pattern
reviewers look for on any `_launch` or create-cleanup PR.

## Task

`async def launch(create, health, kill) -> str`

- `id = (await create())["id"]`
- `await health(id)`; on success return `id`
- On any later exception, `await kill(id)` (ignore kill errors), then re-raise

## Verify

```bash
python3 -m pytest onboarding/exercises/13-cleanup-failed-create/solution -q
```
