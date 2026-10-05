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

func TestOpenEnded(t *testing.T) {
	r, err := ParseRange("bytes=2-", 10)
	if err != nil {
		t.Fatal(err)
	}
	if r.Start != 2 || r.Length != 8 {
		t.Fatalf("%+v", r)
	}
}

func TestClampHugeEnd(t *testing.T) {
	r, err := ParseRange("bytes=0-999999999999", 10)
	if err != nil {
		t.Fatal(err)
	}
	if r.Start != 0 || r.Length != 10 {
		t.Fatalf("clamp failed: %+v", r)
	}
}

func TestSuffix(t *testing.T) {
	r, err := ParseRange("bytes=-3", 10)
	if err != nil {
		t.Fatal(err)
	}
	if r.Start != 7 || r.Length != 3 {
		t.Fatalf("%+v", r)
	}
}
