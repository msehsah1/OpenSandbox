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

from fastapi.testclient import TestClient

from app import SANDBOX_API_KEY_HEADER, create_app


def test_health_is_open() -> None:
    client = TestClient(create_app("secret"))
    assert client.get("/health").status_code == 200


def test_missing_and_wrong_key() -> None:
    client = TestClient(create_app("secret"))
    missing = client.get("/v1/sandboxes")
    assert missing.status_code == 401
    assert missing.json()["code"] == "MISSING_API_KEY"
    wrong = client.get("/v1/sandboxes", headers={SANDBOX_API_KEY_HEADER: "nope"})
    assert wrong.json()["code"] == "INVALID_API_KEY"


def test_valid_key_and_open_mode() -> None:
    locked = TestClient(create_app("secret"))
    ok = locked.get("/v1/sandboxes", headers={SANDBOX_API_KEY_HEADER: "secret"})
    assert ok.status_code == 200
    opened = TestClient(create_app(""))
    assert opened.get("/v1/sandboxes").status_code == 200
