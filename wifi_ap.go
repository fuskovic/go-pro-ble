//go:build darwin

package ble

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

var (
	errUnsupportedOs = errors.New("unsupported operating system")
	errNoInterface   = errors.New("no wireless interface found")
)

// ConnectToWifiAccessPoint is for Darwin systems only.
func ConnectToWifiAccessPoint(ssid, password string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux":
		interfaceName, err := getWirelessInterfaceLinux()
		if err != nil {
			return fmt.Errorf("failed to detect wireless interface for linux: %w", err)
		}
		cmd = exec.Command("nmcli", "dev", "wifi", "connect", ssid, "password", password, "iface", interfaceName)
	case "windows":
		// TODO
	case "darwin":
		cmd = exec.Command("networksetup", "-setairportnetwork", "en0", ssid, password)
	default:
		return errUnsupportedOs
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to connect to %q: %s", err, output)
	}
	return nil
}

func getWirelessInterfaceLinux() (string, error) {
	cmd := exec.Command("nmcli", "-t", "-f", "DEVICE,TYPE", "device")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		fields := strings.Split(line, ":")
		if len(fields) == 2 && fields[1] == "wifi" {
			return fields[0], nil
		}
	}
	return "", errNoInterface
}
