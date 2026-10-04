# OpenSandbox contribution warm-up exercises

45 small drills extracted from the [contributor onboarding plan](../../docs/community/onboarding-plan.md).
Each one isolates one pattern you will meet in a real pull request.

Language is chosen from the real subsystem:

- **Python** for the lifecycle server, Python SDK, and CLI
- **Go** for execd, ingress, egress, the Kubernetes controller, and shared runtime helpers

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
  ./23-json-errors/solution ./24-prefix-router/solution \
  ./26-parse-ingress-route/solution ./27-endpoints-annotation/solution \
  ./28-alloc-status-compat/solution ./29-egress-domain-match/solution \
  ./30-parse-byte-range/solution ./31-sanitize-command-log/solution \
  ./32-strip-hop-headers/solution ./33-constant-time-token/solution \
  ./34-idempotent-pause/solution ./35-requeue-backoff/solution \
  ./36-parse-expires-b36/solution ./37-parse-access-keys/solution \
  ./38-parse-otlp-endpoint/solution ./39-burst-tracker/solution \
  ./40-expand-path-env/solution ./41-host-selector-overlap/solution \
  ./42-match-ip-cidr/solution ./43-normalize-interval-set/solution \
  ./44-oss-target-digest/solution ./45-classify-connect-error/solution
```

Run every Python reference:

```bash
cd onboarding/exercises
python3 -m pytest -q */solution
```

## The 45 exercises

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
| 26 | [Parse an ingress route](26-parse-ingress-route/) | Go | Host/URI routing | `components/ingress/pkg/proxy/host_route_parse.go` |
| 27 | [Endpoints annotation](27-endpoints-annotation/) | Go | BatchSandbox addresses | `kubernetes/pkg/utils/endpoints.go` |
| 28 | [Alloc-status compatibility](28-alloc-status-compat/) | Go | Annotation stability | `sandbox.opensandbox.io/alloc-status` |
| 29 | [Egress FQDN / wildcard](29-egress-domain-match/) | Go | Deny-wins policy | `components/egress/pkg/policy/policy.go` |
| 30 | [Parse a byte Range](30-parse-byte-range/) | Go | File download clamp | execd `ParseRange` |
| 31 | [Sanitize command logs](31-sanitize-command-log/) | Go | Secret masking | `components/execd/pkg/log/sanitize.go` |
| 32 | [Strip hop-by-hop headers](32-strip-hop-headers/) | Go | Reverse-proxy hygiene | `components/ingress/pkg/proxy/header.go` |
| 33 | [Constant-time token compare](33-constant-time-token/) | Go | Access-token auth | execd / egress `subtle` |
| 34 | [Idempotent pause dispatch](34-idempotent-pause/) | Go | Controller reconcile | pause skip when already Paused |
| 35 | [Requeue backoff](35-requeue-backoff/) | Go | Retry jitter | `components/internal/supervisor/backoff.go` |
| 36 | [Parse expires_b36](36-parse-expires-b36/) | Go | Signed-route format | `components/ingress/pkg/signature/signature.go` |
| 37 | [Parse secure-access keys](37-parse-access-keys/) | Go | Ingress key ring | `ParseKeys` (`key_id=base64`) |
| 38 | [Parse an OTLP endpoint](38-parse-otlp-endpoint/) | Go | Telemetry URL | `components/internal/telemetry/endpoint.go` |
| 39 | [Sliding-window burst](39-burst-tracker/) | Go | Crash-loop guard | `components/internal/supervisor/burst.go` |
| 40 | [Expand a path with env](40-expand-path-env/) | Go | Working-dir vars | execd `pathutil.ExpandPathWithEnv` |
| 41 | [Host selector overlap](41-host-selector-overlap/) | Go | OSEP-0023 algebra | `components/egress/pkg/hostselector` |
| 42 | [Match IP / CIDR](42-match-ip-cidr/) | Go | Egress IP targets | `netip` in `policy.go` |
| 43 | [Normalize nft intervals](43-normalize-interval-set/) | Go | Conflicting CIDRs | `components/egress/pkg/nftables/interval.go` |
| 44 | [OSS target digest](44-oss-target-digest/) | Go | Node-agent identity | `components/nodeagent/pkg/identity` |
| 45 | [Classify connect errors](45-classify-connect-error/) | Go | Dial taxonomy | ingress `connectivity.ClassifyConnectError` |

Work in order. 01–07 are the create path. 11–14 are the SDK path. 16–23 are execd. 25 is the docs-first contribution drill. 26–45 are Go-land drills for ingress, egress, controller, telemetry, supervisor, and node-agent helpers.
