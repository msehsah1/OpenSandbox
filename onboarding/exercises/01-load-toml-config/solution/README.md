# Solution steps

1. Use `tomllib` on 3.11+ (`tomli` fallback is fine; this reference uses
   `tomllib` / `tomli` via a tiny helper).
2. Expand `~` with `Path.expanduser()`.
3. After parse, copy `[server]` if present and overlay the env var.
4. Keep the function pure: no global singleton. The real server caches config;
   you do not need that yet.

This matches `server/opensandbox_server/config.py` (`CONFIG_ENV_VAR`,
`API_KEY_ENV_VAR`) at a much smaller scale.
