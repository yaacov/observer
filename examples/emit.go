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

	// Add a listener that logs events
	o.AddListener(func(e interface{}) {
		log.Printf("Received: %s.\n", e.(string))
	})

	// This event will be picked by the listener
	go func() {
		time.Sleep(2 * time.Second)
		o.Emit("Holla")
	}()

	// This event will be picked by the listener
	go func() {
		time.Sleep(1 * time.Second)
		o.Emit("Hello")
	}()

	// Close observer
	time.Sleep(3 * time.Second)
}
