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
	"fmt"
	"strconv"
	"strings"
)

type ByteRange struct {
	Start  int64
	Length int64
}

func ParseRange(header string, size int64) (ByteRange, error) {
	if size <= 0 {
		return ByteRange{}, fmt.Errorf("invalid size")
	}
	if !strings.HasPrefix(header, "bytes=") {
		return ByteRange{}, fmt.Errorf("invalid range")
	}
	ra := strings.TrimSpace(strings.TrimPrefix(header, "bytes="))
	i := strings.Index(ra, "-")
	if i < 0 {
		return ByteRange{}, fmt.Errorf("invalid range")
	}
	start, end := strings.TrimSpace(ra[:i]), strings.TrimSpace(ra[i+1:])
	if start == "" {
		n, err := strconv.ParseInt(end, 10, 64)
		if err != nil || n < 0 {
			return ByteRange{}, fmt.Errorf("invalid range")
		}
		if n > size {
			n = size
		}
		return ByteRange{Start: size - n, Length: n}, nil
	}
	s, err := strconv.ParseInt(start, 10, 64)
	if err != nil || s < 0 || s >= size {
		return ByteRange{}, fmt.Errorf("invalid range")
	}
	if end == "" {
		return ByteRange{Start: s, Length: size - s}, nil
	}
	e, err := strconv.ParseInt(end, 10, 64)
	if err != nil || e < s {
		return ByteRange{}, fmt.Errorf("invalid range")
	}
	if e > size-1 {
		e = size - 1
	}
	return ByteRange{Start: s, Length: e - s + 1}, nil
}
