package ble

import (
	"bytes"

	"github.com/fuskovic/go-pro-ble/internal/packet"
)

type Notification interface {
	CommandID() COMMAND_ID
	Status() TLV_RESPONSE_STATUS
	Payload() *packet.Payload
}

type humanReadableNotification struct {
	cmdID   COMMAND_ID
	status  TLV_RESPONSE_STATUS
	payload *packet.Payload
}

func (n humanReadableNotification) CommandID() COMMAND_ID       { return n.cmdID }
func (n humanReadableNotification) Status() TLV_RESPONSE_STATUS { return n.status }
func (n humanReadableNotification) Payload() *packet.Payload    { return n.payload }

type rawNotification struct {
	buf *bytes.Buffer
}