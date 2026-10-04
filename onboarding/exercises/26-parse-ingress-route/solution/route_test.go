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

func TestParseHeader(t *testing.T) {
	r, err := ParseIngressHeader("https://a0ee-bc99-44772.example")
	if err != nil {
		t.Fatal(err)
	}
	if r.SandboxID != "a0ee-bc99" || r.Port != 44772 {
		t.Fatalf("%+v", r)
	}
}

func TestParseURI(t *testing.T) {
	r, err := ParseIngressURI("/sbx-1/8080/v1/ping")
	if err != nil {
		t.Fatal(err)
	}
	if r.SandboxID != "sbx-1" || r.Port != 8080 || r.Path != "/v1/ping" {
		t.Fatalf("%+v", r)
	}
}

func TestRejects(t *testing.T) {
	if _, err := ParseIngressHeader("noport"); err == nil {
		t.Fatal("expected error")
	}
	if _, err := ParseIngressURI("/onlyid"); err == nil {
		t.Fatal("expected error")
	}
}
