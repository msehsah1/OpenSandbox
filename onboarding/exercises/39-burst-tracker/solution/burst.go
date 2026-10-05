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

import "time"

type BurstTracker struct {
	max    int
	window time.Duration
	now    func() time.Time
	ring   []time.Time
	idx    int
	filled int
}

func NewBurstTracker(max int, window time.Duration, now func() time.Time) *BurstTracker {
	if max < 1 {
		max = 1
	}
	return &BurstTracker{
		max:    max,
		window: window,
		now:    now,
		ring:   make([]time.Time, max),
	}
}

func (b *BurstTracker) Record() {
	b.ring[b.idx] = b.now()
	b.idx = (b.idx + 1) % b.max
	if b.filled < b.max {
		b.filled++
	}
}

func (b *BurstTracker) Exceeded() bool {
	if b.filled < b.max {
		return false
	}
	oldest := b.ring[b.idx]
	return b.now().Sub(oldest) <= b.window
}
