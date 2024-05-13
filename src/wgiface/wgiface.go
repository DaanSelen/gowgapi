package wgiface

// Package for managing the WireGuard interfaces

import (
	"os/exec"
	"strings"
)

func GenPrivKey() string {
	cmd := exec.Command("wg", "genkey")
	output, _ := cmd.CombinedOutput()

	return strings.TrimSpace(string(output))
}

func StartService(serviceName string) error {
	cmd := exec.Command("systemctl", "start", serviceName)
	err := cmd.Run()
	return err
}

func StopService(serviceName string) error {
	cmd := exec.Command("systemctl", "stop", serviceName)
	err := cmd.Run()
	return err
}
