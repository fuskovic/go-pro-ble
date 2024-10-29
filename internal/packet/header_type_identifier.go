package packet

type HEADER_TYPE_IDENTIFIER int

func (h HEADER_TYPE_IDENTIFIER) Byte() byte { return byte(h) }

const (
	// Messages that are 20 bytes or fewer can be sent or received using the following format:
	// https://gopro.github.io/OpenGoPro/ble/protocol/data_protocol.html#general-5-bit-packets
	GENERAL HEADER_TYPE_IDENTIFIER = 0b00

	// Messages that are 8191 bytes or fewer can be sent or received using the following format:
	// https://gopro.github.io/OpenGoPro/ble/protocol/data_protocol.html#extended-13-bit-packets
	// Always use Extended (13-bit) packet headers when sending messages to avoid having to work with multiple packet header formats.
	EXT_13 HEADER_TYPE_IDENTIFIER = 0b01

	// If a message is 8192 bytes or longer, the camera will respond using the following format:
	// https://gopro.github.io/OpenGoPro/ble/protocol/data_protocol.html#extended-16-bit-packets
	// This format can not be used for sending messages to the camera. It is only used to receive messages.
	EXT_16 HEADER_TYPE_IDENTIFIER = 0b10

	RESERVED HEADER_TYPE_IDENTIFIER = 0b11
)