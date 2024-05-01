package wgsetup

// Package for getting the system ready for WireGuard usage.

import "log"

func Install() {
	if repairInstall() {
		verifyInstall()
	}
}

func repairInstall() bool {
	if !statWGDirectory() {
		log.Println("Installation incomplete, installing...")
		createWGDirectoryTree()
		return false
	} else if !statCertDirectory() {
		createCertDirectory()
		return false
	} else if !statWGPackage() {
		installWGPackage()
		modWGQuick()
		return false
	} else {
		return true
	}
}

func verifyInstall() {
	log.Println("Verifying install...")
	if repairInstall() {
		log.Println("Successfully installed.")
	}
}
