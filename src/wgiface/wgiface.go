package wgiface

// Package for managing the WireGuard interfaces

import "os/exec"

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
