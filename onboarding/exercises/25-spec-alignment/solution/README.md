# Solution steps

1. Walk `paths` for a key ending with `/sandboxes` and method `post`.
2. Read `responses.202.description`.
3. Scan architecture text case-insensitively.
4. Do not auto-rewrite docs here — the test only reports drift. The real
   first PR edits `docs/architecture/index.md` §7.1 and
   `docs/architecture/control-plane/server.md`.

See `server/tests/test_routes_create_delete.py`
`test_create_sandbox_openapi_describes_synchronous_provisioning`.
