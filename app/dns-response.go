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
	domain, _ := helpers.ParseDomain(data)

	response = headers.DNSSimpleHeaderResponse(response, id, flag)
	response = questions.NewDNSQuestion(response, domain)
	response = answers.NewDNSAnswer(response, domain, "8.8.8.8")

	return
}
