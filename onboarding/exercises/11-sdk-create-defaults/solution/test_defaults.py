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

from defaults import DEFAULT_ENTRYPOINT, DEFAULT_RESOURCE, prepare_create


def test_image_defaults() -> None:
    body = prepare_create(image="alpine")
    assert body["spec"] == {"image": "alpine"}
    assert body["entrypoint"] == DEFAULT_ENTRYPOINT
    assert body["resource"] == DEFAULT_RESOURCE


def test_xor() -> None:
    with pytest.raises(ValueError, match="Exactly one"):
        prepare_create()
