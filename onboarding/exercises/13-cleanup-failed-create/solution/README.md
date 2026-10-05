# Solution steps

1. Track `sandbox_id` only after create returns.
2. Cleanup in `except` / `finally`-style `if sandbox_id:` — not before create.
3. Swallow cleanup failures so they do not mask the original error (log in
   production; tests just `pass`).

See `Sandbox._launch` in `sdks/sandbox/python/src/opensandbox/sandbox.py`.
