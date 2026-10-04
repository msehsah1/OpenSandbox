# 30 — Parse a byte Range header

**Language: Go.** execd file downloads honor `Range: bytes=`. A recent fix
clamped the end **before** computing length so a huge end cannot overflow.
This is the same shape as `ParseRange` in
`components/execd/pkg/web/controller/range.go`.

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
