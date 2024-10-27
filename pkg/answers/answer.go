package answers

import (
	"bytes"
	"encoding/binary"
	"strings"
)

type DNSAnswer struct {
	Name    string
	Type    uint16
	Class   uint16
	TTL     uint32
	DataLen uint16
	Address [4]byte // IPv4 address (A record)
}

func (a *DNSAnswer) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Convert the name (example.com) into DNS format (length-prefixed labels)
	labels := strings.Split(a.Name, ".")
	for _, label := range labels {
		buf.WriteByte(byte(len(label)))
		buf.WriteString(label)
	}
	buf.WriteByte(0) // End of the domain name

	// Write Type, Class, TTL, Data length, and Address
	binary.Write(buf, binary.BigEndian, a.Type)
	binary.Write(buf, binary.BigEndian, a.Class)
	binary.Write(buf, binary.BigEndian, a.TTL)
	binary.Write(buf, binary.BigEndian, a.DataLen)
	binary.Write(buf, binary.BigEndian, a.Address)

	return buf.Bytes()
}
