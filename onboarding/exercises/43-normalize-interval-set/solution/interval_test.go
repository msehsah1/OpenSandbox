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
	"reflect"
	"testing"
)

func TestNormalize(t *testing.T) {
	got, err := NormalizeIntervalSet([]string{"10.0.0.0/8", "10.1.2.3", "10.0.0.0/8", "192.168.1.1"})
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"10.0.0.0/8", "192.168.1.1"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("%v", got)
	}
}

func TestMaskAndReject(t *testing.T) {
	got, err := NormalizeIntervalSet([]string{"10.1.2.3/16"})
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 1 || got[0] != "10.1.0.0/16" {
		t.Fatalf("%v", got)
	}
	if _, err := NormalizeIntervalSet([]string{"not-an-ip"}); err == nil {
		t.Fatal("expected error")
	}
}
