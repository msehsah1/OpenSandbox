# Solution steps

1. Implement `MaskToken` first — it is independently useful.
2. Use `regexp.MustCompile` with `(?i)` and a replace function.
3. Do not delete the flag name; reviewers want logs to stay readable.
4. The real sanitizer has more patterns. Three is enough for the drill.
