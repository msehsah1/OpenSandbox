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

import (
	"errors"
	"net/netip"
	"strings"
)

type Selector struct {
	base     string
	wildcard bool
}

var errInvalid = errors.New("invalid host selector")

func ParseCanonical(text string) (Selector, error) {
	wildcard := strings.HasPrefix(text, "*.")
	base := text
	if wildcard {
		base = strings.TrimPrefix(text, "*.")
	}
	if !validHost(base) {
		return Selector{}, errInvalid
	}
	return Selector{base: base, wildcard: wildcard}, nil
}

func validHost(host string) bool {
	if len(host) > 253 || !strings.Contains(host, ".") {
		return false
	}
	if _, err := netip.ParseAddr(host); err == nil {
		return false
	}
	for _, label := range strings.Split(host, ".") {
		if len(label) == 0 || len(label) > 63 || label[0] == '-' || label[len(label)-1] == '-' {
			return false
		}
		for i := 0; i < len(label); i++ {
			c := label[i]
			if !((c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '-') {
				return false
			}
		}
	}
	return true
}

func (s Selector) String() string {
	if s.wildcard {
		return "*." + s.base
	}
	return s.base
}

func (s Selector) Matches(host string) bool {
	if s.base == "" {
		return false
	}
	host = strings.TrimSuffix(strings.ToLower(host), ".")
	if !validHost(host) {
		return false
	}
	if s.wildcard {
		return strings.HasSuffix(host, "."+s.base)
	}
	return host == s.base
}

func (s Selector) Overlaps(other Selector) bool {
	if s.base == "" || other.base == "" {
		return false
	}
	if !s.wildcard {
		return other.Matches(s.base)
	}
	if !other.wildcard {
		return s.Matches(other.base)
	}
	return s.base == other.base || strings.HasSuffix(s.base, "."+other.base) || strings.HasSuffix(other.base, "."+s.base)
}
