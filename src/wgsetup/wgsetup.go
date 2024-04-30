package wgsetup

// Package for getting the system ready for WireGuard usage.

import "log"

func CheckInstall() bool {
	if !statDirectory() || !statPackage() {
		log.Println("Installation incomplete, installing...")
		return false
	} else {
		log.Println("Installation is complete.")
		return true
	}
}

func SetupInstall() {
	createDirectoryTree()
	installPackage()
	modQuick()

	log.Println("Rerunning check to verify...")
	if CheckInstall() {
		log.Println("Successfully installed.")
	}
}
