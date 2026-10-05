# Solution steps

1. Do not use `str.startswith` on unresolved paths —
   `/tmp/allowed/../etc/passwd` would pass.
2. `Path.resolve()` then `relative_to(prefix)`.
3. Catch `ValueError` from `relative_to` as deny.

This is the same idea as `ensure_volumes_valid` / Docker volume helpers.
