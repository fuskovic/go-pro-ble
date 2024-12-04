package ble

import (
	"fmt"
	"log"

	pb "github.com/fuskovic/ble/protos"
	"google.golang.org/protobuf/proto"
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

func decodeTlvResponse(b []byte) error {
	for _, pb := range pb.IDs {
		if b[0] == pb.FeatureID() && b[1] == pb.FeatureID() {
			x := pb.DataStructure()
			if err := proto.Unmarshal(b, x); err != nil {
				return fmt.Errorf("failed to unmarshal data structure: %s", err)
			}
		}
	}
	// TODO:
	// Handle non-pb payload.
	//
	// Nope. It is a TLV response
	// if U == GP-0072 (Command) {
	// 	Parse message payload using Command Table with Command scheme
	// }
	// else if U == GP-0074 (Settings) {
	// 	Parse using Setting ID mapping with Command scheme
	// }
	// else if U == GP-0076 (Query) {
	// 	Parse message payload using Query Table with Query scheme
	// }
	return nil
}
