# Solution steps

1. Strip `template_id` before testing truthiness.
2. Do not treat a bare `snapshot_id` as fsb unless the resolved backend is
   `"fsb"` — that is `resolved_snapshot_backend` on the server.
3. Keep `_backend(id)` as a one-liner prefix check. Reviewers expect that
   split to stay boring.
