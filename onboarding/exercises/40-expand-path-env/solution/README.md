# Solution steps

1. Scan with a regexp for `$NAME` / `${NAME}` and collect missing keys.
2. `os.Expand` only after validation — otherwise empty values hide typos.
3. Handle `~` yourself; `os.Expand` does not.
4. Prefer `filepath.Join(home, rest)` after stripping `~/`.

See `ExpandPathWithEnv` in `components/execd/pkg/util/pathutil/path.go`.
