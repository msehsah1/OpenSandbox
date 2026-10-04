# 22 — Background command status

**Language: Go.** Background `/command` returns an id immediately. Clients
poll `GET /command/status/:id`. Status must distinguish running vs finished
and keep the exit code.

## Task

```go
type Store struct { ... }
func (s *Store) Start(run func() int) string
func (s *Store) Status(id string) (running bool, exitCode int, ok bool)
```

`Start` launches `run` in a goroutine and returns an id. `Status` is safe
for concurrent use.

## Verify

```bash
cd onboarding/exercises
go test ./22-background-status/solution -count=1
```
