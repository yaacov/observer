# Observer

[![Build Status](https://travis-ci.org/yaacov/observer.svg?branch=master)](https://travis-ci.org/yaacov/observer)

## Description

Yet another go observer / listener.

## Develop

```
$ make
$ ./obs-example
```

## API

| Method                         | Description                       |
|--------------------------------|-----------------------------------|
| Open()                         | Open the observer channels        |
| Close()                        | Close the observer channels       |
| AddListener(callback Listener) | Add a listener function to run on event |
| Emit(event interface{})        | Emit event                        |
| Watch(files []string)          | Watch for file changes, and emit a file change events |

## Example

``` go
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

## Example watching file

``` go
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
    log.Printf("File modified: %v.\n", e)
  })

  // Watch for changes in LICENSE file
  err := o.Watch([]string{"LICENSE"})
  if err != nil {
    log.Fatal("Error: ", err)
  }
  log.Print("Observer is watching the LICENSE file, try to change it.\n")

  // Wait 10s for changes in file
  time.Sleep(10 * time.Second)
  o.Close()
}
```
