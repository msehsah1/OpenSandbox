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

from pathlib import Path

import pytest

from config import API_KEY_ENV_VAR, CONFIG_ENV_VAR, load_config


def test_loads_toml(tmp_path: Path, monkeypatch: pytest.MonkeyPatch) -> None:
    cfg = tmp_path / "sandbox.toml"
    cfg.write_text('[server]\napi_key = "from-file"\nport = 8080\n', encoding="utf-8")
    monkeypatch.delenv(API_KEY_ENV_VAR, raising=False)
    data = load_config(cfg)
    assert data["server"]["api_key"] == "from-file"
    assert data["server"]["port"] == 8080


def test_env_api_key_wins(tmp_path: Path, monkeypatch: pytest.MonkeyPatch) -> None:
    cfg = tmp_path / "sandbox.toml"
    cfg.write_text('[server]\napi_key = "from-file"\n', encoding="utf-8")
    monkeypatch.setenv(API_KEY_ENV_VAR, "from-env")
    assert load_config(cfg)["server"]["api_key"] == "from-env"


def test_path_env(tmp_path: Path, monkeypatch: pytest.MonkeyPatch) -> None:
    cfg = tmp_path / "alt.toml"
    cfg.write_text('[log]\nlevel = "DEBUG"\n', encoding="utf-8")
    monkeypatch.setenv(CONFIG_ENV_VAR, str(cfg))
    monkeypatch.delenv(API_KEY_ENV_VAR, raising=False)
    assert load_config()["log"]["level"] == "DEBUG"


def test_missing_file(tmp_path: Path) -> None:
    with pytest.raises(FileNotFoundError):
        load_config(tmp_path / "missing.toml")
