# 14 — SSE command client

**Language: Python.** Command execution is handwritten HTTP + SSE, not the
generated client. You must parse `event:` / `data:` lines and infer an exit
code from `execution_complete` / stderr errors.

## Task

`parse_sse(payload: str) -> list[dict]` where each event is
`{"event": name, "data": raw}`.

`execution_from_events(events) -> dict` with `stdout` (list of data strings
for `stdout` events) and `exit_code` (from JSON data on `execution_complete`,
default `0` if missing).

## Verify

```bash
python3 -m pytest onboarding/exercises/14-sse-command-client/solution -q
```
