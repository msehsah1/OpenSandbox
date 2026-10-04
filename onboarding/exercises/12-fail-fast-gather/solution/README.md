# Solution steps

1. `asyncio.ensure_future` each awaitable.
2. `try: return list(await asyncio.gather(*tasks))`.
3. On failure, cancel pending tasks, then `await gather(..., return_exceptions=True)`.
4. Re-raise. Do not use `TaskGroup` if you want 3.10 compatibility — the SDK
   comments this explicitly.
