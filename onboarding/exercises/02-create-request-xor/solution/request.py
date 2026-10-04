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

from pydantic import BaseModel, model_validator


class CreateSandboxRequest(BaseModel):
    image: str | None = None
    snapshot_id: str | None = None

    @model_validator(mode="after")
    def exactly_one_startup_source(self) -> "CreateSandboxRequest":
        image = (self.image or "").strip() or None
        snapshot_id = (self.snapshot_id or "").strip() or None
        self.image = image
        self.snapshot_id = snapshot_id
        if (image is None) == (snapshot_id is None):
            raise ValueError("Exactly one of image or snapshot_id must be specified")
        return self
