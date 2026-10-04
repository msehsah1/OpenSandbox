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

import argparse
from typing import Protocol


class SandboxPort(Protocol):
    def create(self, image: str) -> str: ...
    def run(self, sandbox_id: str, command: list[str]) -> tuple[str, int]: ...


def main(argv: list[str], sandbox: SandboxPort) -> int:
    parser = argparse.ArgumentParser(prog="osb-lite")
    sub = parser.add_subparsers(dest="command")

    create = sub.add_parser("create")
    create.add_argument("--image", required=True)

    run = sub.add_parser("run")
    run.add_argument("sandbox_id")
    run.add_argument("cmd", nargs=argparse.REMAINDER)

    args = parser.parse_args(argv)
    if args.command == "create":
        print(sandbox.create(args.image))
        return 0
    if args.command == "run":
        command = list(args.cmd)
        if command and command[0] == "--":
            command = command[1:]
        stdout, code = sandbox.run(args.sandbox_id, command)
        print(stdout, end="" if stdout.endswith("\n") or stdout == "" else "\n")
        return code
    parser.print_help()
    return 2
