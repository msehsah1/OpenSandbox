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

from launch import launch


@pytest.mark.asyncio
async def test_success_does_not_kill() -> None:
    killed: list[str] = []

    async def create() -> dict:
        return {"id": "sbx-1"}

    async def health(sandbox_id: str) -> None:
        assert sandbox_id == "sbx-1"

    async def kill(sandbox_id: str) -> None:
        killed.append(sandbox_id)

    assert await launch(create, health, kill) == "sbx-1"
    assert killed == []


@pytest.mark.asyncio
async def test_health_failure_kills() -> None:
    killed: list[str] = []

    async def create() -> dict:
        return {"id": "sbx-2"}

    async def health(sandbox_id: str) -> None:
        raise TimeoutError("execd not ready")

    async def kill(sandbox_id: str) -> None:
        killed.append(sandbox_id)

    with pytest.raises(TimeoutError):
        await launch(create, health, kill)
    assert killed == ["sbx-2"]
