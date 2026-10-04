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

import "net/http"

var hop = map[string]struct{}{
	http.CanonicalHeaderKey("Connection"):             {},
	http.CanonicalHeaderKey("Keep-Alive"):             {},
	http.CanonicalHeaderKey("Proxy-Authenticate"):     {},
	http.CanonicalHeaderKey("Proxy-Authorization"):    {},
	http.CanonicalHeaderKey("TE"):                     {},
	http.CanonicalHeaderKey("Trailer"):                {},
	http.CanonicalHeaderKey("Transfer-Encoding"):      {},
	http.CanonicalHeaderKey("Upgrade"):                {},
	http.CanonicalHeaderKey("Proxy-Connection"):       {},
	http.CanonicalHeaderKey("OpenSandbox-Ingress-To"): {},
	http.CanonicalHeaderKey("OPEN-SANDBOX-INGRESS"):   {},
}

func StripForwardHeaders(h map[string][]string) map[string][]string {
	out := make(map[string][]string, len(h))
	for k, v := range h {
		if _, drop := hop[http.CanonicalHeaderKey(k)]; drop {
			continue
		}
		copied := make([]string, len(v))
		copy(copied, v)
		out[k] = copied
	}
	return out
}
