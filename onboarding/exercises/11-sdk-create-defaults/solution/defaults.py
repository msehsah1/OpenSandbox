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

from typing import Any


DEFAULT_ENTRYPOINT = ["tail", "-f", "/dev/null"]
DEFAULT_RESOURCE = {"cpu": "1", "memory": "2Gi"}


def prepare_create(
    *,
    image: str | dict[str, str] | None = None,
    snapshot_id: str | None = None,
    entrypoint: list[str] | None = None,
    resource: dict[str, str] | None = None,
) -> dict[str, Any]:
    if (image is None) == (snapshot_id is None):
        raise ValueError("Exactly one of image or snapshot_id must be specified")
    spec: dict[str, str] | None
    if isinstance(image, str):
        spec = {"image": image}
    else:
        spec = image
    return {
        "spec": spec,
        "snapshot_id": snapshot_id,
        "entrypoint": entrypoint or list(DEFAULT_ENTRYPOINT),
        "resource": resource or dict(DEFAULT_RESOURCE),
    }
