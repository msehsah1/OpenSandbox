# Solution steps

1. Encode the state machine as a switch — do not “set Pausing anyway”.
2. Already `Paused` / `Pausing` → no-op. That is the #2082 lesson.
3. Unknown phases should not dispatch.

See `fix(kubernetes): skip pause dispatch for an already Paused BatchSandbox`.
