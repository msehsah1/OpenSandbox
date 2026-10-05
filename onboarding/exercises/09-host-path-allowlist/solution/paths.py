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

from pathlib import Path


def is_allowed_host_path(path: str, allowed: list[str]) -> bool:
    if not allowed:
        return True
    resolved = Path(path).resolve()
    for raw_prefix in allowed:
        prefix = Path(raw_prefix).resolve()
        if resolved == prefix:
            return True
        try:
            resolved.relative_to(prefix)
            return True
        except ValueError:
            continue
    return False
