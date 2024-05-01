package wgsetup

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

const (
	wireGuard_Pkg string = "wireguard"
	wireGuard_Dir string = "/etc/wireguard"
	wgQuick_Loc   string = "/usr/bin/wg-quick"
)

func statCertDirectory() bool {
	_, err := os.Stat("./certificate")
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("Certificate directory not present.")
			return false
		} else {
			log.Println("Unknown error while checking certificate directory presence.")
			return false
		}
	} else {
		return true
	}
}

func createCertDirectory() {
	err := os.Mkdir("./certificate", 0755)
	if err != nil {
		if os.IsPermission(err) {
			log.Printf("Error creating directory.")
			log.Fatal("Either manually create the directory with correct permissions (755), or run this program as root (sudo).")
		} else {
			log.Println("Unable to create the Certificate directory.")
		}
	} else {
		log.Println("Created the Certificate directory.")
	}
}

func statWGDirectory() bool {
	_, err1 := os.Stat(wireGuard_Dir)
	_, err2 := os.Stat((wireGuard_Dir + "/iface-config"))
	_, err3 := os.Stat((wireGuard_Dir + "/iface-client"))

	if err1 != nil || err2 != nil || err3 != nil {
		if os.IsNotExist(err1) || os.IsNotExist(err2) || os.IsNotExist(err3) {
			log.Println("One of the needed directories is not present.")
			return false
		} else {
			log.Println("Unknown error while checking directory presence.")
			return false
		}
	} else {
		return true
	}
}

func createWGDirectoryTree() {
	err1 := os.Mkdir(wireGuard_Dir, 0755)
	err2 := os.Mkdir((wireGuard_Dir + "/iface-config"), 0755)
	err3 := os.Mkdir((wireGuard_Dir + "/iface-client"), 0755)
	if err1 != nil || err2 != nil || err3 != nil {
		if os.IsPermission(err1) || os.IsPermission(err2) || os.IsPermission(err3) {
			log.Printf("Error creating directory.")
			log.Fatal("Either manually create the directory with correct permissions (755), or run this program as root (sudo).")
		} else {
			log.Println("Unable to create the WireGuard directory.", err1, err2, err3)
		}
	} else {
		log.Println("Created the WireGuard directory tree.")
	}
}

func statWGPackage() bool {
	cmd := exec.Command("dpkg", "-s", wireGuard_Pkg)
	output, _ := cmd.CombinedOutput()

	if strings.Contains(string(output), "Status: install ok installed") {
		return true
	} else if strings.Contains(string(output), "is not installed") {
		return false
	} else {
		return false
	}
}

func installWGPackage() {
	cmd := exec.Command("apt", "install", "-y", wireGuard_Pkg)
	output, _ := cmd.CombinedOutput()

	if strings.Contains(string(output), "Setting up wireguard") {
		log.Println("Installing Wireguard using APT.")
	} else if strings.Contains(string(output), "Permission denied") {
		log.Fatal("APT permission error.")
	}
}

func modWGQuick() bool {
	content, err := os.ReadFile("/usr/bin/wg-quick")
	if err != nil {
		log.Println("Error reading file:", err)
		return false
	}

	oldString := `[[ $CONFIG_FILE =~ ^[a-zA-Z0-9_=+.-]{1,15}$ ]] && CONFIG_FILE="/etc/wireguard/$CONFIG_FILE.conf"` // Only the part after the '&&' matters.
	newString := fmt.Sprintf(`[[ $CONFIG_FILE =~ ^[a-zA-Z0-9_=+.-]{1,15}$ ]] && CONFIG_FILE="%s/$CONFIG_FILE.conf"`, (wireGuard_Dir + "/iface-config"))

	modifiedContent := strings.Replace(string(content), oldString, newString, -1)

	err = os.WriteFile("/usr/bin/wg-quick", []byte(modifiedContent), 0644)
	if err != nil {
		log.Println("Error writing file:", err)
		return false
	}

	log.Println("wg-quick modification complete.")
	return true
}
