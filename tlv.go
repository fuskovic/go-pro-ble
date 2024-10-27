package ble

// https://gopro.github.io/OpenGoPro/ble/protocol/data_protocol.html#commands
const WifiApToggleCmdID = 0x17

type TLV_RESPONSE byte

// https://gopro.github.io/OpenGoPro/ble/protocol/data_protocol.html#command-response
const (
	TLV_RESPONSE_SUCCESS           TLV_RESPONSE = 0x00
	TLV_RESPONSE_ERROR             TLV_RESPONSE = 0x01
	TLV_RESPONSE_INVALID_PARAMETER TLV_RESPONSE = 0x02
)
