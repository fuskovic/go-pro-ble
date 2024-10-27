package ble

import "strconv"

// https://gopro.github.io/OpenGoPro/ble/protocol/data_protocol.html#commands
const WifiApToggleCmdID = 0x17

type TLV_RESPONSE byte

// https://gopro.github.io/OpenGoPro/ble/protocol/data_protocol.html#command-response
const (
	TLV_RESPONSE_SUCCESS           TLV_RESPONSE = 0x00
	TLV_RESPONSE_ERROR             TLV_RESPONSE = 0x01
	TLV_RESPONSE_INVALID_PARAMETER TLV_RESPONSE = 0x02
)

type TlvResponse struct {
}

func decodeTlvResponse(b []byte) {
	// # First byte is the length of this response.
	// # Second byte is the ID
	// # Third byte is the status
	// # The remainder is the payload
	// payload = data[3 : length + 1]
	var payload []byte
	
	length, cmdId, status := int(b[0]), b[1], b[2]
	if len(b) > 3 {
		payload = b[3:length+1]
		
	}
}
