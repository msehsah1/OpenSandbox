# OpenSandbox contribution warm-up exercises

25 small drills extracted from the [contributor onboarding plan](../../docs/community/onboarding-plan.md).
Each one isolates one pattern you will meet in a real pull request.

Language is chosen from the real subsystem:

- **Python** for the lifecycle server, Python SDK, and CLI
- **Go** for execd (and the same request/validation habits used in ingress/egress)

You do not need Docker, Kubernetes, or a running OpenSandbox server.

## How to use each exercise

1. Read the exercise `README.md` (task, signatures, and why the language was chosen).
2. Implement it yourself in a scratch file or by copying the signatures.
3. Compare with `solution/` and read `solution/README.md` for the steps.
4. Run the tests from `solution/` to confirm the reference works.

```text
onboarding/exercises/
  NN-short-name/
    README.md           # the drill
    solution/
      README.md         # how the reference is built
      ...               # working code + tests
```

## Setup

Python 3.10+ (3.11+ preferred for `tomllib`):

```bash
cd onboarding/exercises
python3 -m pip install -r requirements.txt
```

Go 1.22+ is enough for these drills (production execd wants 1.25):

```bash
cd onboarding/exercises
go test ./16-gin-ping-token/solution ./17-validate-command/solution \
  ./18-run-command/solution ./21-safe-join/solution \
  ./23-json-errors/solution ./24-prefix-router/solution
```

Run every Python reference:

```bash
cd onboarding/exercises
python3 -m pytest -q */solution
```

## The 25 exercises

| # | Exercise | Language | Trains you for | Maps to |
| - | -------- | -------- | -------------- | ------- |
| 01 | [Load TOML config](01-load-toml-config/) | Python | Server boot config | `server/opensandbox_server/config.py` |
| 02 | [Create-request XOR](02-create-request-xor/) | Python | Lifecycle request rules | `Sandbox.create`, `CreateSandboxRequest` |
| 03 | [API-key middleware](03-api-key-middleware/) | Python | Server auth | `middleware/auth.py` |
| 04 | [Thin route + service](04-thin-route-service/) | Python | Keep routes thin | `api/lifecycle.py` + `SandboxService` |
| 05 | [Runtime factory](05-runtime-factory/) | Python | `runtime.type` dispatch | `services/factory.py` |
| 06 | [Composite backend routing](06-composite-routing/) | Python | K8s vs FastSandbox | `composite_service.py` |
| 07 | [Offload blocking create](07-offload-blocking-create/) | Python | Docker create thread | `DockerSandboxService.create_sandbox` |
| 08 | [Parse metadata query](08-parse-metadata-query/) | Python | List filters | `api/lifecycle.py` `list_sandboxes` |
| 09 | [Host-path allowlist](09-host-path-allowlist/) | Python | Volume safety | Docker/K8s volume helpers |
| 10 | [Error envelope](10-error-envelope/) | Python | `{code, message}` errors | server `HTTPException` detail |
| 11 | [SDK create defaults](11-sdk-create-defaults/) | Python | Client facade | `sandbox.py` `Sandbox.create` |
| 12 | [Fail-fast gather](12-fail-fast-gather/) | Python | Endpoint readiness | `sandbox.py` `_gather_fail_fast` |
| 13 | [Cleanup after failed create](13-cleanup-failed-create/) | Python | No zombie sandboxes | `Sandbox._launch` |
| 14 | [SSE command client](14-sse-command-client/) | Python | Exec streaming | `CommandsAdapter.run` |
| 15 | [Thin CLI wrapper](15-thin-cli-wrapper/) | Python | CLI over SDK | `cli/` command groups |
| 16 | [Gin ping + access token](16-gin-ping-token/) | Go | execd router/auth | `components/execd/pkg/web/router.go` |
| 17 | [Validate run-command](17-validate-command/) | Go | Request validation | `RunCommandRequest.Validate` |
| 18 | [Run a shell command](18-run-command/) | Go | Foreground exec | `runtime.runCommand` |
| 19 | [SSE command stream](19-sse-command-stream/) | Go | Streaming output | `RunCommand` SSE |
| 20 | [Cancel process group](20-cancel-process-group/) | Go | Interrupt/timeout | execd process-group kill |
| 21 | [Reject path traversal](21-safe-join/) | Go | File APIs | execd filesystem handlers |
| 22 | [Background command status](22-background-status/) | Go | `/command/status` | `GetCommandStatus` |
| 23 | [Structured JSON errors](23-json-errors/) | Go | Error codes | execd `RespondError` |
| 24 | [ID-prefix router](24-prefix-router/) | Go | `fsb-` backend split | `CompositeSandboxService._backend` |
| 25 | [Spec/docs alignment test](25-spec-alignment/) | Python | First recommended PR | create-path 202 vs Pending docs |

Work in order. 01–07 are the create path. 11–14 are the SDK path. 16–23 are execd. 25 is the docs-first contribution drill.
