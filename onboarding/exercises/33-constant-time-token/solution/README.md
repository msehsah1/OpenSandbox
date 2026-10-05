# Solution steps

1. Empty expected → `true` (local/dev).
2. `sha256.Sum256` both sides, then `subtle.ConstantTimeCompare` on the digests.
   That avoids leaking `len(provided) == len(expected)`.
3. Do not log either token.

Production also has to think about empty provided vs empty expected; here
empty expected is the only open mode.
