# 01 — Load TOML config

**Language: Python.** The lifecycle server is a FastAPI process that boots from
`~/.sandbox.toml` (override `SANDBOX_CONFIG_PATH`) and can take
`OPENSANDBOX_SERVER_API_KEY` from the environment. Config loading is the first
server file you will touch.

## Task

Implement `load_config(path=None) -> dict` in `config.py`:

1. Resolve the file from `path`, else `SANDBOX_CONFIG_PATH`, else `~/.sandbox.toml`.
2. Parse TOML.
3. If `OPENSANDBOX_SERVER_API_KEY` is set and non-empty, it wins over
   `[server].api_key`.
4. Missing file → `FileNotFoundError`. Invalid TOML → `ValueError`.

## Verify

```bash
python3 -m pytest onboarding/exercises/01-load-toml-config/solution -q
```
