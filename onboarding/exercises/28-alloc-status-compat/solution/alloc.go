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

import "encoding/json"

type AllocStatus struct {
	Pods       []string `json:"pods"`
	PoolRef    string   `json:"poolRef,omitempty"`
	Generation int64    `json:"generation,omitempty"`
}

func ParseAllocStatus(raw string) (AllocStatus, error) {
	var status AllocStatus
	if err := json.Unmarshal([]byte(raw), &status); err != nil {
		return AllocStatus{}, err
	}
	if status.Pods == nil {
		status.Pods = []string{}
	}
	return status, nil
}

func FormatAllocStatus(s AllocStatus) (string, error) {
	if s.Pods == nil {
		s.Pods = []string{}
	}
	raw, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(raw), nil
}
