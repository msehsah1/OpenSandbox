# 23 — Structured JSON errors

**Language: Go.** execd `RespondError` writes `{"code","message"}` with an
HTTP status. Handlers should not `c.String(400, err.Error())`.

## Task

```go
type APIError struct {
    Status  int    `json:"-"`
    Code    string `json:"code"`
    Message string `json:"message"`
}

func WriteError(w http.ResponseWriter, err APIError)
```

Set `Content-Type: application/json`.

## Verify

```bash
cd onboarding/exercises
go test ./23-json-errors/solution -count=1
```
