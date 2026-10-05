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

func TestExpiresRoundTrip(t *testing.T) {
	raw := FormatExpiresB36(36)
	if raw != "10" {
		t.Fatalf("format: %s", raw)
	}
	n, err := ParseExpiresB36(raw)
	if err != nil || n != 36 {
		t.Fatalf("%d %v", n, err)
	}
	if n, err := ParseExpiresB36("0"); err != nil || n != 0 {
		t.Fatalf("zero: %d %v", n, err)
	}
}

func TestExpiresRejects(t *testing.T) {
	for _, s := range []string{"", "00", "K", "10!", "00000000000000"} {
		if _, err := ParseExpiresB36(s); err == nil {
			t.Fatalf("expected error for %q", s)
		}
	}
}

func TestSignatureFormat(t *testing.T) {
	if err := ValidateSignatureFormat("deadbeefa"); err != nil {
		t.Fatal(err)
	}
	for _, s := range []string{"deadbeef", "DEADBEEFa", "deadbeef!"} {
		if err := ValidateSignatureFormat(s); err == nil {
			t.Fatalf("expected error for %q", s)
		}
	}
}
