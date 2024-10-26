package ble

// https://gopro.github.io/OpenGoPro/ble/protocol/ble_setup.html#ble-characteristics
var (
	wifiAccessPoint   = Service(format("0001"))
	cameraNetworkMgmt = Service(format("0090"))
	ctrlAndQuery      = Service("0000fea6-0000-1000-8000-00805f9b34fb")
	internal          = Service("00002a19-0000-1000-8000-00805f9b34fb")
)

var Services = []Service{wifiAccessPoint, cameraNetworkMgmt, ctrlAndQuery}

// Service is the ble service registered on the device.
type Service string

// Prints the human-readable name for the service.
func (s Service) Name() string {
	var name string
	switch s {
	case wifiAccessPoint:
		name = "wifi-access-point"
	case cameraNetworkMgmt:
		name = "camera-network-management"
	case ctrlAndQuery:
		name = "control-and-query"
	case internal:
		name = "internal"
	default:
		name = "unknown"
	}
	return name
}
