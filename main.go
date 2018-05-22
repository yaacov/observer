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
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"sync"
	"time"

	"github.com/yaacov/observer/observer"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func runScript(s string) (err error) {
	// Get command line to run on events.
	c := strings.Split(s, " ")

	// If script has no args, just run it, o/w sent args.
	if len(c) == 1 {
		err = exec.Command(c[0]).Run()
	} else if len(c) > 1 {
		err = exec.Command(c[0], c[1:]...).Run()
	}

	return
}

func printUsage() {
	fmt.Println("Usage:")
	flag.PrintDefaults()

	fmt.Println("")
	fmt.Println("Examples:")
	fmt.Println("  observer -w main.c -r ./run.sh")
	fmt.Println("  observer -w main.c -w src/*.c -r run.sh -d 1")

	os.Exit(1)
}

func main() {
	var err error
	var watchFiles arrayFlags
	var scripts arrayFlags

	// Create a mutex for the event listener.
	mutex := &sync.Mutex{}

	// Parse cli arguments.
	flag.Var(&watchFiles, "w", "list of files to watch.")
	flag.Var(&scripts, "r", "list of scripts to run on file modifiaction event.")
	bufferSecPtr := flag.Int("d", 0, "buffer events for N sec.")
	verbosePtr := flag.Bool("V", false, "dump debug data.")

	flag.Usage = printUsage
	flag.Parse()

	// Check user input.
	if len(watchFiles) < 1 {
		fmt.Println("[Error] no watch files.")
		flag.Usage()
	}

	// Sanity check.
	if len(scripts) < 1 {
		fmt.Println("[Error] no scripts to run.")
		flag.Usage()
	}

	// Open observer and start watching.
	o := observer.Observer{}
	defer o.Close()

	// Set verbosity.
	o.Verbose = *verbosePtr

	// Set damping time.
	if *bufferSecPtr != 0 {
		sec := time.Duration(*bufferSecPtr) * time.Second
		o.SetBufferDuration(sec)
	}

	// Watch for changes in files.
	err = o.Watch(watchFiles)
	if err != nil {
		log.Fatal("[Error] watch files: ", err)
	}

	// Add a listener for events.
	o.AddListener(func(e interface{}) {
		// Lock the listener.
		mutex.Lock()
		defer mutex.Unlock()

		// Log the event.
		log.Printf("[Info]Received: %v\n", e)

		for _, s := range scripts {
			// Try to run a script. and check for errors running script.
			if err = runScript(s); err != nil {
				log.Printf("[Error] running event listener: %s\n", err)
			}
		}
	})

	// Log watcher starting.
	log.Print("Observer starting.")
	log.Print("Press Ctrl+C to exit.")

	// Wait for Ctrl+C.
	waitCtrlC := make(chan os.Signal, 1)
	signal.Notify(waitCtrlC, os.Interrupt)

	<-waitCtrlC
}
