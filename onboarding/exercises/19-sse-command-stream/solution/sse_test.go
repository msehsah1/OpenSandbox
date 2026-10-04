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
	"strings"
	"testing"
)

func TestStream(t *testing.T) {
	var b strings.Builder
	if err := StreamCommand(&b, []string{"hello"}, 0); err != nil {
		t.Fatal(err)
	}
	got := b.String()
	if !strings.Contains(got, "event: stdout\ndata: hello\n") {
		t.Fatalf("stdout missing: %q", got)
	}
	if !strings.Contains(got, `event: execution_complete`) || !strings.Contains(got, `"exit_code":0`) {
		t.Fatalf("complete missing: %q", got)
	}
}
