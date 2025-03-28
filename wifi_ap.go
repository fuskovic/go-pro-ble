//go:build darwin

package ble

import (
	"errors"
	"fmt"
	"os"
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
	switch runtime.GOOS {
	case "linux":
		return connectLinux(ssid, password)
	case "windows":
		return connectWindows(ssid, password)
	case "darwin":
		return connectDarwin(ssid, password)
	default:
		return errUnsupportedOs
	}
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

var windowsWifiProfileXML = `
<?xml version="1.0"?>
<WLANProfile xmlns="http://www.microsoft.com/networking/WLAN/profile/v1">
	<name>%s</name>
	<SSIDConfig>
		<SSID>
			<name>%s</name>
		</SSID>
	</SSIDConfig>
	<connectionType>ESS</connectionType>
	<connectionMode>auto</connectionMode>
	<MSM>
		<security>
			<authEncryption>
				<authentication>WPA2PSK</authentication>
				<encryption>AES</encryption>
				<useOneX>false</useOneX>
			</authEncryption>
			<sharedKey>
				<keyType>passPhrase</keyType>
				<protected>false</protected>
				<keyMaterial>%s</keyMaterial>
			</sharedKey>
		</security>
	</MSM>
</WLANProfile>
`

func connectWindows(ssid, password string) error {
	profile := fmt.Sprintf(windowsWifiProfileXML, ssid, ssid, password)

	tmpFile := fmt.Sprintf("%s.xml", ssid)
	err := os.WriteFile(tmpFile, []byte(profile), 0644)
	if err != nil {
		return fmt.Errorf("failed to write profile XML: %w", err)
	}
	defer os.Remove(tmpFile)

	addCmd := exec.Command("netsh", "wlan", "add", "profile", fmt.Sprintf("filename=%s", tmpFile), "user=current")
	addOut, err := addCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to add profile: %v\nOutput: %s", err, string(addOut))
	}

	connectCmd := exec.Command("netsh", "wlan", "connect", fmt.Sprintf("name=%s", ssid))
	connectOut, err := connectCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to connect: %w;\noutput: %s", err, connectOut)
	}
	return nil
}

func connectLinux(ssid, password string) error {
	interfaceName, err := getWirelessInterfaceLinux()
	if err != nil {
		return fmt.Errorf("failed to detect wireless interface for linux: %w", err)
	}
	cmd := exec.Command("nmcli", "dev", "wifi", "connect", ssid, "password", password, "iface", interfaceName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to connect: %w;\noutput: %s", err, output)
	}
	return nil
}

func connectDarwin(ssid, password string) error {
	cmd := exec.Command("networksetup", "-setairportnetwork", "en0", ssid, password)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to connect: %w;\noutput: %s", err, output)
	}
	return nil
}
