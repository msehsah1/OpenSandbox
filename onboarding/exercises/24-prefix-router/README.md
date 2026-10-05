# 24 — ID-prefix router

**Language: Go.** The same split as exercise 06, in Go, because the
Kubernetes controller and ingress also branch on sandbox identity. Keep the
rule identical: IDs starting with `fsb-` are FastSandbox.

## Objective

Contribute Go-side backend dispatch that matches the Python
composite service (exercise 06). Ingress and the Kubernetes
controller also branch on `fsb-`. If the prefix rule drifts between
languages, FastSandbox IDs hit the wrong provider. This drill is
that shared identity split.

## Task

```go
func BackendFor(sandboxID string) string // "fsb" or "kubernetes"
func Dispatch(sandboxID string, fsb, kube func(string) string) string
```

`Dispatch` calls exactly one backend function with the id.

## Verify

```bash
cd onboarding/exercises
go test ./24-prefix-router/solution -count=1
```
