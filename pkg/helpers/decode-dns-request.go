package helpers

import (
	"encoding/binary"
	"strings"
)

func GetDNSId(data []byte) uint16 {
	return binary.BigEndian.Uint16(data[0:2])
}

func ParseDomain(data []byte) (string, int) {
	var domainParts []string
	position := 12 // DNS header is 12 bytes long

	for {
		// Read the length of the next label
		labelLength := int(data[position])
		if labelLength == 0 {
			// End of the domain name (null byte)
			position++ // Move past the null byte
			break
		}

		// Extract the label and add it to the domain parts
		position++
		label := string(data[position : position+labelLength])
		domainParts = append(domainParts, label)

		// Move position to the next label
		position += labelLength
	}

	// Join the labels with dots to form the full domain name
	domain := strings.Join(domainParts, ".")

	return domain, position
}
