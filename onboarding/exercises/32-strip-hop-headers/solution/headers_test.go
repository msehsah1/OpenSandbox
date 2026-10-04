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

func TestStrip(t *testing.T) {
	in := map[string][]string{
		"Authorization":          {"Bearer x"},
		"X-Request-ID":           {"rid"},
		"Transfer-Encoding":      {"chunked"},
		"OpenSandbox-Ingress-To": {"sbx-1-44772"},
	}
	out := StripForwardHeaders(in)
	if _, ok := out["Transfer-Encoding"]; ok {
		t.Fatal("hop header leaked")
	}
	if _, ok := out["OpenSandbox-Ingress-To"]; ok {
		t.Fatal("routing header leaked")
	}
	if out["Authorization"][0] != "Bearer x" || out["X-Request-ID"][0] != "rid" {
		t.Fatalf("%v", out)
	}
	if _, ok := in["Transfer-Encoding"]; !ok {
		t.Fatal("input was mutated")
	}
}
