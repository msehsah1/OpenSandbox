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

import json


def parse_sse(payload: str) -> list[dict[str, str]]:
    events: list[dict[str, str]] = []
    for block in payload.split("\n\n"):
        event = "message"
        data_lines: list[str] = []
        for line in block.splitlines():
            if not line or line.startswith(":"):
                continue
            if line.startswith("event:"):
                event = line[6:].strip()
            elif line.startswith("data:"):
                data_lines.append(line[5:].lstrip())
        if data_lines or event != "message":
            events.append({"event": event, "data": "\n".join(data_lines)})
    return events


def execution_from_events(events: list[dict[str, str]]) -> dict:
    stdout: list[str] = []
    exit_code = 0
    for event in events:
        if event["event"] == "stdout":
            stdout.append(event["data"])
        elif event["event"] == "execution_complete" and event["data"]:
            body = json.loads(event["data"])
            exit_code = int(body.get("exit_code", 0))
    return {"stdout": stdout, "exit_code": exit_code}
