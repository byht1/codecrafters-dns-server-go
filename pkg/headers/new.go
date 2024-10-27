package headers

func DNSSimpleHeaderResponse(res []byte, id uint16, flag uint16) []byte {
	opcode := (flag >> 11) & 0xF
	rcode := flag & 0xF // 0xF = 0000 1111
	if opcode != 0 {
		rcode = 4 // 4 = Not Implemented
	}

	header := DNSHeader{
		ID:      id,
		QDCount: 1,
		ANCount: 1,
		NSCount: 0,
		ARCount: 0,
		Flags: HeaderFlags{
			QR:     1,
			OPCODE: opcode,
			AA:     0,
			TC:     0,
			RD:     (flag >> 8) & 0x1,
			RA:     1,
			Z:      0,
			RCODE:  rcode,
		},
	}

	return append(res, header.Serialize()...)
}
