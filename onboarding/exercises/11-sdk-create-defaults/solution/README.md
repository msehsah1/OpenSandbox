# Solution steps

1. Reuse the XOR check from exercise 02 — do not invent a different message.
2. Defaults belong in the facade, not the generated OpenAPI client.
3. Do not invent a default image. The real SDK requires the caller to pass one.

See `Sandbox.create` in `sdks/sandbox/python/src/opensandbox/sandbox.py`.
