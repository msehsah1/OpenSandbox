# Solution steps

1. Mutex around the map. Do not return a pointer the caller can race.
2. Mark running before starting the goroutine.
3. On finish, store exit code and `running=false`.
4. Unknown id → `ok == false` (the real API maps this to 404).

See `GetCommandStatus` and `commandKernel` in execd.
