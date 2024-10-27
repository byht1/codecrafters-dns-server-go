package helpers

import (
	"encoding/binary"
	"strings"
)

func GetDNSId(data []byte) uint16 {
	return binary.BigEndian.Uint16(data[0:2])
}

func GetDNSFlags(data []byte) uint16 {
	return binary.BigEndian.Uint16(data[2:4])
}

func GetQueryCount(data []byte) uint16 {
	return binary.BigEndian.Uint16(data[4:6])
}

func ParseName(data []byte, offset int) (string, int, error) {
	var name string
	for {
		length := data[offset]
		offset++
		if length == 0 {
			break // End of the name
		}
		name += string(data[offset:offset+int(length)]) + "."
		offset += int(length)
	}
	name = strings.TrimSuffix(name, ".") // Remove the trailing dot
	return name, offset, nil
}
