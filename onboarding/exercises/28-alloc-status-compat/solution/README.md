# Solution steps

1. Unmarshal into the struct with `omitempty` so old documents still decode.
2. Do not require `poolRef` on read.
3. When formatting, include `poolRef`/`generation` only when set so tests can
   still emit compact current JSON.
4. Never rename the `pods` key.

See `kubernetes/AGENTS.md` Annotation Contracts.
