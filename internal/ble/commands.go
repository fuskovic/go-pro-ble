package ble

var (
	// https://gopro.github.io/OpenGoPro/tutorials/connect-wifi#enable-wifi-ap
	WifiApControlEnable = []byte{0x03, 0x17, 0x01, 0x01}
	// https://gopro.github.io/OpenGoPro/tutorials/connect-wifi#enable-wifi-ap
	WifiApControlDisable = []byte{0x03, 0x17, 0x01, 0x00}
)
