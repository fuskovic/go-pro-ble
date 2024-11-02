package __

import "google.golang.org/protobuf/proto"

type OPERATION int

const (
	OPERATION_REQUEST OPERATION = iota
	OPERATION_RESPONSE
	OPERATION_NOTIFICATION
)

type ID interface {
	FeatureID() byte
	ActionID() byte
	Operation() byte
	DataStructure() proto.Message
}

type protobufID struct {
	featureID     byte
	actionID      byte
	op            OPERATION
	dataStructure proto.Message
}

func (i protobufID) FeatureID() byte              { return i.featureID }
func (i protobufID) ActionID() byte               { return i.actionID }
func (i protobufID) Operation() OPERATION         { return i.op }
func (i protobufID) DataStructure() proto.Message { return i.dataStructure }

type protobufIDs []protobufID

var (
	IDs = []protobufID{StartScanRequest, GetApEntries}
)

var (
	StartScanRequest = protobufID{
		featureID:     0x02,
		actionID:      0x02,
		op:            OPERATION_REQUEST,
		dataStructure: new(RequestStartScan),
	}

	GetApEntries = protobufID{
		featureID:     0x02,
		actionID:      0x03,
		op:            OPERATION_REQUEST,
		dataStructure: new(RequestGetApEntries),
	}
)
