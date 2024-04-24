package wgsetup

import "log"

const (
	wireGuard_Pkg string = "wireguard"
	wireGuard_Dir string = "/etc/wireguard"
	wgQuick_Loc   string = "/usr/bin/wg-quick"
)

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
