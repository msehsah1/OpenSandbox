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

from routing import route_create, route_existing


def test_create_routes() -> None:
    assert route_create("tpl-1", None) == "fsb"
    assert route_create("  ", "fsb") == "fsb"
    assert route_create(None, "pod") == "kubernetes"
    assert route_create(None, None) == "kubernetes"


def test_existing_routes() -> None:
    assert route_existing("fsb-abc") == "fsb"
    assert route_existing("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11") == "kubernetes"
