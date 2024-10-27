package questions

import (
	"encoding/binary"

	"github.com/codecrafters-io/dns-server-starter-go/pkg/helpers"
)

func ParseQuestions(data []byte, offset int) ([]DNSQuestion, error) {
	var questions []DNSQuestion

	for i := 0; i < int(binary.BigEndian.Uint16(data[4:6])); i++ {
		var question DNSQuestion

		// Parse the name
		name, newOffset, err := helpers.ParseName(data, offset)
		if err != nil {
			return nil, err
		}
		question.Name = name
		offset = newOffset

		// Parse Type and Class
		question.Type = binary.BigEndian.Uint16(data[offset : offset+2])
		offset += 2
		question.Class = binary.BigEndian.Uint16(data[offset : offset+2])
		offset += 2

		questions = append(questions, question)
	}

	return questions, nil
}
