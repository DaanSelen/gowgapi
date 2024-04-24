package main

import (
	"gowgapi/wgrest"
	"gowgapi/wgsetup"
	"gowgapi/wgsqlite"
	"log"
)

func main() {
	// Setting up the WireGuard application and its dependencies.
	log.Println("Checking WireGuard status on system.")
	if !wgsetup.CheckInstall() {
		wgsetup.SetupInstall()
	}
	// Done checking installation, everything should be setup.

	// Begin setting up persistency.
	if !wgsqlite.InitDatabase() {
		log.Fatal("Failed to setup persistency.")
	}
	// Done setting up persistency

	// Start the frontend API.
	wgrest.InitFrontend()
	// Done setting up the program.
}
