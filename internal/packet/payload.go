package packet

type Payload struct {
	rawBytes []byte
}

func (p *Payload) Bytes() []byte { return p.rawBytes }

// The BLE protocol limits packet size to 20 bytes per packet.
// To accommodate, GoPro cameras use start and continuation packets
// to serialize larger payloads. If a message is less than 20 bytes, it can be
// sent with a single packet containing the start packet header. Otherwise, it
// must be split into multiple packets with the first packet containing
// a start packet header and subsequent packets containing continuation packet headers.
// https://gopro.github.io/OpenGoPro/ble/protocol/data_protocol.html#packetization
// https://gopro.github.io/OpenGoPro/tutorials/parse-ble-responses#accumulating-the-response
func (p *Payload) Accumulate(buf []byte) {
	if buf[0] == CONTINUATION_MASK.Byte() {
		// pop the first byte
		buf = buf[1:]
	} else {
		// 	<< is used for "times 2" and >> is for "divided by 2" - and the number after it is how many times.
		// So n << x is "n times 2, x times". And y >> z is "y divided by 2, z times".
		// For example, 1 << 5 is "1 times 2, 5 times" or 32. And 32 >> 5 is "32 divided by 2, 5 times" or 1.
		header := ((buf[0] & HEADER_MASK.Byte()) >> 5)
		switch header {
		case GENERAL.Byte():
			buf = buf[1:]
		case EXT_13.Byte():
			buf = buf[2:]
		case EXT_16.Byte():
			buf = buf[3:]
		}
	}

	// # Append payload to buffer and update remaining / complete
	p.rawBytes = append(p.rawBytes, buf...)
}
