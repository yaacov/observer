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
	o.Run()

	// This event will not run any listener
	o.Emit("Hello")

	// Add a listner that logs events
	o.AddListener(func(e interface{}) {
		log.Printf("Recived: %s.\n", e.(string))
	})

	// This events will be loged
	time.Sleep(2 * time.Second)
	o.Emit("Holla")
	time.Sleep(2 * time.Second)
	o.Emit("Hy")

	// Close observer
	time.Sleep(3 * time.Second)
	o.Close()
}
