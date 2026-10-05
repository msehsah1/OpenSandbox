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

from collections.abc import Callable

from fastapi import FastAPI, Request, Response
from fastapi.responses import JSONResponse
from starlette.middleware.base import BaseHTTPMiddleware

SANDBOX_API_KEY_HEADER = "OPEN-SANDBOX-API-KEY"
EXEMPT_PATHS = ("/health",)


class AuthMiddleware(BaseHTTPMiddleware):
    def __init__(self, app, api_key: str | None = None) -> None:
        super().__init__(app)
        key = (api_key or "").strip()
        self.valid_keys = {key} if key else set()

    async def dispatch(self, request: Request, call_next: Callable) -> Response:
        if any(request.url.path.startswith(p) for p in EXEMPT_PATHS):
            return await call_next(request)
        if not self.valid_keys:
            return await call_next(request)
        provided = request.headers.get(SANDBOX_API_KEY_HEADER)
        if not provided:
            return JSONResponse(
                status_code=401,
                content={
                    "code": "MISSING_API_KEY",
                    "message": f"Provide API key via {SANDBOX_API_KEY_HEADER} header.",
                },
            )
        if provided not in self.valid_keys:
            return JSONResponse(
                status_code=401,
                content={"code": "INVALID_API_KEY", "message": "Invalid API key."},
            )
        return await call_next(request)


def create_app(api_key: str | None = "secret") -> FastAPI:
    app = FastAPI()
    app.add_middleware(AuthMiddleware, api_key=api_key)

    @app.get("/health")
    def health() -> dict[str, str]:
        return {"status": "healthy"}

    @app.get("/v1/sandboxes")
    def list_sandboxes() -> dict[str, list]:
        return {"items": []}

    return app
