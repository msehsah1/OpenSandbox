# Solution steps

1. `argparse` with subparsers. Do not open sockets.
2. `run` uses `nargs=argparse.REMAINDER` after `--`.
3. Return process-style exit codes; do not `sys.exit` inside `main` so tests
   can call it.

See `cli/src/opensandbox_cli/commands/` for the real split.
