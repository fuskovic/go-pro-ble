package ble

// https://gopro.github.io/OpenGoPro/ble/protocol/ble_setup.html#ble-characteristics
var (
	wifiAccessPoint   = Service(format("0001"))
	cameraNetworkMgmt = Service(format("0090"))
	ctrlAndQuery      = Service(format("FEA6"))
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
	default:
		name = "unknown"
	}
	return name
}
