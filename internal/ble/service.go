package ble

var (
	wifiAccessPoint = Service(format("0001"))
	cameraMgmt      = Service(format("0090"))
	ctrlAndQuery    = Service(format("FEA6"))
)

var Services = []Service{wifiAccessPoint, cameraMgmt, ctrlAndQuery}

type Service string

func (s Service) Name() string {
	var name string
	switch s {
	case wifiAccessPoint:
		name = "wifi-access-point"
	case cameraMgmt:
		name = "camera-management"
	case ctrlAndQuery:
		name = "control-and-query"
	default:
		name = "unknown"
	}
	return name
}
