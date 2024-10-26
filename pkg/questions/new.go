package questions

func NewDNSQuestion(res []byte, domain string) []byte {
	question := DNSQuestion{
		Name:  domain,
		Type:  1, // Type A
		Class: 1, // IN (Internet)
	}

	return append(res, question.Serialize()...)
}
