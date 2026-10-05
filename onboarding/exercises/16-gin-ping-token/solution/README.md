# Solution steps

1. `gin.New()` + `Recovery`. Release mode is optional in tests.
2. Register `/ping` before the auth middleware, or skip auth when path is `/ping`.
   execd uses an allowlist (`preInitPaths`) including `/ping`, `/ready`,
   `/internal/init`.
3. Compare the header to the configured token only when the token is non-empty.
4. Do not invent Bearer tokens — execd uses `X-EXECD-ACCESS-TOKEN`.
