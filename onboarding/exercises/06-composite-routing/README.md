# 06 — Composite backend routing

**Language: Python.** `CompositeSandboxService` sends `templateId` (and fsb
snapshot restores) to FastSandbox, and routes existing IDs by the `fsb-`
prefix.

## Task

```python
def route_create(template_id: str | None, snapshot_backend: str | None) -> str
def route_existing(sandbox_id: str) -> str
```

Return `"fsb"` or `"kubernetes"`.

Create → fsb when `template_id` is non-blank **or**
`snapshot_backend == "fsb"`. Existing → fsb iff `sandbox_id.startswith("fsb-")`.

## Verify

```bash
python3 -m pytest onboarding/exercises/06-composite-routing/solution -q
```
