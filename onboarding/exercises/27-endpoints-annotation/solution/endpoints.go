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
	"encoding/json"
	"fmt"
)

const AnnotationEndpoints = "sandbox.opensandbox.io/endpoints"

func GetEndpoints(annotations map[string]string) ([]string, error) {
	if annotations == nil {
		return nil, fmt.Errorf("no annotations")
	}
	raw, ok := annotations[AnnotationEndpoints]
	if !ok || raw == "" {
		return nil, fmt.Errorf("missing %s annotation", AnnotationEndpoints)
	}
	var endpoints []string
	if err := json.Unmarshal([]byte(raw), &endpoints); err != nil {
		return nil, fmt.Errorf("failed to parse endpoints annotation: %w", err)
	}
	if len(endpoints) == 0 {
		return nil, fmt.Errorf("endpoints annotation contains no IPs")
	}
	return endpoints, nil
}

func SetEndpoints(annotations map[string]string, ips []string) error {
	if annotations == nil {
		return fmt.Errorf("annotations map is nil")
	}
	if len(ips) == 0 {
		return fmt.Errorf("endpoints annotation contains no IPs")
	}
	raw, err := json.Marshal(ips)
	if err != nil {
		return err
	}
	annotations[AnnotationEndpoints] = string(raw)
	return nil
}
