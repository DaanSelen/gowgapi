package wgsetup

import "log"

const (
	wireGuard_Pkg string = "wireguard"
	wireGuard_Dir string = "/etc/wireguard"
)

func SetupInstall() {
	var dirCheck bool
	var pkgCheck bool

	if !statDirectory() {
		createDirectory()
		dirCheck = statDirectory()
	} else {
		dirCheck = true
		log.Println("Directory is present.")
	}

	if !statPackage() {
		installPackage()
		pkgCheck = statPackage()
	} else {
		pkgCheck = true
		log.Println("Package is present.")
	}

	if dirCheck && pkgCheck {
		log.Println("All checks passed. WireGuard is set up.")
	}
}
