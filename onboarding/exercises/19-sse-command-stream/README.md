# 19 — SSE command stream

**Language: Go.** execd writes Server-Sent Events (`event:` / `data:` /
blank line) as the process emits lines. The SDK client from exercise 14
consumes this shape.

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
