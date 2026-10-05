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

import time

import pytest

from offload import create_sandbox


@pytest.mark.asyncio
async def test_returns_worker_result() -> None:
    def provision(image: str) -> dict:
        time.sleep(0.05)
        return {"id": "sbx-1", "image": image}

    result = await create_sandbox(provision, "alpine")
    assert result == {"id": "sbx-1", "image": "alpine"}


@pytest.mark.asyncio
async def test_propagates_errors() -> None:
    def provision() -> None:
        raise ValueError("pull failed")

    with pytest.raises(ValueError, match="pull failed"):
        await create_sandbox(provision)
