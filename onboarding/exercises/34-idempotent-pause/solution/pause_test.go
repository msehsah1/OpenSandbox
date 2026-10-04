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

import "testing"

func TestDispatch(t *testing.T) {
	if !ShouldDispatchPause(PhaseRunning) {
		t.Fatal("running")
	}
	if ShouldDispatchPause(PhasePaused) || ShouldDispatchPause(PhasePausing) {
		t.Fatal("already pausing/paused")
	}
}

func TestNext(t *testing.T) {
	next, changed := NextPausePhase(PhasePaused)
	if changed || next != PhasePaused {
		t.Fatalf("%v %v", next, changed)
	}
	next, changed = NextPausePhase(PhaseRunning)
	if !changed || next != PhasePausing {
		t.Fatalf("%v %v", next, changed)
	}
}
