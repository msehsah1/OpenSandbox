# 10 — Error envelope

**Language: Python.** Lifecycle errors are `{"code": "...", "message": "..."}`,
not FastAPI’s default `{"detail": ...}` string. Tests and SDKs match on
`code`.

## Objective

Contribute API errors that SDKs and tests can match. Lifecycle
clients key off `{"code","message"}`, not FastAPI’s default
`{"detail": ...}`. Changing the envelope breaks every language SDK.
This drill is the `HTTPException` shape you must keep when adding a
new error code or handler.

## Task

`raise_sandbox_error(status: int, code: str, message: str)` raises
`HTTPException` whose `detail` is that object.

`install_error_handler(app)` makes uncaught `HTTPException` render `detail`
as the JSON body when `detail` is a dict.

A demo route `GET /boom` raises `INVALID_PARAMETER` / 400.

## Verify

```bash
python3 -m pytest onboarding/exercises/10-error-envelope/solution -q
```
