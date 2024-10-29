package packet

import "log"

type Payload struct {
	rawBytes       []byte
	bytesRemaining int
	received       bool
}

func (p *Payload) Complete() bool { return p.received }

// The BLE protocol limits packet size to 20 bytes per packet.
// To accommodate, GoPro cameras use start and continuation packets
// to serialize larger payloads. If a message is less than 20 bytes, it can be
// sent with a single packet containing the start packet header. Otherwise, it
// must be split into multiple packets with the first packet containing
// a start packet header and subsequent packets containing continuation packet headers.
// https://gopro.github.io/OpenGoPro/ble/protocol/data_protocol.html#packetization
// https://gopro.github.io/OpenGoPro/tutorials/parse-ble-responses#accumulating-the-response
func NewPayload(buf []byte) *Payload {
	p := new(Payload)
	if buf[0] == CONTINUATION_MASK.Byte() {
		log.Println("continuation packet received")
		// pop the first byte
		buf = buf[1:]
	} else {
		// This is a new packet so start with an empty byte array
		p.rawBytes = []byte{}
		header := ((buf[0] & HEADER_MASK.Byte()) >> 5)
		switch header {
		case GENERAL.Byte():
			log.Println("general packet received")
			p.bytesRemaining = int(buf[0] & GENERAL_LEN_MASK.Byte())
			buf = buf[1:]
		case EXT_13.Byte():
			log.Println("extended 13-bit packet received")
			p.bytesRemaining = int(((buf[0] & EXTENDED_13_BIT_MASK.Byte()) << 8) + buf[1])
			buf = buf[2:]
		case EXT_16.Byte():
			log.Println("extended 16-bit packet received")
			p.bytesRemaining = int((buf[1] << 8) + buf[2])
			buf = buf[3:]
		}
	}

	// # Append payload to buffer and update remaining / complete
	log.Printf("appending %d bytes", len(buf))
	p.rawBytes = append(p.rawBytes, buf...)
	p.bytesRemaining -= len(buf)
	log.Printf("bytes-remaining: %d\n", p.bytesRemaining)
	return p
}
