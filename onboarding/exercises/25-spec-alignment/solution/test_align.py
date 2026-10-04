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

from align import assert_create_contract

SPEC_OK = {
    "paths": {
        "/v1/sandboxes": {
            "post": {
                "responses": {
                    "202": {
                        "description": (
                            "Sandbox created and provisioned successfully. "
                            "status.state: Running"
                        )
                    }
                }
            }
        }
    }
}


def test_aligned() -> None:
    assert assert_create_contract(SPEC_OK, "Clients poll execd health after create.") == []


def test_detects_stale_docs() -> None:
    problems = assert_create_contract(
        SPEC_OK,
        "Creation is asynchronous from the API perspective.",
    )
    assert any("asynchronous" in p for p in problems)
