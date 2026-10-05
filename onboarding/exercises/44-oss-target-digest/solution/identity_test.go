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

func TestCanonical(t *testing.T) {
	got, err := CanonicalOSSEndpoint("HTTPS://OSS.Example.com:443")
	if err != nil || got != "https://oss.example.com" {
		t.Fatalf("%s %v", got, err)
	}
	got, err = CanonicalOSSEndpoint("https://oss.example.com:8443")
	if err != nil || got != "https://oss.example.com:8443" {
		t.Fatalf("%s %v", got, err)
	}
	if _, err := CanonicalOSSEndpoint("http://oss.example.com"); err == nil {
		t.Fatal("http")
	}
	if _, err := CanonicalOSSEndpoint("https://oss.example.com/bucket"); err == nil {
		t.Fatal("path")
	}
}

func TestDigestStable(t *testing.T) {
	a := TargetDigest("v1\x00", "oss", "https://oss.example.com", "bkt")
	b := TargetDigest("v1\x00", "oss", "https://oss.example.com", "bkt")
	if a != b || !hasPrefix(a, "sha256:") || len(a) != 7+64 {
		t.Fatalf("%s", a)
	}
	c := TargetDigest("v1\x00", "oss", "https://oss.example.com", "other")
	if a == c {
		t.Fatal("collision")
	}
}

func hasPrefix(s, p string) bool {
	return len(s) >= len(p) && s[:len(p)] == p
}
