# 16 — Gin ping + access token

**Language: Go.** execd is a Gin daemon. `/ping` is reachable before the
runtime-init gate and without `X-EXECD-ACCESS-TOKEN`. Business routes are not.

## Task

`NewRouter(accessToken string) *gin.Engine`

- `GET /ping` → `{"status":"ok"}` always
- `POST /command` requires header `X-EXECD-ACCESS-TOKEN` equal to `accessToken`
  when the token is non-empty; otherwise 401 `{"code":"UNAUTHORIZED","message":...}`
- With a valid token (or empty configured token), `POST /command` → 200
  `{"accepted":true}`

## Verify

```bash
cd onboarding/exercises
go test ./16-gin-ping-token/solution -count=1
```
