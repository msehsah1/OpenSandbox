# 25 — Spec/docs alignment test

**Language: Python.** Your recommended first contribution is to fix architecture
docs that still say create returns `Pending` while the spec and OpenAPI test
say create waits and returns `Running`. This drill is a regression test you
could add next to `test_create_sandbox_openapi_describes_synchronous_provisioning`.

## Objective

Contribute the recommended first docs PR: architecture text still
describes create as returning `Pending` / async, while
`specs/sandbox-lifecycle.yml` and the server wait and return
`Running` (HTTP 202). This drill is a regression test you could land
next to the OpenAPI create-path test — a real, reviewable change
that does not need Docker or a cluster.

## Task

`assert_create_contract(openapi: dict, architecture_text: str) -> list[str]`

Return a list of problem strings (empty = aligned):

1. `POST /v1/sandboxes` (or `/sandboxes`) 202 description must mention
   `provisioned successfully` (case-insensitive).
2. That description must mention `Running`.
3. `architecture_text` must **not** claim that creation returns before the
   sandbox is running as a standalone fact. Flag the phrase
   `returns before the sandbox is running` or
   `creation is asynchronous from the api perspective`.

## Verify

```bash
python3 -m pytest onboarding/exercises/25-spec-alignment/solution -q
```
