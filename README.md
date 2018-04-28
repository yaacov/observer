# Observer

Go event emitter and listener with builtin file watcher package.

[![Build Status](https://travis-ci.org/yaacov/observer.svg?branch=master)](https://travis-ci.org/yaacov/observer)
[![GoDoc](https://godoc.org/github.com/yaacov/observer?status.svg)](https://godoc.org/github.com/yaacov/observer)

## Description

This observer implements event emitter and listener pattern in go,
the observer register a list of listener functions and implement an event emitter,
once an event is emitted, all listener functions will be called.

This observer also abstracts watching for file changes, users can register a list for files to watch,
once a file is watched, events will be emitted automatically on each file modification.
Common use cases are watching for changing in config files, and wating for code changes.

This observer is using golang [channels](https://gobyexample.com/channels) for emiting events and [fsnotify](https://github.com/fsnotify/fsnotify) for watching file changes.

#### Examples:
  - [Emit string events](#emit-string-events)
  - [Watch files, emit file cahnge events](#watch-files-emit-file-change-events)

## Develop

```
$ go get -u github.com/golang/dep/cmd/dep
$ make vendor
$ make
$ ./observe
```

## Install

```
go get -u github.com/yaacov/observer
```

Install using `go get` will install the package and the CLI tool _observer_

## CLI

_observer_ is a cli tool for watching files and executing shell commands on file modification, it is used to call
an action on file change, examples of use can be restart and app when config file changes, recompile code when code updates and send image to server when image change.

[inotify-tools](https://github.com/rvoicilas/inotify-tools/wiki), are a set of linux tools that monitor files for changes, they may be better choice for file monitoring on more complex cases.

#### Get help:
``` sh
observer -h
```

#### Call a server api when config file chage:
``` sh
observer -r "curl -X POST http://127.0.0.1:8000/api/v1/-/restart" -w "/root/.aws/config"
```

## API

See [examples](#examples-1) for usage examples.

| Method                         | Description                       |
|--------------------------------|-----------------------------------|
| Open()                         | Open the observer channels        |
| Close()                        | Close the observer channels       |
| AddListener(callback Listener) | Add a listener function to run on event |
| Emit(event interface{})        | Emit event                        |
| Watch(files []string)          | Watch for file changes, and emit a file change events |

| Type                           |                                   | Description |
|--------------------------------|-----------------------------------|-------------|
| WatchEvent                     | struct{ Name string, Op uint32 }  | Event type emitted by file watcher |
| Listener                       | func(interface{})                 | Function type for listeners        |
| Observer                       | struct{ Verbose bool }            | The observer object                |

## Watching files for modifications

Watching files can be done using exact file name, or shell pattern matching.

#### Watching for exact file names:
``` go
Watch([]string{"./aws/config", "./aws/credentials"})
```

#### Watching files using shell pattern matching:
``` go
Watch([]string{"./kube/*.yml"})
```

#### Note:
We can not expand tilde to home directory, `~/.config` will not work as expected.
If needed users can use golang's [os/user/](https://golang.org/pkg/os/user/) package.

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
o.Watch([]string{"../LICENSE"})
defer o.Close()

// Add a listener that logs events
o.AddListener(func(e interface{}) {
  watchEvent := e.(observer.WatchEvent)
  log.Printf("File modified: %s.\n", e.Name)
})
```

[file-watch-pattern.go](/examples/file-watch-pattern.go)

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
