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

// Package main
package main

import (
	"log"
	"time"

	"github.com/yaacov/observer/observer"
)

func main() {
	log.Print("Observer Example\n")

	// Open observer and start running
	o := observer.Observer{}
	o.Open()
	defer o.Close()

	// Watch for changes in LICENSE file
	err := o.Watch([]string{"LICENSE"})
	if err != nil {
		log.Fatal("Error: ", err)
	}
	log.Print("Observer is watching LICENSE file, try to change it.\n")

	// Add a listener that logs events
	o.AddListener(func(e interface{}) {
		log.Printf("Received: %v.\n", e)
	})

	// This events will be loged
	go func() {
		time.Sleep(2 * time.Second)
		o.Emit("Holla")
	}()

	go func() {
		time.Sleep(1 * time.Second)
		o.Emit("Hy")
	}()

	// Wait for events
	time.Sleep(10 * time.Second)
}
