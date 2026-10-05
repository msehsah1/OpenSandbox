# Solution steps

1. Format is `event: NAME\ndata: DATA\n\n`.
2. Flush is the caller’s problem (`http.Flusher` in a real handler).
3. Keep JSON on the complete event only — stdout is raw text.

See `setServerEventsHandler` / `writeSingleEvent` in execd controllers.
