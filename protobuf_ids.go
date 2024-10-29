package ble

import pb "github.com/fuskovic/ble/protos"

type OpElement int

func (o OpElement) String() string {
	var s string
	switch o {
	case OpElementNotification:
		s = "notification"
	case OpElementRequest:
		s = "request"
	case OpElementResponse:
		s = "response"
	default:
		s = "unknown"
	}
	return s
}

const (
	OpElementUnknown OpElement = iota
	OpElementRequest
	OpElementResponse
	OpElementNotification
)

type ProtobufId[T any] struct {
	FeatureID byte
	ActionID  byte
	OpElement OpElement
	PbMessage *T
}

// Notifications
var (
	NotifStartScanning = ProtobufId[pb.NotifStartScanning]{
		FeatureID: 0x02,
		ActionID:  0x0B,
		OpElement: OpElementNotification,
		PbMessage: new(pb.NotifStartScanning),
	}

// 0x02	0x0B	"Scan for Access Points" Notification	NotifStartScanning
// 0x02	0x0C	"Connect to Provisioned Access Point" Notification	NotifProvisioningState
// 0x02	0x0C	"Connect to a New Access Point" Notification	NotifProvisioningState
)

// Feature ID	Action ID	Operation Element	Protobuf Message
// 0x02	0x02	"Scan for Access Points" Request	RequestStartScan
// 0x02	0x03	"Get AP Scan Results" Request	RequestGetApEntries
// 0x02	0x04	"Connect to Provisioned Access Point" Request	RequestConnect
// 0x02	0x05	"Connect to a New Access Point" Request	RequestConnectNew
// // 0x02	0x0B	"Scan for Access Points" Notification	NotifStartScanning
// // 0x02	0x0C	"Connect to Provisioned Access Point" Notification	NotifProvisioningState
// // 0x02	0x0C	"Connect to a New Access Point" Notification	NotifProvisioningState
// 0x02	0x82	"Scan for Access Points" Response	ResponseStartScanning
// 0x02	0x83	"Get AP Scan Results" Response	ResponseGetApEntries
// 0x02	0x84	"Connect to Provisioned Access Point" Response	ResponseConnect
// 0x02	0x85	"Connect to a New Access Point" Response	ResponseConnectNew
// 0xF1	0x64	"Update Custom Preset" Request	RequestCustomPresetUpdate
// 0xF1	0x65	"Set COHN Setting" Request	RequestSetCOHNSetting
// 0xF1	0x66	"Clear COHN Certificate" Request	RequestClearCOHNCert
// 0xF1	0x67	"Create COHN Certificate" Request	RequestCreateCOHNCert
// 0xF1	0x69	"Set Camera Control" Request	RequestSetCameraControlStatus
// 0xF1	0x6B	"Set Turbo Transfer" Request	RequestSetTurboActive
// 0xF1	0x79	"Set Livestream Mode" Request	RequestSetLiveStreamMode
// 0xF1	0xE4	"Update Custom Preset" Response	ResponseGeneric
// 0xF1	0xE5	"Set COHN Setting" Response	ResponseGeneric
// 0xF1	0xE6	"Clear COHN Certificate" Response	ResponseGeneric
// 0xF1	0xE7	"Create COHN Certificate" Response	ResponseGeneric
// 0xF1	0xE9	"Set Camera Control" Response	ResponseGeneric
// 0xF1	0xEB	"Set Turbo Transfer" Response	ResponseGeneric
// 0xF1	0xF9	"Set Livestream Mode" Response	ResponseGeneric
// 0xF5	0x6D	"Get Last Captured Media" Request	RequestGetLastCapturedMedia
// 0xF5	0x6E	"Get COHN Certificate" Request	RequestCOHNCert
// 0xF5	0x6F	"Get COHN Status" Request	RequestGetCOHNStatus
// 0xF5	0x72	"Get Available Presets" Request	RequestGetPresetStatus
// 0xF5	0x74	"Get Livestream Status" Request	RequestGetLiveStreamStatus
// 0xF5	0xED	"Get Last Captured Media" Response	ResponseLastCapturedMedia
// 0xF5	0xEE	"Get COHN Certificate" Response	ResponseCOHNCert
// 0xF5	0xEF	"Get COHN Status" Response	NotifyCOHNStatus
// 0xF5	0xEF	"Get COHN Status" Notification	NotifyCOHNStatus
// 0xF5	0xF2	"Get Available Presets" Response	NotifyPresetStatus
// 0xF5	0xF3	"Get Available Presets" Notification	NotifyPresetStatus
// 0xF5	0xF4	"Get Livestream Status" Response	NotifyLiveStreamStatus
// 0xF5	0xF5	"Get Livestream Status" Notification	NotifyLiveStreamStatus
