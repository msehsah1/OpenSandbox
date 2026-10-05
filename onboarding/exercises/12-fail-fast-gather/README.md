# 12 — Fail-fast gather

**Language: Python.** After create, the SDK resolves execd and egress
endpoints together. `asyncio.gather` would leave the sibling retrying after
a 401. `_gather_fail_fast` cancels the rest.

## Objective

Contribute endpoint-readiness code that cancels siblings on the first
failure. After create, the SDK resolves execd and egress together; a
401 must not leave the other side retrying. This drill is
`_gather_fail_fast` in `sandbox.py`. You will need it when you add
another parallel probe or change create-time health checks.

## Task

`async def gather_fail_fast(*awaitables) -> list`

- First exception cancels unfinished siblings
- Await cancelled siblings with `return_exceptions=True`
- Re-raise the original exception
- Preserve `CancelledError`

## Verify

```bash
python3 -m pytest onboarding/exercises/12-fail-fast-gather/solution -q
```
