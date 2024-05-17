package wgparser

import (
	"fmt"
	"gowgapi/wgsetup"
	"gowgapi/wgsqlite"
	"log"
	"os"
	"strings"
)

// Package for converting 'SQLified' config to .conf file.

func ParseAll() {
	allInterface := wgsqlite.QueryAllInterfaces()
	for _, iface := range allInterface {
		var configPath string = wgsetup.WireGuard_Dir + "/iface-config/" + iface.Name + ".conf"

		_, err := os.Stat(configPath)
		log.Println(err)
		if err != nil || os.IsNotExist(err) {
			os.Create(configPath)
		}

		log.Println(configPath)
		err = os.WriteFile(configPath, []byte(parse(iface.Name, iface.Address, iface.Port, iface.Out_Interface)), 0644)
		if err != nil {
			log.Println(err, "test")
		}
	}
}

func CreateAndParseInterface(ifaceName, address, port, outiface string) {
	log.Println(parse(ifaceName, address, port, outiface))
}

func parse(ifaceName, address, port, out_iface string) string {
	//GENERATE The WireGuard Interface Config part.
	privKey := wgsqlite.QuieryPrivKey(ifaceName)
	fw_rules := strings.ReplaceAll(default_fw_rules, "<WG-INTERFACE>", ifaceName)
	fw_rules = strings.ReplaceAll(fw_rules, "<OUTGOING-INTERFACE>", out_iface)

	config_template := strings.ReplaceAll(interface_template, "<FW-RULES>", fw_rules)
	config_template = fmt.Sprintf(config_template, address, port, privKey)

	//GENERATE The WireGuard Peers part.
	return config_template
}
