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


def route_create(template_id: str | None, snapshot_backend: str | None) -> str:
    if (template_id or "").strip():
        return "fsb"
    if (snapshot_backend or "").strip() == "fsb":
        return "fsb"
    return "kubernetes"


def route_existing(sandbox_id: str) -> str:
    return "fsb" if sandbox_id.startswith("fsb-") else "kubernetes"
