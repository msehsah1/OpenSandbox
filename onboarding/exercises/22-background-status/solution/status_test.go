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

func TestStatusLifecycle(t *testing.T) {
	s := NewStore()
	started := make(chan struct{})
	release := make(chan struct{})
	id := s.Start(func() int {
		close(started)
		<-release
		return 3
	})
	<-started
	running, _, ok := s.Status(id)
	if !ok || !running {
		t.Fatalf("want running, ok=%v running=%v", ok, running)
	}
	close(release)
	deadline := time.Now().Add(time.Second)
	for {
		running, code, ok := s.Status(id)
		if ok && !running {
			if code != 3 {
				t.Fatalf("exit %d", code)
			}
			break
		}
		if time.Now().After(deadline) {
			t.Fatal("timeout waiting for finish")
		}
		time.Sleep(5 * time.Millisecond)
	}
	_, _, ok = s.Status("missing")
	if ok {
		t.Fatal("missing id should be !ok")
	}
}
