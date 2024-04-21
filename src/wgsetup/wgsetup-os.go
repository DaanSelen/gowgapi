package wgsetup

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func statDirectory() bool {
	_, err := os.Stat(wireGuard_Dir)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("Directory does not exist.")
			return false
		} else {
			log.Println("Unable to check the WireGuard directory.")
			return false
		}
	} else {
		return true
	}
}

func createDirectory() {
	err := os.Mkdir(wireGuard_Dir, 0755)
	if err != nil {
		if os.IsPermission(err) {
			log.Printf("Error creating directory: %s\n", err)
			log.Fatal("Either manually create the directory with correct permissions (755), or run this program as root (sudo).")
		} else {
			log.Println("Unable to create the WireGuard directory.")
		}
	} else {
		log.Println("Created the WireGuard directory.")
	}
}

func statPackage() bool {
	cmd := exec.Command("dpkg", "-s", wireGuard_Pkg)
	output, _ := cmd.CombinedOutput()

	if strings.Contains(string(output), "install ok installed") {
		return true
	} else if strings.Contains(string(output), "is not installed") {
		log.Println("Package is not preset.")
		return false
	} else {
		return false
	}
}

func installPackage() {
	cmd := exec.Command("apt", "install", "-y", wireGuard_Pkg)
	output, _ := cmd.CombinedOutput()

	if strings.Contains(string(output), "Setting up wireguard") {
		log.Println("Installing Wireguard using APT.")
	} else if strings.Contains(string(output), "Permission denied") {
		log.Fatal("APT permission error.")
	}
}
