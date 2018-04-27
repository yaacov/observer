# Observer

Go event emitter and listener with builtin file watcher.

[![Build Status](https://travis-ci.org/yaacov/observer.svg?branch=master)](https://travis-ci.org/yaacov/observer)

## Description

This observer emplements event emitter and listener pattern in go,
the observer register a list of listener functions and implement an event emitter,
once an event is emited, all listener functions will be called.

This observer also abstruct watching for file changes, users can register a list for files to watch,
once a file is watched, events will be emitted automatically on each file modification.

This observer is using golang [channels](https://gobyexample.com/channels) for emiting events and [fsnotify](https://github.com/fsnotify/fsnotify) for watching file changes.

#### Examples:
  - [Emit string events](#emit-string-events)
  - [Watch files, emit file cahnge events](#watch-files-emit-file-change-events)

## Develop

```
$ make vendor
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
| Watch(files []string)          | Watch for file changes, and emit a file change events, file names can have   |

## Examples

### Emit string events

Example of event listener and emitter.

[emit.go](/examples/emit.go)

``` go
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
```

### Watch files, emit file change events

Example of file watching and listener.

[file-watch.go](/examples/file-watch.go)

``` go
// Open an observer and start watching for files by file name
o := observer.Observer{}
o.Watch([]string{"../LICENSE", "~/.kube/config"})
defer o.Close()

// Add a listener that logs events
o.AddListener(func(e interface{}) {
  log.Printf("File modified: %v.\n", e)
})
```

``` go
// Open an observer and start watching for file matching shell pattern
o := observer.Observer{}
o.Watch([]string{"*.html", "css/*.scss"})
defer o.Close()

// Add a listener that logs events
o.AddListener(func(e interface{}) {
  log.Printf("File modified: %v.\n", e)
})
```
