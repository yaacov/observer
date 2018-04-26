// Package main
package main

import (
	"log"
	"time"

	"github.com/yaacov/observer/observer"
)

func main() {
	// Open an observer and start watching for file modifications
	o := observer.Observer{}
	err := o.Watch([]string{"../LICENSE", "../README.md"})
	if err != nil {
		log.Fatal("Error: ", err)
	}
	defer o.Close()

	// Add a listener that logs events
	o.AddListener(func(e interface{}) {
		log.Printf("File modified: %v.\n", e)
	})

	// Wait 10s for changes in file
	log.Print("Observer is watching the LICENSE and README.md files.\n")
	time.Sleep(12 * time.Second)
}
