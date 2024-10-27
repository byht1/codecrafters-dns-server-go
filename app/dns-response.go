package main

import (
	"context"
	"net"

	"github.com/codecrafters-io/dns-server-starter-go/pkg/answers"
	"github.com/codecrafters-io/dns-server-starter-go/pkg/headers"
	"github.com/codecrafters-io/dns-server-starter-go/pkg/helpers"
	"github.com/codecrafters-io/dns-server-starter-go/pkg/questions"
)

func DNSResponse(data []byte, resolver *net.Resolver) (response []byte) {
	id := helpers.GetDNSId(data)
	flag := helpers.GetDNSFlags(data)
	qbCount := helpers.GetQueryCount(data)

	response = headers.DNSSimpleHeaderResponse(response, id, flag, qbCount)

	questionsReq, _ := questions.ParseQuestions(data, 12)
	for _, question := range questionsReq {
		response = append(response, question.Serialize()...)
	}

	for _, question := range questionsReq {
		ips, err := resolver.LookupIP(context.Background(), "ip4", question.Name)
		if err != nil {
			continue
		}

		response = answers.NewDNSAnswer(response, question.Name, ips[0].To4())
	}

	return
}
