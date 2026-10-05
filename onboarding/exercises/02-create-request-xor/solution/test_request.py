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

import pytest
from pydantic import ValidationError

from request import CreateSandboxRequest


def test_image_only() -> None:
    req = CreateSandboxRequest(image="alpine")
    assert req.image == "alpine"
    assert req.snapshot_id is None


def test_snapshot_only() -> None:
    req = CreateSandboxRequest(snapshot_id="snap-1")
    assert req.snapshot_id == "snap-1"


@pytest.mark.parametrize(
    "kwargs",
    [{}, {"image": "alpine", "snapshot_id": "snap-1"}, {"image": "  ", "snapshot_id": ""}],
)
def test_rejects_not_exactly_one(kwargs: dict) -> None:
    with pytest.raises(ValidationError, match="Exactly one"):
        CreateSandboxRequest(**kwargs)
