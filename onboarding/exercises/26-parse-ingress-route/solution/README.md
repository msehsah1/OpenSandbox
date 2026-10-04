# Solution steps

1. Strip `http://` / `https://`, then take the first DNS label.
2. Split on `-`. Join all but the last part as the sandbox id; `Atoi` the last.
3. For URI mode, `strings.Split` on `/` after `TrimPrefix(path, "/")`.
4. Remaining segments become `"/" + strings.Join(...)`. Empty remainder → `"/"`.

See `parseHostRoute` / `parseURIRoute` in
`components/ingress/pkg/proxy/host_route_parse.go`.
