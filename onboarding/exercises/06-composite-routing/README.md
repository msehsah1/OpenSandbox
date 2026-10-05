# 06 — Composite backend routing

**Language: Python.** `CompositeSandboxService` sends `templateId` (and fsb
snapshot restores) to FastSandbox, and routes existing IDs by the `fsb-`
prefix.

## Objective

Contribute create/get/kill paths that stay on the correct backend.
`templateId` and `fsb-` IDs go to FastSandbox; everything else stays
on Kubernetes. A wrong branch orphans sandboxes or talks to the wrong
control plane. This drill is `CompositeSandboxService._backend` — the
first thing to get right in any composite-service PR.

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
