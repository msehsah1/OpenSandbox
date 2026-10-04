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

func TestMatch(t *testing.T) {
	if !MatchDomain("api.example.com", "*.example.com") {
		t.Fatal("wildcard")
	}
	if MatchDomain("notexample.com", "*.example.com") {
		t.Fatal("label escape")
	}
	if MatchDomain("example.com", "*.example.com") {
		t.Fatal("bare host must not match wildcard")
	}
	if !MatchDomain("EXAMPLE.com", "example.com") {
		t.Fatal("case")
	}
}

func TestDecideDenyWins(t *testing.T) {
	if Decide("evil.example.com", []string{"*.example.com"}, []string{"*.example.com"}) != "deny" {
		t.Fatal("deny must win")
	}
	if Decide("api.openai.com", nil, []string{"*.openai.com"}) != "allow" {
		t.Fatal("allow")
	}
	if Decide("example.org", nil, []string{"*.openai.com"}) != "default" {
		t.Fatal("default")
	}
}
