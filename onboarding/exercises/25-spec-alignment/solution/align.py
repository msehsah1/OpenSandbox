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


STALE_PHRASES = (
    "returns before the sandbox is running",
    "creation is asynchronous from the api perspective",
)


def _create_description(openapi: dict) -> str:
    paths = openapi.get("paths") or {}
    for path, item in paths.items():
        if path.rstrip("/").endswith("/sandboxes"):
            desc = (
                ((item.get("post") or {}).get("responses") or {})
                .get("202", {})
                .get("description")
                or ""
            )
            return desc
    return ""


def assert_create_contract(openapi: dict, architecture_text: str) -> list[str]:
    problems: list[str] = []
    desc = _create_description(openapi)
    lowered = desc.lower()
    if "provisioned successfully" not in lowered:
        problems.append("202 description must say provisioned successfully")
    if "running" not in lowered:
        problems.append("202 description must mention Running")
    arch = architecture_text.lower()
    for phrase in STALE_PHRASES:
        if phrase in arch:
            problems.append(f"architecture docs still say: {phrase}")
    return problems
