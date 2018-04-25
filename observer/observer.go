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

	"github.com/fsnotify/fsnotify"
)

// Listener is the function type to run on events.
type Listener func(interface{})

// Observer emplements the observer pattern
type Observer struct {
	quit      chan bool
	events    chan interface{}
	watcher   *fsnotify.Watcher
	listeners []Listener
}

// Open the observer channles and run observer loop
func (o *Observer) Open() error {
	if o.events != nil {
		return fmt.Errorf("Observer already inititated.")
	}

	// Create the observer channels
	o.quit = make(chan bool)
	o.events = make(chan interface{})

	// Init an empty listeners list
	o.listeners = make([]Listener, 0)

	// Run the observer
	return o.Run()
}

// Close the observer channles
func (o *Observer) Close() error {
	// Send a quit signal
	o.quit <- true

	// Close channels
	close(o.quit)
	close(o.events)

	// Close file watcher
	if o.watcher == nil {
		o.watcher.Close()
	}

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

// AddListener adds a listener function to run on event.
func (o *Observer) AddListener(l Listener) error {
	o.listeners = append(o.listeners, l)

	return nil
}

// Emit an event
func (o *Observer) Emit(event interface{}) error {
	o.events <- event

	return nil
}

// RunWatch run a Watcher for file changes
func (o *Observer) RunWatch() error {
	var err error

	o.watcher, err = fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	// Listen for file changes
	go func() {
		for {
			select {
			case event := <-o.watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					o.Emit(event)
				}
			case err := <-o.watcher.Errors:
				if err != nil {
					o.Emit(err)
				}
			}
		}
	}()

	return nil
}

// Watch for file changes
func (o *Observer) Watch(files []string) error {
	// Init watcher on first call
	if o.watcher == nil {
		err := o.RunWatch()
		if err != nil {
			return err
		}
	}

	// Add files to watch
	for _, f := range files {
		err := o.watcher.Add(f)
		if err != nil {
			return err
		}
	}

	return nil
}
