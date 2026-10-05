# 14 — SSE command client

**Language: Python.** Command execution is handwritten HTTP + SSE, not the
generated client. You must parse `event:` / `data:` lines and infer an exit
code from `execution_complete` / stderr errors.

## Objective

Contribute command-streaming client changes without relying on the
generated OpenAPI client. SSE (`event:` / `data:`) is handwritten in
`CommandsAdapter`. Event names and `execution_complete` are a public
protocol shared with execd. This drill is the parser you must keep
aligned when you add an event type or change how exit codes are
inferred.

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
