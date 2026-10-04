# 15 — Thin CLI wrapper

**Language: Python.** `osb` commands are thin wrappers over the Python SDK.
`cli/AGENTS.md` rejects reimplementing HTTP in a new stable command.

## Task

`main(argv: list[str], sandbox: SandboxPort) -> int`

- `create --image IMAGE` prints the id, exit 0
- `run SANDBOX_ID -- command...` prints stdout, exit = command exit code
- Unknown command → return 2

`SandboxPort` is a protocol with `create(image) -> str` and
`run(sandbox_id, command: list[str]) -> tuple[str, int]`.

## Verify

```bash
python3 -m pytest onboarding/exercises/15-thin-cli-wrapper/solution -q
```
