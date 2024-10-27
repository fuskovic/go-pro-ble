package ble

var (
	// Command for enabling the wifi access point on the GoPro.
	// https://gopro.github.io/OpenGoPro/tutorials/connect-wifi#enable-wifi-ap
	WifiApControlEnable = []byte{0x03, 0x17, 0x01, 0x01}

	// Command for disabling the wifi access point on the GoPro.
	// https://gopro.github.io/OpenGoPro/tutorials/connect-wifi#enable-wifi-ap
	WifiApControlDisable = []byte{0x03, 0x17, 0x01, 0x00}
)
