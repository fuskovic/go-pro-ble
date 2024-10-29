package ble

// https://gopro.github.io/OpenGoPro/ble/protocol/data_protocol.html#commands
type COMMAND []byte

func (c COMMAND) Bytes() []byte { return []byte(c) }

var (
	// Command for enabling the wifi access point on the GoPro.
	// https://gopro.github.io/OpenGoPro/tutorials/connect-wifi#enable-wifi-ap
	WIFI_AP_CONTROL_ENABLE COMMAND = []byte{0x03, 0x17, 0x01, 0x01}

	// Command for disabling the wifi access point on the GoPro.
	// https://gopro.github.io/OpenGoPro/tutorials/connect-wifi#enable-wifi-ap
	WIFI_AP_CONTROL_DISABLE COMMAND = []byte{0x03, 0x17, 0x01, 0x00}
)

// // https://gopro.github.io/OpenGoPro/ble/protocol/data_protocol.html#commands
type COMMAND_ID byte

func (c COMMAND_ID) Byte() byte { return byte(c) }

const (
	// ID for the command that toggles the wifi-access-point on the GoPro.
	// https://gopro.github.io/OpenGoPro/ble/protocol/data_protocol.html#commands
	WIFI_AP_TOGGLE_COMMAND_ID COMMAND_ID = 0x17
)
