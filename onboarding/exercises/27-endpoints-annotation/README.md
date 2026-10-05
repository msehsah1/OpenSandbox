# 27 — Endpoints annotation

**Language: Go.** BatchSandbox publishes reachable addresses in
`sandbox.opensandbox.io/endpoints` as a JSON array. Ingress and the server
read that annotation; a missing or empty list is an error, not “no routes”.

## Objective

Contribute the BatchSandbox endpoints contract used by ingress and
the server. `sandbox.opensandbox.io/endpoints` is a JSON array;
missing or `[]` is an error, not “no routes.” Inventing a default IP
hides provision failures. `kubernetes/AGENTS.md` treats this key as
stability-sensitive. This drill is `pkg/utils/endpoints.go`.

## Task

```go
const AnnotationEndpoints = "sandbox.opensandbox.io/endpoints"

func GetEndpoints(annotations map[string]string) ([]string, error)
func SetEndpoints(annotations map[string]string, ips []string) error
```

`GetEndpoints` fails when the map is nil, the key is missing, JSON is invalid,
or the list is empty. `SetEndpoints` writes compact JSON.

## Verify

```bash
cd onboarding/exercises
go test ./27-endpoints-annotation/solution -count=1
```
