# 07 — Offload blocking create

**Language: Python.** Docker create is awaited, but the Docker SDK is
blocking. The server starts a daemon thread and completes an
`asyncio.Future`. You will break the event loop if you call Docker on the
main coroutine.

## Objective

Contribute Docker-service changes without freezing the control plane.
The Docker SDK is blocking; calling it on the FastAPI event loop
stalls every other create. This drill is the daemon-thread +
`asyncio.Future` pattern in `DockerSandboxService.create_sandbox`,
which you must keep (or replace with an equivalent offload) whenever
you touch that path.

## Task

`async def create_sandbox(provision, *args)`:

- Run `provision(*args)` on a daemon thread.
- Await the result.
- Propagate exceptions from the worker.

`provision` is any blocking callable (the tests use `time.sleep`).

## Verify

```bash
python3 -m pytest onboarding/exercises/07-offload-blocking-create/solution -q
```
