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
	"sync"
)

type record struct {
	running  bool
	exitCode int
}

type Store struct {
	mu   sync.Mutex
	next int
	m    map[string]record
}

func NewStore() *Store {
	return &Store{m: make(map[string]record)}
}

func (s *Store) Start(run func() int) string {
	s.mu.Lock()
	s.next++
	id := fmt.Sprintf("cmd-%d", s.next)
	s.m[id] = record{running: true}
	s.mu.Unlock()
	go func() {
		code := run()
		s.mu.Lock()
		s.m[id] = record{running: false, exitCode: code}
		s.mu.Unlock()
	}()
	return id
}

func (s *Store) Status(id string) (running bool, exitCode int, ok bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	rec, ok := s.m[id]
	return rec.running, rec.exitCode, ok
}
