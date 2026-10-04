# Solution steps

1. Call `parse_qsl` inside `try` and wrap any exception.
2. `dict(parsed)` last-key-wins, same as the server.
3. Do not write your own splitter — `%26` and `+` are easy to get wrong.

See `list_sandboxes` in `server/opensandbox_server/api/lifecycle.py`.
