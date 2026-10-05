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

from sse import execution_from_events, parse_sse

STREAM = (
    "event: stdout\n"
    "data: hello\n"
    "\n"
    "event: execution_complete\n"
    "data: {\"exit_code\": 0}\n"
    "\n"
)


def test_parse_and_fold() -> None:
    events = parse_sse(STREAM)
    assert [e["event"] for e in events] == ["stdout", "execution_complete"]
    assert execution_from_events(events) == {"stdout": ["hello"], "exit_code": 0}
