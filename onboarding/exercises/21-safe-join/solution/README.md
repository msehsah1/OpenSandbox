# Solution steps

1. `filepath.Clean` both sides.
2. Join, clean again.
3. `filepath.Rel(root, joined)` and reject `..` prefix or `..` itself.
4. Do not use string prefix checks only — ` /tmp/root-evil` vs `/tmp/root`.

execd filesystem handlers do this before any read/write.
