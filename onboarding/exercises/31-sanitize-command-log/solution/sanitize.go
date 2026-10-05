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
	"regexp"
)

var (
	kvPattern     = regexp.MustCompile(`(?i)\b(password|token)=(\S+)`)
	bearerPattern = regexp.MustCompile(`(?i)(Authorization:\s*Bearer\s+)(\S+)`)
)

func MaskToken(token string) string {
	if len(token) <= 8 {
		return "****"
	}
	return token[:4] + "****" + token[len(token)-4:]
}

func SanitizeCommand(cmd string) string {
	cmd = kvPattern.ReplaceAllStringFunc(cmd, func(m string) string {
		parts := kvPattern.FindStringSubmatch(m)
		return parts[1] + "=" + MaskToken(parts[2])
	})
	cmd = bearerPattern.ReplaceAllStringFunc(cmd, func(m string) string {
		parts := bearerPattern.FindStringSubmatch(m)
		return parts[1] + MaskToken(parts[2])
	})
	return cmd
}
