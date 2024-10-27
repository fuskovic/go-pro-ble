package ble

import (
	"strings"
)

// GP-XXXX is shorthand for GoPro’s 128-bit UUID: b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b
// https://gopro.github.io/OpenGoPro/ble/protocol/ble_setup.html#ble-characteristics
var baseFormat = "b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b"

// format takes 'b5f9XXXX-aa8d-11e3-9046-0002a5d5c51b' and replaces 'XXXX' with s.
func format(s string) string {
	return strings.Replace(baseFormat, "XXXX", s, 1)
}
