package ble

// Notification represents the final accumulated response for a given command.
type Notification interface {
	CommandID() COMMAND_ID
	Status() TLV_RESPONSE_STATUS
	Payload() []byte
	// Match confirms if the notification came from the targetCmdID and status matches the targetStatus
	Match(targetCmdID COMMAND_ID, targetStatus TLV_RESPONSE_STATUS) bool
}

type notification struct {
	cmdID   COMMAND_ID
	status  TLV_RESPONSE_STATUS
	payload *payload
}

func (n notification) CommandID() COMMAND_ID       { return n.cmdID }
func (n notification) Status() TLV_RESPONSE_STATUS { return n.status }
func (n notification) Payload() []byte             { return n.payload.bytes }
func (n notification) Match(targetCmdID COMMAND_ID, targetStatus TLV_RESPONSE_STATUS) bool {
	return n.CommandID() == targetCmdID && n.Status() == targetStatus
}
