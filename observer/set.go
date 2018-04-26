// Copyright 2018 Yaacov Zamir <kobi.zamir@gmail.com>
// and other contributors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package observer for events
package observer

import (
	"fmt"
)

// Set implements an array of unique strings
type Set struct {
	set map[string]struct{}
}

// Add a new string value to the set
func (s *Set) Add(v string) error {
	// Check for value exist
	if _, ok := s.set[v]; ok {
		return fmt.Errorf("Value already in set.")
	}

	// Check for empty set
	if s.set == nil {
		s.set = make(map[string]struct{})
	}

	s.set[v] = struct{}{}
	return nil
}

// Clear all values in the set
func (s *Set) Clear() {
	s.set = nil
}

// Get set value as array of strings
func (s Set) Get() (keys []string) {
	keys = make([]string, 0, len(s.set))
	for k := range s.set {
		keys = append(keys, k)
	}

	return
}

// Has check if set has a string
func (s Set) Has(v string) (ok bool) {
	_, ok = s.set[v]

	return
}
