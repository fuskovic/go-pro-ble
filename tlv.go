package ble

import (
	"log"
)

type TLV_RESPONSE_STATUS byte

func (r TLV_RESPONSE_STATUS) String() string {
	var s string
	switch r {
	case TLV_RESPONSE_SUCCESS:
		s = "success"
	case TLV_RESPONSE_ERROR:
		s = "error"
	case TLV_RESPONSE_INVALID_PARAMETER:
		s = "invalid-parameter"
	default:
		log.Printf("unknown response: %d\n", r)
		s = "unknown"
	}
	return s
}

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
