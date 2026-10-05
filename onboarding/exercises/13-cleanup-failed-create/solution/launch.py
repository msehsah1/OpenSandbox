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

from collections.abc import Awaitable, Callable
from typing import Any


async def launch(
    create: Callable[[], Awaitable[dict[str, Any]]],
    health: Callable[[str], Awaitable[None]],
    kill: Callable[[str], Awaitable[None]],
) -> str:
    sandbox_id: str | None = None
    try:
        created = await create()
        sandbox_id = created["id"]
        await health(sandbox_id)
        return sandbox_id
    except BaseException:
        if sandbox_id is not None:
            try:
                await kill(sandbox_id)
            except Exception:
                pass
        raise
