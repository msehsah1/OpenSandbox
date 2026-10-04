# Solution steps

1. Treat nil annotations as an error on read (the real helper does).
2. `json.Unmarshal` into `[]string`.
3. On write, allocate the map if needed; `json.Marshal` the slice.
4. Do not invent a default IP.

See `kubernetes/pkg/utils/endpoints.go`.
