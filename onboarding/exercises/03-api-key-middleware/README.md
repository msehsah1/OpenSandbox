# 03 — API-key middleware

**Language: Python.** Every lifecycle call except a short allowlist must send
`OPEN-SANDBOX-API-KEY`. This is the first middleware you will read in
`server/opensandbox_server/middleware/auth.py`.

## Task

Wrap a FastAPI app:

- Exempt `/health` (prefix match).
- If no keys are configured, allow all (local unauthenticated mode).
- Otherwise require the header and a matching key.
- Failures return 401 JSON `{"code":"MISSING_API_KEY"|"INVALID_API_KEY","message":...}`.

`create_app(api_key: str | None)` should return the app.

## Verify

```bash
python3 -m pytest onboarding/exercises/03-api-key-middleware/solution -q
```
