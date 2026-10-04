# Solution steps

1. Use Starlette `BaseHTTPMiddleware`.
2. Exempt with `request.url.path.startswith(path)` — the real server also
   exempts `/version`, `/docs`, `/redoc`, `/openapi.json`.
3. Read `OPEN-SANDBOX-API-KEY` (not `Authorization`).
4. Return `JSONResponse` yourself; do not raise inside middleware if you want
   a stable body.

Startup confirmation (`OPENSANDBOX_INSECURE_SERVER`) is out of scope.
