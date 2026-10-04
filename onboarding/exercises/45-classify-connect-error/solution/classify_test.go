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
	"context"
	"errors"
	"fmt"
	"net"
	"syscall"
	"testing"
)

type timeoutErr struct{}

func (timeoutErr) Error() string   { return "i/o timeout" }
func (timeoutErr) Timeout() bool   { return true }
func (timeoutErr) Temporary() bool { return true }

func TestClassify(t *testing.T) {
	if ClassifyConnectError(nil) != ResultSuccess {
		t.Fatal("nil")
	}
	if ClassifyConnectError(fmt.Errorf("wrap: %w", syscall.ECONNREFUSED)) != ResultRefused {
		t.Fatal("refused")
	}
	if ClassifyConnectError(fmt.Errorf("wrap: %w", syscall.EHOSTUNREACH)) != ResultUnreachable {
		t.Fatal("unreach")
	}
	if ClassifyConnectError(&net.DNSError{Err: "no such host", Name: "x"}) != ResultDNS {
		t.Fatal("dns")
	}
	if ClassifyConnectError(context.DeadlineExceeded) != ResultTimeout {
		t.Fatal("deadline")
	}
	if ClassifyConnectError(context.Canceled) != ResultCanceled {
		t.Fatal("canceled")
	}
	if ClassifyConnectError(timeoutErr{}) != ResultTimeout {
		t.Fatal("net timeout")
	}
	if ClassifyConnectError(errors.New("boom")) != ResultOther {
		t.Fatal("other")
	}
}
