# Solution steps

1. `loop = asyncio.get_running_loop()` then `future = loop.create_future()`.
2. Worker `try/except` must `call_soon_threadsafe(future.set_result|set_exception)`.
3. Guard `set_exception` against `RuntimeError` if the loop is already closed
   — the real Docker service does this.
4. `await future` on the handler.

See `DockerSandboxService.create_sandbox` around the `Thread(target=_run)`.
