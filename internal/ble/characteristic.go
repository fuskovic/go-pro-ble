package ble

import "strings"

// INFO:root:Enabling notification on char 00002a19-0000-1000-8000-00805f9b34fb
// INFO:root:Enabling notification on char b5f90073-aa8d-11e3-9046-0002a5d5c51b
// INFO:root:Enabling notification on char b5f90075-aa8d-11e3-9046-0002a5d5c51b
// INFO:root:Enabling notification on char b5f90077-aa8d-11e3-9046-0002a5d5c51b
// INFO:root:Enabling notification on char b5f90079-aa8d-11e3-9046-0002a5d5c51b
// INFO:root:Enabling notification on char b5f90092-aa8d-11e3-9046-0002a5d5c51b
// INFO:root:Enabling notification on char b5f90081-aa8d-11e3-9046-0002a5d5c51b
// INFO:root:Enabling notification on char b5f90083-aa8d-11e3-9046-0002a5d5c51b
// INFO:root:Enabling notification on char b5f90084-aa8d-11e3-9046-0002a5d5c51b

// https://gopro.github.io/OpenGoPro/ble/protocol/ble_setup.html#ble-characteristics

// wifi-access-point
var (
	// GP-0002	WiFi AP SSID: Read + Write.
	//
	// GP-XXXX is shorthand for GoPro’s 128-bit UUID: b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b
	wifiApSsidUuid = Characteristic(format("0002"))

	// GP-0003	WiFi AP Password: Read + Write
	//
	// GP-XXXX is shorthand for GoPro’s 128-bit UUID: b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b
	wifiApPasswordUuid = Characteristic(format("0003"))

	// GP-0004	WiFi AP Power: Write
	//
	// GP-XXXX is shorthand for GoPro’s 128-bit UUID: b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b
	wifiApPowerUuid = Characteristic(format("0004"))
	// GP-0005	WiFi AP State: Read + Indicate
	//
	// GP-XXXX is shorthand for GoPro’s 128-bit UUID: b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b
	wifiApStateUuid = Characteristic(format("0005"))
)

// camera-management
var (
	// GP-0091	Network Management Command: Write
	//
	// GP-XXXX is shorthand for GoPro’s 128-bit UUID: b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b
	networkMgmtReqUuid Characteristic = Characteristic(format("0091"))
	// GP-0092	Network Management Response: Notify
	//
	// GP-XXXX is shorthand for GoPro’s 128-bit UUID: b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b
	networkMgmtRespUuid Characteristic = Characteristic(format("0092"))
)

// control+query
var (
	// GP-0072	Command:Write
	//
	// GP-XXXX is shorthand for GoPro’s 128-bit UUID: b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b
	cmdRequestUuid Characteristic = Characteristic(format("0072"))

	// GP-0073	Command Response: Notify
	//
	// GP-XXXX is shorthand for GoPro’s 128-bit UUID: b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b
	cmdResponseUuid Characteristic = Characteristic(format("0073"))

	// GP-0074	Settings: Write
	//
	// GP-XXXX is shorthand for GoPro’s 128-bit UUID: b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b
	settingsReqUuid Characteristic = Characteristic(format("0074"))

	// GP-0075	Settings Response: Notify
	//
	// GP-XXXX is shorthand for GoPro’s 128-bit UUID: b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b
	settingsRespUuid Characteristic = Characteristic(format("0075"))

	// GP-0076	Query: Write
	//
	// GP-XXXX is shorthand for GoPro’s 128-bit UUID: b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b
	queryReqUuid Characteristic = Characteristic(format("0076"))

	// GP-0077	Query Response: Notify
	//
	// GP-XXXX is shorthand for GoPro’s 128-bit UUID: b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b
	queryRespUuid Characteristic = Characteristic(format("0077"))
	// TODO: move this
	// modelCode           Characteristic = format("00002a26-0000-1000-8000-00805f9b34fb")
)

var (
	Characterstics = []Characteristic{
		cmdRequestUuid,
		cmdResponseUuid,
		settingsReqUuid,
		settingsRespUuid,
		queryReqUuid,
		queryRespUuid,
		wifiApSsidUuid,
		wifiApPasswordUuid,
		networkMgmtReqUuid,
		networkMgmtRespUuid,
		// modelCode,
	}
)

type Characteristic string

func (c Characteristic) Name() string {
	var name string
	switch c {
	case cmdRequestUuid:
		name = "command-request"
	case cmdResponseUuid:
		name = "command-response"
	case settingsReqUuid:
		name = "settings-request"
	case settingsRespUuid:
		name = "settings_response"
	case queryReqUuid:
		name = "query-request"
	case queryRespUuid:
		name = "query-response"
	case wifiApSsidUuid:
		name = "wifi-access-point-ssid"
	case wifiApPasswordUuid:
		name = "wifi-access-point-password"
	case wifiApPowerUuid:
		name = "wifi-access-point-power"
	case wifiApStateUuid:
		name = "wifi-access-point-state"
	case networkMgmtReqUuid:
		name = "network-management-request"
	case networkMgmtRespUuid:
		name = "network-management-response"
	default:
		name = "unknown"
	}
	return name
}

func (c Characteristic) Notifiable() bool {
	return strings.Contains(c.Name(), "response")
}

func (c Characteristic) String() string {
	return string(c)
}
