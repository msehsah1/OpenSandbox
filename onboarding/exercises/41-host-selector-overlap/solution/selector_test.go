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

func TestMatches(t *testing.T) {
	s, err := ParseCanonical("*.example.com")
	if err != nil {
		t.Fatal(err)
	}
	if !s.Matches("api.example.com") || s.Matches("example.com") || s.Matches("notexample.com") {
		t.Fatal("wildcard semantics")
	}
}

func TestOverlaps(t *testing.T) {
	a, _ := ParseCanonical("*.example.com")
	b, _ := ParseCanonical("api.example.com")
	if !a.Overlaps(b) {
		t.Fatal("wildcard vs host")
	}
	c, _ := ParseCanonical("*.a.example.com")
	if !a.Overlaps(c) {
		t.Fatal("nested wildcards")
	}
	d, _ := ParseCanonical("foo.com")
	e, _ := ParseCanonical("bar.com")
	if d.Overlaps(e) {
		t.Fatal("disjoint")
	}
}

func TestRejects(t *testing.T) {
	for _, s := range []string{"localhost", "1.2.3.4", "EXAMPLE.com", "-bad.com"} {
		if _, err := ParseCanonical(s); err == nil {
			t.Fatalf("expected error for %q", s)
		}
	}
}
