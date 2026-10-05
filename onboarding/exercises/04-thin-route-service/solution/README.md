# Solution steps

1. Keep `create_sandbox` on the router to `return await service.create_sandbox(...)`.
2. Inject the service with `app.state` or a module-level default you can
   monkeypatch — the real server uses a module-level
   `sandbox_service = create_sandbox_service()`.
3. Status 202 matches `POST /v1/sandboxes` even though provision is awaited.

See `server/opensandbox_server/api/lifecycle.py` `create_sandbox`.
