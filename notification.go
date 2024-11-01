package ble

// Notification represents the final accumulated response for a given command.
type Notification interface {
	CommandID() COMMAND_ID
	Status() TLV_RESPONSE_STATUS
	Payload() []byte
}

type notification struct {
	cmdID   COMMAND_ID
	status  TLV_RESPONSE_STATUS
	payload *payload
}

func (n notification) CommandID() COMMAND_ID       { return n.cmdID }
func (n notification) Status() TLV_RESPONSE_STATUS { return n.status }
func (n notification) Payload() []byte             { return n.payload.bytes }