# Observer

[![Build Status](https://travis-ci.org/yaacov/observer.svg?branch=master)](https://travis-ci.orgyaacov/observer)

## Description

Yet another go observer / listener.

## Example

```
import (
	"log"
	"time"

	"github.com/yaacov/observer/observer"
)

main() {
  // Open an observer and start running
  o := observer.Observer{}
  o.Open()

  // Add a listener that logs events
  o.AddListener(func(e interface{}) {
    log.Printf("Received: %s.\n", e.(string))
  })

  // This event will be picked by the listener
  o.Emit("Hello")

  // Close observer
  time.Sleep(3 * time.Second)
  o.Close()
}
```
