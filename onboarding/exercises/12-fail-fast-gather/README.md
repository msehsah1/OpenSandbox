# 12 — Fail-fast gather

**Language: Python.** After create, the SDK resolves execd and egress
endpoints together. `asyncio.gather` would leave the sibling retrying after
a 401. `_gather_fail_fast` cancels the rest.

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
