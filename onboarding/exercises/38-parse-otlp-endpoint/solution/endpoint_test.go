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

func TestParseOTLP(t *testing.T) {
	host, port, ok := ParseOTLPEndpoint("https://otel.example.com/v1/metrics")
	if !ok || host != "otel.example.com" || port != "443" {
		t.Fatalf("%s %s %v", host, port, ok)
	}
	host, port, ok = ParseOTLPEndpoint("http://10.0.0.1:4318")
	if !ok || host != "10.0.0.1" || port != "4318" {
		t.Fatalf("%s %s %v", host, port, ok)
	}
	host, port, ok = ParseOTLPEndpoint("https://otel.example.com.")
	if !ok || host != "otel.example.com" {
		t.Fatalf("root dot: %s %v", host, ok)
	}
}

func TestParseOTLPRejects(t *testing.T) {
	for _, s := range []string{"", "otel.example.com:4318", "://", "https://"} {
		if _, _, ok := ParseOTLPEndpoint(s); ok {
			t.Fatalf("expected !ok for %q", s)
		}
	}
}
