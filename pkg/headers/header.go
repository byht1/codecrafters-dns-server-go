package headers

import (
	"bytes"
	"encoding/binary"
)

type DNSHeader struct {
	ID      uint16 // The request ID must match the response ID.
	Flags   uint16 // Flags to indicate the type of message (request or response)
	QDCount uint16 // Number of requests
	ANCount uint16 // Number of responses
	NSCount uint16 // Number of authoritative server records
	ARCount uint16 // Number of additional records
}

func (h *DNSHeader) Serialize() []byte {
	buf := new(bytes.Buffer)

	binary.Write(buf, binary.BigEndian, h.ID)
	binary.Write(buf, binary.BigEndian, h.Flags)
	binary.Write(buf, binary.BigEndian, h.QDCount)
	binary.Write(buf, binary.BigEndian, h.ANCount)
	binary.Write(buf, binary.BigEndian, h.NSCount)
	binary.Write(buf, binary.BigEndian, h.ARCount)

	return buf.Bytes()
}
