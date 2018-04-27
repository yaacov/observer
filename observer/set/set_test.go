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

// Package set for unique collection of strings.
package set

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var s Set
	var e error

	if len(s.Values()) != 0 {
		t.Error("error init a new Set.")
	}

	e = s.Add("hello")
	if len(s.Values()) != 1 || e != nil {
		t.Error("error adding a new value to Set.")
	}

	e = s.Add("hello")
	if len(s.Values()) != 1 || e == nil {
		t.Error("error adding a recurrent value to Set.")
	}

	e = s.Add("world")
	if len(s.Values()) != 2 || e != nil {
		t.Error("error adding a second value to Set.")
	}
}

func TestClear(t *testing.T) {
	var s Set

	s.Add("hello")
	s.Add("world")
	s.Clear()

	if len(s.Values()) != 0 {
		t.Error("error clearing a Set.")
	}
}

func TestValues(t *testing.T) {
	var s Set

	s.Add("hello")
	s.Add("world")

	if s.Values()[0] != "world" && s.Values()[1] != "world" {
		t.Error("error getting Values from Set.")
	}
}

func TestHas(t *testing.T) {
	var s Set

	s.Add("hello")
	s.Add("world")

	if !s.Has("hello") {
		t.Error("error checking value is in Set.")
	}

	if !s.Has("world") {
		t.Error("error checking value is in Set (2).")
	}

	if s.Has("World") {
		t.Error("error checking value is not in Set.")
	}
}
