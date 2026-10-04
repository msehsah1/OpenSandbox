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

import "strings"

func MatchDomain(host, rule string) bool {
	host = strings.ToLower(strings.TrimSuffix(host, "."))
	rule = strings.ToLower(strings.TrimSuffix(rule, "."))
	if rule == "*" || host == rule {
		return true
	}
	if strings.HasPrefix(rule, "*.") {
		// "*.example.com" matches "a.example.com" but not "example.com".
		suffix := strings.TrimPrefix(rule, "*")
		return strings.HasSuffix(host, suffix) && host != strings.TrimPrefix(rule, "*.")
	}
	return false
}

func Decide(host string, deny, allow []string) string {
	for _, rule := range deny {
		if MatchDomain(host, rule) {
			return "deny"
		}
	}
	for _, rule := range allow {
		if MatchDomain(host, rule) {
			return "allow"
		}
	}
	return "default"
}
