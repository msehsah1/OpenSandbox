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
	"testing"
	"time"
)

func mid() float64 { return 0.5 }

func TestDoublesAndClamps(t *testing.T) {
	min := 100 * time.Millisecond
	max := 800 * time.Millisecond
	d := NextBackoff(0, min, max, 0.1, mid)
	if d != min {
		t.Fatalf("start %s", d)
	}
	d = NextBackoff(min, min, max, 0.1, mid)
	if d != 200*time.Millisecond {
		t.Fatalf("double %s", d)
	}
	d = NextBackoff(max, min, max, 0.1, mid)
	if d != max {
		t.Fatalf("clamp %s", d)
	}
}
