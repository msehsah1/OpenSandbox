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

import asyncio

import pytest

from gather import gather_fail_fast


@pytest.mark.asyncio
async def test_success() -> None:
    async def one() -> int:
        return 1

    async def two() -> int:
        return 2

    assert await gather_fail_fast(one(), two()) == [1, 2]


@pytest.mark.asyncio
async def test_cancels_sibling() -> None:
    cancelled = asyncio.Event()

    async def boom() -> None:
        raise RuntimeError("401")

    async def slow() -> None:
        try:
            await asyncio.sleep(10)
        except asyncio.CancelledError:
            cancelled.set()
            raise

    with pytest.raises(RuntimeError, match="401"):
        await gather_fail_fast(boom(), slow())
    assert cancelled.is_set()
