package packet

import "log"

const (
	// We use HEADER_MASK to distinguish the header bytes from the payload.
	headerMask = 0b01100000
	// When sending or receiving a message that is longer than 20 bytes, the message must be
	// split into N packets with packet 1 containing a start packet header and packets 2..N
	// containing a continuation packet header.
	continuationMask = 0b10000000
)

const (
	// Messages that are 20 bytes or fewer can be sent or received using the following format:
	// https://gopro.github.io/OpenGoPro/ble/protocol/data_protocol.html#general-5-bit-packets
	general = 0b00

	// Messages that are 8191 bytes or fewer can be sent or received using the following format:
	// https://gopro.github.io/OpenGoPro/ble/protocol/data_protocol.html#extended-13-bit-packets
	// Always use Extended (13-bit) packet headers when sending messages to avoid having to work with multiple packet header formats.
	ext13 = 0b01

	// If a message is 8192 bytes or longer, the camera will respond using the following format:
	// https://gopro.github.io/OpenGoPro/ble/protocol/data_protocol.html#extended-16-bit-packets
	// This format can not be used for sending messages to the camera. It is only used to receive messages.
	ext16 = 0b10
)

type Payload struct {
	rawBytes []byte
	offset   int
}

// Offset should be used by the caller to distinguish between where
// the payload starts. Which should be after the command id and status.
func (p *Payload) Offset() int { return p.offset }

// Bytes should only be called after all the necessary bluetooth packets
// have been accumulated.
func (p *Payload) Bytes() []byte { return p.rawBytes }

// The BLE protocol limits packet size to 20 bytes per packet.
// To accommodate, GoPro cameras use start and continuation packets
// to serialize larger payloads. If a message is less than 20 bytes, it can be
// sent with a single packet containing the start packet header. Otherwise, it
// must be split into multiple packets with the first packet containing
// a start packet header and subsequent packets containing continuation packet headers.
// https://gopro.github.io/OpenGoPro/tutorials/parse-ble-responses#accumulating-the-response
func (p *Payload) Accumulate(buf []byte) {
	if buf[0] == continuationMask {
		// pop the first byte
		log.Println("received continuation packet")
		buf = buf[1:]
	} else {
		// 	<< is used for "times 2" and >> is for "divided by 2" - and the number after it is how many times.
		// So n << x is "n times 2, x times". And y >> z is "y divided by 2, z times".
		// For example, 1 << 5 is "1 times 2, 5 times" or 32. And 32 >> 5 is "32 divided by 2, 5 times" or 1.
		header := ((buf[0] & headerMask) >> 5)
		switch header {
		case general:
			log.Println("received general packet")
			buf = buf[1:]
		case ext13:
			p.offset = 2
			log.Println("received extended-13 packet")
			buf = buf[2:]
		case ext16:
			p.offset = 2
			log.Println("received extended-16 packet")
			buf = buf[3:]
		}
	}
	// append payload to buffer and update remaining / complete
	p.rawBytes = append(p.rawBytes, buf...)
}
