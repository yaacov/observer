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
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"

	"github.com/yaacov/observer/observer"
)

func main() {
	var err error

	// Parse cli arguments
	watchPtr := flag.String("w", "./*", "space sperated list of files to watch.")
	runPtr := flag.String("r", "./run.sh", "shell command to run.")
	verbosePtr := flag.Bool("V", false, "dump debug data.")
	flag.Parse()

	// Get watchFiles
	watchFiles := strings.Split(*watchPtr, " ")

	// Get command line to run on events
	cmd := strings.Split(*runPtr, " ")

	// Open observer and start watching
	o := observer.Observer{}
	o.Verbose = *verbosePtr

	defer o.Close()

	// Watch for changes in files
	err = o.Watch(watchFiles)
	if err != nil {
		log.Fatal("[Error] watch files: ", err)
	}

	// Add a listener for events
	o.AddListener(func(e interface{}) {
		log.Printf("Received: %v.\n", e)

		if len(cmd) == 1 {
			err = exec.Command(cmd[0]).Run()
		} else if len(cmd) > 1 {
			err = exec.Command(cmd[0], cmd[1:]...).Run()
		}

		if err != nil {
			log.Printf("[Error] running event listener: %s.\n", err)
		}
	})

	// Log watcher starting.
	log.Print("Observer starting.")
	log.Print("Press Ctrl+C to exit.")

	// Wait for Ctrl+C
	waitCtrlC := make(chan os.Signal, 1)
	signal.Notify(waitCtrlC, os.Interrupt)

	<-waitCtrlC
}
