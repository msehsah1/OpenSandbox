# Solution steps

1. `strings.HasPrefix(sandboxID, "fsb-")`.
2. Do not parse UUIDs. The server’s composite service is this one prefix.
3. `Dispatch` exists so you practice not duplicating the predicate at each
   call site (`_backend` on the Python service).
