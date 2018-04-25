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

// Package server API REST server
package observer

import (
	"fmt"
)

type Listener func(interface{})

type Observer struct {
	quit      chan bool
	events    chan interface{}
	listeners []Listener
}

// Open the observer channles
func (o *Observer) Open() error {
	if o.events != nil {
		return fmt.Errorf("Observer already inititated.")
	}

	// Create the observer channels
	o.quit = make(chan bool)
	o.events = make(chan interface{})

	// Init an empty listeners list
	o.listeners = make([]Listener, 0)

	return nil
}

// Close the observer channles
func (o *Observer) Close() error {
	// Send a quit signal
	o.quit <- true

	// Close channels
	close(o.quit)
	close(o.events)

	return nil
}

// Run the observer main loop
func (o *Observer) Run() error {
	// Run observer
	go func() {
		for {
			select {
			case event := <-o.events:
				// Run all listeners for this event
				for _, listener := range o.listeners {
					go listener(event)
				}
			case <-o.quit:
				return
			}
		}
	}()

	return nil
}

// Emit an event
func (o *Observer) AddListener(l Listener) error {
	o.listeners = append(o.listeners, l)

	return nil
}

// Emit an event
func (o *Observer) Emit(event interface{}) error {
	o.events <- event

	return nil
}
