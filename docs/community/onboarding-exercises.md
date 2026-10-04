---
title: Onboarding Exercises
description: 35 Python and Go warm-up drills that train the patterns used in OpenSandbox contributions.
---

# Onboarding Exercises

Thirty-five small drills live in
[`onboarding/exercises/`](https://github.com/opensandbox-group/OpenSandbox/tree/main/onboarding/exercises).
They were extracted from the [contributor onboarding plan](/community/onboarding-plan).

Language follows the real subsystem:

- **Python** for the lifecycle server, Python SDK, and CLI
- **Go** for execd, ingress, egress, the Kubernetes controller, and shared runtime helpers

Each exercise folder has a task `README.md` and a `solution/` directory with
working code plus a short step-by-step `README.md`.

Start at the pack index:

[onboarding/exercises/README.md](https://github.com/opensandbox-group/OpenSandbox/blob/main/onboarding/exercises/README.md)

| # | Language | Drill |
| - | -------- | ----- |
| 01–15 | Python | Server config, auth, factory, SDK create, SSE client, CLI |
| 16–24 | Go | execd router, command exec, SSE, process groups, file safety |
| 25 | Python | Spec vs docs alignment (recommended first PR) |
| 26–28 | Go | Ingress host/URI parse, endpoints annotation, alloc-status compat |
| 29 | Go | Egress FQDN / wildcard match (deny-wins) |
| 30–31 | Go | execd Range clamp and command-log sanitization |
| 32–33 | Go | Hop-by-hop header strip and constant-time tokens |
| 34–35 | Go | Idempotent pause dispatch and requeue backoff |
