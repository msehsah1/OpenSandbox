# Solution steps

1. Split on blank lines (`\n\n`).
2. Read `event:` and `data:` prefixes. Ignore comments (`:`).
3. Only JSON-decode `execution_complete` data; stdout stays text.
4. The real adapter is more involved (retry unwrap, background early break).
   This drill is the parser only.

See `CommandsAdapter` and `aiter_sse_events`.
