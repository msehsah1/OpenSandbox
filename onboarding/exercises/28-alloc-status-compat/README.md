# 28 — Alloc-status backward compatibility

**Language: Go.** `sandbox.opensandbox.io/alloc-status` used to be
`{"pods":["pod-1"]}`. Current writes add `poolRef` and `generation`. Readers
must accept the legacy shape; writers should emit the current shape.

`kubernetes/AGENTS.md` treats this JSON as stability-sensitive.

## Task

```go
type AllocStatus struct {
    Pods       []string `json:"pods"`
    PoolRef    string   `json:"poolRef,omitempty"`
    Generation int64    `json:"generation,omitempty"`
}

func ParseAllocStatus(raw string) (AllocStatus, error)
func FormatAllocStatus(s AllocStatus) (string, error)
```

Legacy input without `poolRef` must parse. Empty `pods` is allowed (released
allocation). Invalid JSON is an error.

## Verify

```bash
cd onboarding/exercises
go test ./28-alloc-status-compat/solution -count=1
```
