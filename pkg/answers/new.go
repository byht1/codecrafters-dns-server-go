package answers

import "net"

func NewDNSAnswer(res []byte, domain string, ipAddress string) []byte {
	const dataLen uint16 = 4 // IPv4 address is 4 bytes

	ip := net.ParseIP(ipAddress).To4()

	question := DNSAnswer{
		Name:    domain,
		Type:    1,  // Type A
		Class:   1,  // IN (Internet)
		TTL:     60, // 60 seconds TTL
		DataLen: dataLen,
		Address: [dataLen]byte{ip[0], ip[1], ip[2], ip[3]},
	}

	return append(res, question.Serialize()...)
}
