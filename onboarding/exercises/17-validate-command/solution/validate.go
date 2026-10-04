// Copyright 2026 The OpenSandbox Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package solution

import (
	"fmt"
	"path"
	"strings"
)

type RunCommandRequest struct {
	Command   string
	Cwd       string
	TimeoutMs int
}

func (r RunCommandRequest) Validate() error {
	if strings.TrimSpace(r.Command) == "" {
		return fmt.Errorf("command must not be empty")
	}
	if r.TimeoutMs < 0 {
		return fmt.Errorf("timeoutMs must be >= 0")
	}
	if r.Cwd != "" && !path.IsAbs(r.Cwd) {
		return fmt.Errorf("cwd must be an absolute path")
	}
	return nil
}
