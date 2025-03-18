//go:build darwin

package ble

import (
	"fmt"
	"os/exec"
)

const interfaceName = "en0"

// ConnectToWifiAccessPoint is for Darwin systems only.
func ConnectToWifiAccessPoint(ssid, password string) error {
	cmd := exec.Command("networksetup", "-setairportnetwork", interfaceName, ssid, password)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to connect to %q: %s", err, output)
	}
	return nil
}
