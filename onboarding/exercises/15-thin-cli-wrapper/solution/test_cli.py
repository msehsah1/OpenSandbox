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

from cli import main


class Fake:
    def create(self, image: str) -> str:
        assert image == "alpine"
        return "sbx-9"

    def run(self, sandbox_id: str, command: list[str]) -> tuple[str, int]:
        assert sandbox_id == "sbx-9"
        assert command == ["echo", "hi"]
        return ("hi\n", 0)


def test_create_and_run(capsys) -> None:
    fake = Fake()
    assert main(["create", "--image", "alpine"], fake) == 0
    assert capsys.readouterr().out.strip() == "sbx-9"
    assert main(["run", "sbx-9", "--", "echo", "hi"], fake) == 0
    assert capsys.readouterr().out == "hi\n"


def test_unknown() -> None:
    assert main([], Fake()) == 2
