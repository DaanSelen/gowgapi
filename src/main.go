package main

import (
	"gowgapi/wgsetup"
	"log"
)

func main() {
	// Setting up WireGuard
	log.Println("Checking WireGuard status on system.")
	if !wgsetup.CheckInstall() {
		wgsetup.SetupInstall()
	}
}
