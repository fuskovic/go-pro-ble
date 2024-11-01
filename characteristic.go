package ble

import "strings"

// The following code is based on the bluetooth-low-energy spec defined at:
// https://gopro.github.io/OpenGoPro/ble/protocol/ble_setup.html#ble-characteristics

// Characteristic represents a ble service characteristic.
type Characteristic struct {
	uuid       string
	name       string
	readable   bool
	writeable  bool
	notifiable bool
}

// read-only
func (c Characteristic) Uuid() string     { return c.uuid }
func (c Characteristic) Name() string     { return c.name }
func (c Characteristic) Readable() bool   { return c.readable }
func (c Characteristic) Writeable() bool  { return c.writeable }
func (c Characteristic) Notifiable() bool { return c.notifiable }

var Characterstics = []Characteristic{
	CmdRequest,
	CmdResponse,
	SettingsReq,
	SettingsResp,
	QueryReq,
	QueryResp,
	WifiApSsid,
	WifiApPassword,
	WifiApPower,
	WifiApState,
	NetworkMgmtReq,
	NetworkMgmtResp,
	ModelCode,
}

// https://gopro.github.io/OpenGoPro/ble/protocol/ble_setup.html#ble-characteristics
var baseFormat = "b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b"

// format takes the base format('b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b') and replaces 'XXXX' with s.
func format(s string) string {
	return strings.Replace(baseFormat, "XXXX", s, 1)
}

// wifi-access-point
var (
	// GP-0002	WiFi AP SSID: Read + Write.
	//
	// GP-XXXX is shorthand for GoPro’s 128-bit UUID: b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b
	WifiApSsid = Characteristic{
		uuid:      format("0002"),
		name:      "wifi-access-point-ssid",
		readable:  true,
		writeable: true,
	}

	// GP-0003	WiFi AP Password: Read + Write
	//
	// GP-XXXX is shorthand for GoPro’s 128-bit UUID: b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b
	WifiApPassword = Characteristic{
		uuid:      format("0003"),
		name:      "wifi-access-point-password",
		readable:  true,
		writeable: true,
	}

	// GP-0004	WiFi AP Power: Write
	//
	// GP-XXXX is shorthand for GoPro’s 128-bit UUID: b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b
	WifiApPower = Characteristic{
		uuid:      format("0004"),
		name:      "wifi-access-point-power",
		writeable: true,
	}
	// GP-0005	WiFi AP State: Read + Indicate
	//
	// GP-XXXX is shorthand for GoPro’s 128-bit UUID: b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b
	WifiApState = Characteristic{
		uuid:     format("0005"),
		name:     "wifi-access-point-state",
		readable: true,
	}
)

// camera-management
var (
	// GP-0091	Network Management Command: Write
	//
	// GP-XXXX is shorthand for GoPro’s 128-bit UUID: b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b
	NetworkMgmtReq = Characteristic{
		uuid:       format("0091"),
		name:       "network-management-request",
		readable:   false,
		writeable:  true,
		notifiable: false,
	}
	// GP-0092	Network Management Response: Notify
	//
	// GP-XXXX is shorthand for GoPro’s 128-bit UUID: b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b
	NetworkMgmtResp = Characteristic{
		uuid:       format("0092"),
		name:       "network-management-response",
		readable:   false,
		writeable:  false,
		notifiable: true,
	}
)

// control+query
var (
	// GP-0072	Command:Write
	//
	// GP-XXXX is shorthand for GoPro’s 128-bit UUID: b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b
	CmdRequest = Characteristic{
		uuid:       format("0072"),
		name:       "command-request",
		readable:   false,
		writeable:  true,
		notifiable: false,
	}

	// GP-0073	Command Response: Notify
	//
	// GP-XXXX is shorthand for GoPro’s 128-bit UUID: b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b
	CmdResponse = Characteristic{
		uuid:       format("0073"),
		name:       "command-response",
		readable:   false,
		writeable:  false,
		notifiable: true,
	}

	// GP-0074	Settings: Write
	//
	// GP-XXXX is shorthand for GoPro’s 128-bit UUID: b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b
	SettingsReq = Characteristic{
		uuid:       format("0074"),
		name:       "settings-request",
		readable:   false,
		writeable:  true,
		notifiable: false,
	}

	// GP-0075	Settings Response: Notify
	//
	// GP-XXXX is shorthand for GoPro’s 128-bit UUID: b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b
	SettingsResp = Characteristic{
		uuid:       format("0075"),
		name:       "settings-response",
		readable:   false,
		writeable:  false,
		notifiable: true,
	}

	// GP-0076	Query Request: Write
	//
	// GP-XXXX is shorthand for GoPro’s 128-bit UUID: b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b
	QueryReq = Characteristic{
		uuid:       format("0076"),
		name:       "query-request",
		readable:   false,
		writeable:  true,
		notifiable: false,
	}

	// GP-0077	Query Response: Notify
	//
	// GP-XXXX is shorthand for GoPro’s 128-bit UUID: b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b
	QueryResp = Characteristic{
		uuid:       format("0077"),
		name:       "query-response",
		readable:   false,
		writeable:  false,
		notifiable: true,
	}
)

// used for debugging purposes
var ModelCode = Characteristic{
	uuid:       "00002a26-0000-1000-8000-00805f9b34fb",
	name:       "model-code",
	readable:   true,
	writeable:  false,
	notifiable: false,
}
