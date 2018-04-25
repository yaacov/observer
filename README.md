# Observer

Event emitter and listener with builtin file watcher.

[![Build Status](https://travis-ci.org/yaacov/observer.svg?branch=master)](https://travis-ci.org/yaacov/observer)

## Description

This observer emplements event emitter and listener pattern in go,
the observer register a list of listener functions and implement an event emitter,
once an event is emited, all listener functions will be called.

This observer also abstruct watching for file changes, users can register a list for files to watch,
once a file is watched, events will be emitted automatically on each file modification.

This Observe is using golang [channels](https://gobyexample.com/channels) for emiting events and [fsnotify](https://github.com/fsnotify/fsnotify) for watching file changes.

#### Examples:
  - [Emit string events](#emit-string-events)
  - [Watch files, emit file cahnge events](#watch-emit-events-on-file-modification)

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

## Examples

### Emit string events

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
  defer o.Close()

  // Add a listener that logs events
  o.AddListener(func(e interface{}) {
    log.Printf("Received: %s.\n", e.(string))
  })

  // This event will be picked by the listener
  o.Emit("Hello")

  // Close observer
  time.Sleep(3 * time.Second)
}
```

### Watch, emit events on file modification

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
  defer o.Close()

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
}
```
