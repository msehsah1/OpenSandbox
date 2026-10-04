# 04 — Thin route + service

**Language: Python.** `server/AGENTS.md` says routes stay thin: validation in
the API layer, work in `SandboxService`. Business logic in a handler is a
review reject.

## Task

- `SandboxService` ABC with `async create_sandbox(request) -> dict`.
- `InMemorySandboxService` assigns an `id` and returns
  `{"id": ..., "status": {"state": "Running"}}`.
- FastAPI `POST /v1/sandboxes` body `{image: str}` returns **202** and only
  calls the service.

## Verify

```bash
python3 -m pytest onboarding/exercises/04-thin-route-service/solution -q
```
