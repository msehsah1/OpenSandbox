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

from paths import is_allowed_host_path


def test_empty_allowlist_permits_all(tmp_path: Path) -> None:
    assert is_allowed_host_path(str(tmp_path / "x"), [])


def test_prefix_and_escape(tmp_path: Path) -> None:
    allowed = tmp_path / "data"
    allowed.mkdir()
    outside = tmp_path / "etc"
    outside.mkdir()
    (allowed / "vol").mkdir()
    assert is_allowed_host_path(str(allowed / "vol"), [str(allowed)])
    assert not is_allowed_host_path(str(outside), [str(allowed)])
    sneaky = allowed / "vol" / ".." / ".." / "etc"
    assert not is_allowed_host_path(str(sneaky), [str(allowed)])
