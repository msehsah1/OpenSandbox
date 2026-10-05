# Solution steps

1. Default status 500 if `Status == 0`.
2. `json.NewEncoder(w).Encode`.
3. Keep `Status` out of the body (`json:"-"`).

See execd `RespondError` and the lifecycle `ErrorResponse` model.
