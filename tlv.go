package ble


type TLV_RESPONSE_STATUS byte

// https://gopro.github.io/OpenGoPro/ble/protocol/data_protocol.html#command-response
const (
	TLV_RESPONSE_SUCCESS           TLV_RESPONSE_STATUS = 0x00
	TLV_RESPONSE_ERROR             TLV_RESPONSE_STATUS = 0x01
	TLV_RESPONSE_INVALID_PARAMETER TLV_RESPONSE_STATUS = 0x02
)

type TlvResponse struct {
	CommandID byte
	Status    TLV_RESPONSE_STATUS
	Payload   []byte
}

func decodeTlvResponse(b []byte) {
	// # First byte is the length of this response.
	// # Second byte is the ID
	// # Third byte is the status
	// # The remainder is the payload
	// payload = data[3 : length + 1]
	// var payload []byte

	// length, cmdId, status := int(b[0]), b[1], b[2]
	// if len(b) > 3 {
	// 	payload = b[3 : length+1]

	// }
}
