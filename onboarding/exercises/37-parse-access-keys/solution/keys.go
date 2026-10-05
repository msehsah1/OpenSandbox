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
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
)

func ParseKeys(s string) (map[string][]byte, error) {
	if strings.TrimSpace(s) == "" {
		return nil, errors.New("empty keys string")
	}
	out := make(map[string][]byte)
	for _, seg := range strings.Split(s, ",") {
		seg = strings.TrimSpace(seg)
		if seg == "" {
			continue
		}
		key, val, ok := strings.Cut(seg, "=")
		if !ok || key == "" || val == "" {
			return nil, fmt.Errorf("invalid keys segment %q (want key_id=base64)", seg)
		}
		key = strings.TrimSpace(key)
		val = strings.TrimSpace(val)
		if len(key) != 1 {
			return nil, fmt.Errorf("key_id must be exactly 1 character, got %q", key)
		}
		r := key[0]
		if r >= 'A' && r <= 'Z' {
			return nil, fmt.Errorf("key_id must be lowercase [0-9a-z], got %q", key)
		}
		if !((r >= '0' && r <= '9') || (r >= 'a' && r <= 'z')) {
			return nil, fmt.Errorf("key_id must be [0-9a-z], got %q", key)
		}
		raw, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			return nil, fmt.Errorf("decode secret for key %q: %w", key, err)
		}
		if len(raw) == 0 {
			return nil, fmt.Errorf("empty secret for key %q", key)
		}
		out[key] = raw
	}
	if len(out) == 0 {
		return nil, errors.New("no keys parsed")
	}
	return out, nil
}
