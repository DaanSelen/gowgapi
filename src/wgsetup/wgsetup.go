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
	var installFault int = 0
	if !statWGDirectory() {
		log.Println("Installation incomplete, installing...")
		createWGDirectoryTree()
		installFault++
	}
	if !statCertDirectory() {
		createCertDirectory()
		installFault++
	}
	if !statPackage(wireGuard_Pkg) {
		installPackage(wireGuard_Pkg)
		modWGQuick()
		installFault++
	}
	if !statPackage(openssl_Pkg) {
		installPackage(openssl_Pkg)
		installFault++
	}
	if installFault == 0 {
		return true
	} else {
		return false
	}
}

func verifyInstall() {
	log.Println("Verifying install...")
	if repairInstall() {
		log.Println("Successfully installed.")
	}
}
