// Package main
package main

import (
	"log"
	"time"

	"github.com/yaacov/observer/observer"
)

func main() {
	// Open an observer and start running
	o := observer.Observer{}
	o.Open()
	defer o.Close()

	// Set event damping buffer of 2 sec
	o.SetBufferDuration(2 * time.Second)

	// Add a listener that logs events
	o.AddListener(func(e interface{}) {
		log.Printf("Received: %v.\n", e)
	})

	// This event will be buffered for 2 sec
	go func() {
		o.Emit("Holla")
		o.Emit("Ciao")
	}()

	// This event will be grouped together with previues events
	go func() {
		time.Sleep(1 * time.Second)
		o.Emit("Hello")
		o.Emit("Bonjour")
	}()

	// Close observer
	time.Sleep(3 * time.Second)
}
