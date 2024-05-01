package wgsetup

// Package for getting the system ready for WireGuard usage.

import "log"

const (
	wireGuard_Pkg string = "wireguard" // Name of the package for WireGuard.
	openssl_Pkg   string = "openssl"   // Name of the package for OpenSSL.
)

func Install() {
	if repairInstall() {
		verifyInstall()
	}
	ensureCert()
}

func repairInstall() bool {
	if !statWGDirectory() {
		log.Println("Installation incomplete, installing...")
		createWGDirectoryTree()
		return false
	} else if !statCertDirectory() {
		createCertDirectory()
		return false
	} else if !statPackage(wireGuard_Pkg) {
		installPackage(wireGuard_Pkg)
		modWGQuick()
		return false
	} else if !statPackage(openssl_Pkg) {
		installPackage(openssl_Pkg)
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
