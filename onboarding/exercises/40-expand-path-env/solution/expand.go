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
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

var envVarPattern = regexp.MustCompile(`\$\{([A-Za-z_][A-Za-z0-9_]*)\}|\$([A-Za-z_][A-Za-z0-9_]*)`)

func ExpandPath(path string, env map[string]string) (string, error) {
	if path == "" {
		return "", nil
	}
	if env == nil {
		env = map[string]string{}
	}
	matches := envVarPattern.FindAllStringSubmatch(path, -1)
	missingSet := map[string]struct{}{}
	for _, m := range matches {
		name := m[1]
		if name == "" {
			name = m[2]
		}
		if _, ok := env[name]; !ok {
			missingSet[name] = struct{}{}
		}
	}
	if len(missingSet) > 0 {
		missing := make([]string, 0, len(missingSet))
		for name := range missingSet {
			missing = append(missing, name)
		}
		sort.Strings(missing)
		return "", fmt.Errorf("path references undefined environment variables: %s", strings.Join(missing, ","))
	}
	expanded := os.Expand(path, func(key string) string {
		return env[key]
	})
	if expanded == "~" || strings.HasPrefix(expanded, "~/") {
		home, ok := env["HOME"]
		if !ok || home == "" {
			return "", fmt.Errorf("HOME is not set")
		}
		if expanded == "~" {
			return home, nil
		}
		return filepath.Join(home, expanded[2:]), nil
	}
	return expanded, nil
}
