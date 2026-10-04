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
	"strconv"
	"strings"
)

type Route struct {
	SandboxID string
	Port      int
	Path      string
}

func ParseIngressHeader(value string) (Route, error) {
	trimmed := strings.TrimPrefix(strings.TrimPrefix(value, "https://"), "http://")
	label := strings.Split(trimmed, ".")[0]
	parts := strings.Split(label, "-")
	if len(parts) <= 1 || parts[0] == "" {
		return Route{}, fmt.Errorf("invalid ingress target: %s", value)
	}
	port, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil || port <= 0 || port > 65535 {
		return Route{}, fmt.Errorf("invalid port: %s", value)
	}
	return Route{SandboxID: strings.Join(parts[:len(parts)-1], "-"), Port: port}, nil
}

func ParseIngressURI(path string) (Route, error) {
	trimmed := strings.TrimPrefix(path, "/")
	if trimmed == "" {
		return Route{}, fmt.Errorf("invalid uri: %s", path)
	}
	parts := strings.Split(trimmed, "/")
	if len(parts) < 2 || parts[0] == "" {
		return Route{}, fmt.Errorf("invalid uri: %s", path)
	}
	port, err := strconv.Atoi(parts[1])
	if err != nil || port <= 0 || port > 65535 {
		return Route{}, fmt.Errorf("invalid port: %s", path)
	}
	rest := "/"
	if len(parts) > 2 {
		rest = "/" + strings.Join(parts[2:], "/")
	}
	return Route{SandboxID: parts[0], Port: port, Path: rest}, nil
}
