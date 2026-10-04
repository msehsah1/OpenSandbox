# Copyright 2026 The OpenSandbox Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

from __future__ import annotations

import os
from pathlib import Path
from typing import Any

try:
    import tomllib
except ModuleNotFoundError:  # Python 3.10
    import tomli as tomllib  # type: ignore[no-redef]

CONFIG_ENV_VAR = "SANDBOX_CONFIG_PATH"
API_KEY_ENV_VAR = "OPENSANDBOX_SERVER_API_KEY"
DEFAULT_CONFIG_PATH = Path.home() / ".sandbox.toml"


def load_config(path: str | os.PathLike[str] | None = None) -> dict[str, Any]:
    resolved = Path(path) if path is not None else Path(
        os.environ.get(CONFIG_ENV_VAR, DEFAULT_CONFIG_PATH)
    )
    resolved = resolved.expanduser()
    if not resolved.is_file():
        raise FileNotFoundError(f"config not found: {resolved}")
    try:
        data = tomllib.loads(resolved.read_text(encoding="utf-8"))
    except tomllib.TOMLDecodeError as exc:
        raise ValueError(f"invalid TOML in {resolved}: {exc}") from exc

    server = dict(data.get("server") or {})
    env_key = os.environ.get(API_KEY_ENV_VAR, "").strip()
    if env_key:
        server["api_key"] = env_key
        data["server"] = server
    return data
