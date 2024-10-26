package main

import (
	"github.com/codecrafters-io/dns-server-starter-go/pkg/headers"
	"github.com/codecrafters-io/dns-server-starter-go/pkg/helpers"
	"github.com/codecrafters-io/dns-server-starter-go/pkg/questions"
)

func DNSResponse(data []byte) (response []byte) {
	id := helpers.GetDNSId(data)
	domain, _ := helpers.ParseDomain(data)

	response = headers.DNSSimpleHeaderResponse(response, id)
	response = questions.NewDNSQuestion(response, domain)

	return
}
