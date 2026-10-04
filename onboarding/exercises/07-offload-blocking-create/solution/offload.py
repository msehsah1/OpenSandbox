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

import asyncio
from collections.abc import Callable
from threading import Thread
from typing import Any, TypeVar

T = TypeVar("T")


async def create_sandbox(provision: Callable[..., T], *args: Any) -> T:
    loop = asyncio.get_running_loop()
    future: asyncio.Future[T] = loop.create_future()

    def _run() -> None:
        try:
            result = provision(*args)
            loop.call_soon_threadsafe(future.set_result, result)
        except BaseException as exc:
            try:
                loop.call_soon_threadsafe(future.set_exception, exc)
            except RuntimeError:
                pass

    Thread(target=_run, daemon=True).start()
    return await future
