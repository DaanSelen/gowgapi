package main

// Main function of program, used for initial start-up.

import (
	"gowgapi/wgrest"
	"gowgapi/wgsetup"
	"gowgapi/wgsqlite"
	"log"
	"sync"
)

func main() {
	var waitGroup sync.WaitGroup

	// Setting up the WireGuard application and its dependencies.
	log.Println("Checking dependency presence on system.")
	wgsetup.Install()
	// Done checking installation, everything should be setup.

	// Begin setting up persistency.
	wgsqlite.InitDatabase()
	// Done setting up persistency

	// Start the frontend API, with an increment to the waitGroup delta.
	waitGroup.Add(1)
	go wgrest.InitFrontend(&waitGroup)
	log.Println("GoWGAPI Ready")
	// Done setting up the program.

	waitGroup.Wait()
}
