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

import "testing"

func TestValidateOK(t *testing.T) {
	req := RunCommandRequest{Command: "echo hi", Cwd: "/tmp", TimeoutMs: 1000}
	if err := req.Validate(); err != nil {
		t.Fatal(err)
	}
}

func TestValidateRejects(t *testing.T) {
	cases := []RunCommandRequest{
		{Command: "  "},
		{Command: "echo", TimeoutMs: -1},
		{Command: "echo", Cwd: "tmp"},
	}
	for _, req := range cases {
		if err := req.Validate(); err == nil {
			t.Fatalf("expected error for %+v", req)
		}
	}
}
