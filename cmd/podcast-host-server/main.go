// Implements the server app that will run podcast-host

package main

import (
	_ "codenex.us/ralph/podcast-host"
	"codenex.us/ralph/podcast-host/system"
	"log"
	"os"
)

func main() {
	sys, err := system.NewSystem() // Create the system

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	sys.Start() // Start the system
}
