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

import "net/netip"

func ClassifyTarget(rule string) string {
	if _, err := netip.ParseAddr(rule); err == nil {
		return "ip"
	}
	if _, err := netip.ParsePrefix(rule); err == nil {
		return "cidr"
	}
	if rule == "" {
		return "invalid"
	}
	if stringsLooksDomain(rule) {
		return "domain"
	}
	return "invalid"
}

func stringsLooksDomain(rule string) bool {
	for _, c := range rule {
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '.' || c == '*' || c == '-' {
			continue
		}
		if c >= '0' && c <= '9' {
			continue
		}
		return false
	}
	return rule != ""
}

func MatchIP(addr, rule string) bool {
	ip, err := netip.ParseAddr(addr)
	if err != nil {
		return false
	}
	if exact, err := netip.ParseAddr(rule); err == nil {
		return ip == exact
	}
	if prefix, err := netip.ParsePrefix(rule); err == nil {
		return prefix.Contains(ip)
	}
	return false
}
