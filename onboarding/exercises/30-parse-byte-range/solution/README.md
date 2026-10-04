# Solution steps

1. Require the `bytes=` prefix.
2. Split once on `-`. Empty start → suffix-length. Empty end → through EOF.
3. Clamp `end` with `min(end, size-1)` **before** `length = end - start + 1`.
4. That clamp is the bugfix from `fix(execd): clamp the Range end...`.

See `ParseRange` and the commit on this checkout’s history.
