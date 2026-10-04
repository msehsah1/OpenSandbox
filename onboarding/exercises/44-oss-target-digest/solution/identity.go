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
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net"
	"net/url"
	"strings"
)

func CanonicalOSSEndpoint(raw string) (string, error) {
	u, err := url.Parse(raw)
	if err != nil || !strings.EqualFold(u.Scheme, "https") || u.Host == "" || u.User != nil || u.RawQuery != "" || u.Fragment != "" || (u.Path != "" && u.Path != "/") {
		return "", errors.New("OSS endpoint must be an HTTPS origin")
	}
	host := strings.ToLower(u.Hostname())
	if host == "" {
		return "", errors.New("OSS endpoint must be an HTTPS origin")
	}
	if port := u.Port(); port != "" && port != "443" {
		host = net.JoinHostPort(host, port)
	} else if strings.Contains(host, ":") {
		host = "[" + host + "]"
	}
	return "https://" + host, nil
}

func TargetDigest(domain string, parts ...string) string {
	h := sha256.New()
	_, _ = h.Write([]byte(domain))
	for _, part := range parts {
		_, _ = fmt.Fprintf(h, "%d:%s", len([]byte(part)), part)
	}
	return "sha256:" + hex.EncodeToString(h.Sum(nil))
}
