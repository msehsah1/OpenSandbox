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
	"fmt"
	"strconv"
)

const maxExpiresB36Len = 13

func ParseExpiresB36(s string) (uint64, error) {
	if s == "" {
		return 0, errors.New("empty expires_b36")
	}
	if len(s) > maxExpiresB36Len {
		return 0, errors.New("expires_b36 too long")
	}
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			return 0, errors.New("expires_b36 must be lowercase [0-9a-z]")
		}
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'z')) {
			return 0, errors.New("invalid character in expires_b36")
		}
	}
	if len(s) > 1 && s[0] == '0' {
		return 0, errors.New("expires_b36 must not have leading zeros")
	}
	return strconv.ParseUint(s, 36, 64)
}

func FormatExpiresB36(sec uint64) string {
	return strconv.FormatUint(sec, 36)
}

func ValidateSignatureFormat(signature string) error {
	if len(signature) != 9 {
		return fmt.Errorf("signature must be 9 characters, got %d", len(signature))
	}
	for i := 0; i < 8; i++ {
		c := signature[i]
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f')) {
			return errors.New("signature hex8 must be lowercase hex")
		}
	}
	c := signature[8]
	if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'z')) {
		return errors.New("signed_key_id must be [0-9a-z]")
	}
	return nil
}
