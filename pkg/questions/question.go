package questions

import (
	"bytes"
	"encoding/binary"
	"strings"
)

type DNSQuestion struct {
	Name  string
	Type  uint16
	Class uint16
}

func (q *DNSQuestion) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Convert the name (example.com) into DNS format (length-prefixed labels)
	labels := strings.Split(q.Name, ".")
	for _, label := range labels {
		buf.WriteByte(byte(len(label)))
		buf.WriteString(label)
	}
	buf.WriteByte(0) // End of the domain name

	// Write Type and Class
	binary.Write(buf, binary.BigEndian, q.Type)
	binary.Write(buf, binary.BigEndian, q.Class)

	return buf.Bytes()
}
