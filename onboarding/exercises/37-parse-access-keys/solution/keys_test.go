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
	"bytes"
	"testing"
)

func TestParseKeys(t *testing.T) {
	got, err := ParseKeys("a=c2VjcmV0, b=dG9rZW4=")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got["a"], []byte("secret")) || !bytes.Equal(got["b"], []byte("token")) {
		t.Fatalf("%v", got)
	}
}

func TestParseKeysRejects(t *testing.T) {
	for _, s := range []string{"", "   ,  ", "A=c2VjcmV0", "ab=c2VjcmV0", "a=", "a=@@@", "noequals"} {
		if _, err := ParseKeys(s); err == nil {
			t.Fatalf("expected error for %q", s)
		}
	}
}
