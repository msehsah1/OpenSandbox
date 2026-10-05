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

func NextBackoff(prev, min, max time.Duration, jitter float64, rng func() float64) time.Duration {
	var d time.Duration
	if prev <= 0 {
		d = min
	} else {
		d = prev * 2
		if d < min {
			d = min
		}
		if d > max {
			d = max
		}
	}
	if jitter > 0 && rng != nil {
		span := float64(d) * jitter
		delta := time.Duration((rng()*2 - 1) * span)
		d += delta
	}
	if d < min {
		d = min
	}
	if d > max {
		d = max
	}
	if d < time.Nanosecond {
		d = time.Nanosecond
	}
	return d
}
