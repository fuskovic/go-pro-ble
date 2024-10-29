package packet

// MASK is a mask used to differentiate between different packet headers.
type MASK int

func (m MASK) Byte() byte { return byte(m) }

const (
	// When sending or receiving a message that is longer than 20 bytes, the message must be
	// split into N packets with packet 1 containing a start packet header and packets 2..N
	// containing a continuation packet header.
	//
	// NOTE: Counters start at 0x0 and reset after 0xF.
	CONTINUATION_MASK    MASK = 0b10000000
	HEADER_MASK          MASK = 0b01100000
	GENERAL_LEN_MASK     MASK = 0b00011111
	EXTENDED_13_BIT_MASK MASK = 0b00011111
)