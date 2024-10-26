package headers

func DNSSimpleHeaderResponse(res []byte, id uint16) []byte {
	header := DNSHeader{
		ID:      id,
		Flags:   0x8180,
		QDCount: 1,
		ANCount: 1,
		NSCount: 0,
		ARCount: 0,
	}

	return append(res, header.Serialize()...)
}
