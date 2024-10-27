package headers

type HeaderFlags struct {
	Flag   uint16 // Flag bits
	QR     uint16 // Query/Response flag
	OPCODE uint16 // Operation code
	AA     uint16 // Authoritative Answer flag
	TC     uint16 // Truncated message flag
	RD     uint16 // Recursion Desired flag
	RA     uint16 // Recursion Available flag
	Z      uint16 // Reserved (must be 0)
	RCODE  uint16 // Response code
}

func (f *HeaderFlags) ToUint16() uint16 {
	var flags uint16

	flags |= (f.QR << 15)
	flags |= (f.OPCODE << 11)
	flags |= (f.AA << 10)
	flags |= (f.TC << 9)
	flags |= (f.RD << 8)
	flags |= (f.RA << 7)
	flags |= (f.Z << 4)
	flags |= (f.RCODE & 0xF)

	return flags
}
