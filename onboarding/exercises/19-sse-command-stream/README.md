# 19 — SSE command stream

**Language: Go.** execd writes Server-Sent Events (`event:` / `data:` /
blank line) as the process emits lines. The SDK client from exercise 14
consumes this shape.

## Objective

Contribute the execd side of the command-stream protocol. SDKs
(exercise 14) parse `event:` / `data:` and
`execution_complete`. Renaming events or changing the JSON is a
cross-repo break. This drill is the writer in `RunCommand` so a
streaming PR stays wire-compatible.

## Task

`WriteSSE(w io.Writer, event, data string) error` writes one event.

`StreamCommand(w io.Writer, lines []string, exitCode int) error` writes a
`stdout` event per line, then `execution_complete` with JSON
`{"exit_code":N}`.

## Verify

```bash
cd onboarding/exercises
go test ./19-sse-command-stream/solution -count=1
```
