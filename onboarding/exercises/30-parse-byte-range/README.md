# 30 — Parse a byte Range header

**Language: Go.** execd file downloads honor `Range: bytes=`. A recent fix
clamped the end **before** computing length so a huge end cannot overflow.
This is the same shape as `ParseRange` in
`components/execd/pkg/web/controller/range.go`.

## Objective

Contribute execd file-download Range handling without reintroducing
the overflow. End must be clamped to `size-1` **before**
`length = end - start + 1` (see `fix(execd): clamp the Range end…`).
This drill is `ParseRange`. A download PR that computes length first
fails the same class of huge-`end` request.

## Task

```go
type ByteRange struct {
    Start  int64
    Length int64
}

func ParseRange(header string, size int64) (ByteRange, error)
```

Support one range only: `bytes=start-end`, `bytes=start-`, `bytes=-suffix`.
Clamp `end` to `size-1`. Reject `start >= size`, missing `bytes=`, or
`end < start`.

## Verify

```bash
cd onboarding/exercises
go test ./30-parse-byte-range/solution -count=1
```
