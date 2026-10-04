# Solution steps

1. Declare both fields `str | None = None`.
2. Use `@model_validator(mode="after")`.
3. Treat whitespace-only `snapshot_id` as missing (`strip()`), matching the
   server’s `(request.template_id or "").strip()` style.
4. Do not default an image — the real SDK raises instead.

See `sdks/sandbox/python/src/opensandbox/sandbox.py` `Sandbox.create`.
