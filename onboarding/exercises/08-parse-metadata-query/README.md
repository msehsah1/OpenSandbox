# 08 — Parse metadata query

**Language: Python.** `GET /v1/sandboxes?metadata=project%3DApollo%26env%3Dprod`
uses `urllib.parse.parse_qsl(..., strict_parsing=True)`. Malformed metadata
must become 400 `INVALID_METADATA_FORMAT`, not a 500.

## Objective

Contribute list-filter changes that fail 400, not 500. `GET
/v1/sandboxes?metadata=` is parsed with `parse_qsl(strict_parsing=True)`.
A malformed query must become `INVALID_METADATA_FORMAT`. This drill is
the helper in `api/lifecycle.py` so a list-API PR does not turn
operator typos into unhandled exceptions.

## Task

`parse_metadata(metadata: str | None) -> dict[str, str]`.

- `None` or `""` → `{}`
- `strict_parsing=True`, `keep_blank_values=True`
- On parse failure raise `ValueError` whose message starts with `Invalid metadata format:`

## Verify

```bash
python3 -m pytest onboarding/exercises/08-parse-metadata-query/solution -q
```
