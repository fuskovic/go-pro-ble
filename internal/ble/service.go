package ble

// Service is the ble service registered on the device.
type Service struct {
	uuid string
	name string
}

// read-only fields
func (s Service) Uuid() string { return s.uuid }
func (s Service) Name() string { return s.name }

var Services = []Service{
	wifiAccessPoint,
	cameraNetworkMgmt,
	ctrlAndQuery,
	internal,
}

// https://gopro.github.io/OpenGoPro/ble/protocol/ble_setup.html#ble-characteristics
var (
	wifiAccessPoint = Service{
		uuid: format("0001"),
		name: "wifi-access-point",
	}
	cameraNetworkMgmt = Service{
		uuid: format("0090"),
		name: "camera-network-management",
	}
	ctrlAndQuery = Service{
		uuid: "0000fea6-0000-1000-8000-00805f9b34fb",
		name: "control-and-query",
	}
	internal = Service{
		uuid: "00002a19-0000-1000-8000-00805f9b34fb",
		name: "internal",
	}
)
