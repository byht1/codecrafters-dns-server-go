package main

import (
	"github.com/codecrafters-io/dns-server-starter-go/pkg/answers"
	"github.com/codecrafters-io/dns-server-starter-go/pkg/headers"
	"github.com/codecrafters-io/dns-server-starter-go/pkg/helpers"
	"github.com/codecrafters-io/dns-server-starter-go/pkg/questions"
)

func DNSResponse(data []byte) (response []byte) {
	id := helpers.GetDNSId(data)
	flag := helpers.GetDNSFlags(data)
	qbCount := helpers.GetQueryCount(data)

	response = headers.DNSSimpleHeaderResponse(response, id, flag, qbCount)

	questionsReq, _ := questions.ParseQuestions(data, 12)
	for _, question := range questionsReq {
		response = append(response, question.Serialize()...)
	}

	for _, question := range questionsReq {
		response = answers.NewDNSAnswer(response, question.Name, "8.8.8.8")
	}

	return
}
