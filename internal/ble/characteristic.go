package ble

import (
	"slices"
)

// The following code is based on the bluetooth-low-energy spec defined at:
// https://gopro.github.io/OpenGoPro/ble/protocol/ble_setup.html#ble-characteristics

// wifi-access-point
var (
	// GP-0002	WiFi AP SSID: Read + Write.
	//
	// GP-XXXX is shorthand for GoPro’s 128-bit UUID: b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b
	WifiApSsidUuid = Characteristic(format("0002"))

	// GP-0003	WiFi AP Password: Read + Write
	//
	// GP-XXXX is shorthand for GoPro’s 128-bit UUID: b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b
	WifiApPasswordUuid = Characteristic(format("0003"))

	// GP-0004	WiFi AP Power: Write
	//
	// GP-XXXX is shorthand for GoPro’s 128-bit UUID: b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b
	WifiApPowerUuid = Characteristic(format("0004"))
	// GP-0005	WiFi AP State: Read + Indicate
	//
	// GP-XXXX is shorthand for GoPro’s 128-bit UUID: b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b
	WifiApStateUuid = Characteristic(format("0005"))
)

// camera-management
var (
	// GP-0091	Network Management Command: Write
	//
	// GP-XXXX is shorthand for GoPro’s 128-bit UUID: b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b
	NetworkMgmtReqUuid Characteristic = Characteristic(format("0091"))
	// GP-0092	Network Management Response: Notify
	//
	// GP-XXXX is shorthand for GoPro’s 128-bit UUID: b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b
	NetworkMgmtRespUuid Characteristic = Characteristic(format("0092"))
)

// control+query
var (
	// GP-0072	Command:Write
	//
	// GP-XXXX is shorthand for GoPro’s 128-bit UUID: b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b
	CmdRequestUuid Characteristic = Characteristic(format("0072"))

	// GP-0073	Command Response: Notify
	//
	// GP-XXXX is shorthand for GoPro’s 128-bit UUID: b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b
	CmdResponseUuid Characteristic = Characteristic(format("0073"))

	// GP-0074	Settings: Write
	//
	// GP-XXXX is shorthand for GoPro’s 128-bit UUID: b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b
	SettingsReqUuid Characteristic = Characteristic(format("0074"))

	// GP-0075	Settings Response: Notify
	//
	// GP-XXXX is shorthand for GoPro’s 128-bit UUID: b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b
	SettingsRespUuid Characteristic = Characteristic(format("0075"))

	// GP-0076	Query: Write
	//
	// GP-XXXX is shorthand for GoPro’s 128-bit UUID: b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b
	QueryReqUuid Characteristic = Characteristic(format("0076"))

	// GP-0077	Query Response: Notify
	//
	// GP-XXXX is shorthand for GoPro’s 128-bit UUID: b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b
	QueryRespUuid Characteristic = Characteristic(format("0077"))
)

// used for debugging purposes
var ModelCode Characteristic = Characteristic(format("00002a26-0000-1000-8000-00805f9b34fb"))

var (
	Characterstics = []Characteristic{
		CmdRequestUuid,
		CmdResponseUuid,
		SettingsReqUuid,
		SettingsRespUuid,
		QueryReqUuid,
		QueryRespUuid,
		WifiApSsidUuid,
		WifiApPasswordUuid,
		WifiApPowerUuid,
		WifiApStateUuid,
		NetworkMgmtReqUuid,
		NetworkMgmtRespUuid,
		ModelCode,
	}

	ReadableCharacteristics = []Characteristic{
		WifiApSsidUuid,
		WifiApSsidUuid,
		WifiApStateUuid,
		ModelCode,
	}

	WriteableCharacteristics = []Characteristic{
		WifiApSsidUuid,
		WifiApPasswordUuid,
		WifiApPowerUuid,
		NetworkMgmtReqUuid,
		CmdRequestUuid,
		SettingsReqUuid,
		QueryReqUuid,
	}

	NotifiableCharacteristics = []Characteristic{
		NetworkMgmtRespUuid,
		CmdResponseUuid,
		SettingsRespUuid,
		QueryRespUuid,
	}
)

type Characteristic string

func (c Characteristic) Name() string {
	var name string
	switch c {
	case CmdRequestUuid:
		name = "command-request"
	case CmdResponseUuid:
		name = "command-response"
	case SettingsReqUuid:
		name = "settings-request"
	case SettingsRespUuid:
		name = "settings_response"
	case QueryReqUuid:
		name = "query-request"
	case QueryRespUuid:
		name = "query-response"
	case WifiApSsidUuid:
		name = "wifi-access-point-ssid"
	case WifiApPasswordUuid:
		name = "wifi-access-point-password"
	case WifiApPowerUuid:
		name = "wifi-access-point-power"
	case WifiApStateUuid:
		name = "wifi-access-point-state"
	case NetworkMgmtReqUuid:
		name = "network-management-request"
	case NetworkMgmtRespUuid:
		name = "network-management-response"
	default:
		name = "unknown"
	}
	return name
}

func (c Characteristic) Readable() bool {
	return slices.Contains(ReadableCharacteristics, c)
}

func (c Characteristic) Writeable() bool {
	return slices.Contains(WriteableCharacteristics, c)
}

func (c Characteristic) Notifiable() bool {
	return slices.Contains(NotifiableCharacteristics, c)
}

func (c Characteristic) String() string {
	return string(c)
}
