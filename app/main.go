package main

import (
	"fmt"
	"net"
	"strings"
)

func ParseDomain(data []byte) (string, int) {
	var domainParts []string
	position := 12 // DNS header is 12 bytes long

	for {
		// Read the length of the next label
		labelLength := int(data[position])
		if labelLength == 0 {
			// End of the domain name (null byte)
			position++ // Move past the null byte
			break
		}

		// Extract the label and add it to the domain parts
		position++
		label := string(data[position : position+labelLength])
		domainParts = append(domainParts, label)

		// Move position to the next label
		position += labelLength
	}

	// Join the labels with dots to form the full domain name
	domain := strings.Join(domainParts, ".")

	return domain, position
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:2053")
	if err != nil {
		fmt.Println("Failed to resolve UDP address:", err)
		return
	}

	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Failed to bind to address:", err)
		return
	}
	defer udpConn.Close()

	buf := make([]byte, 512)

	for {
		size, source, err := udpConn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error receiving data:", err)
			break
		}

		receivedData := string(buf[:size])
		fmt.Printf("Received %d bytes from %s: %s\n", size, source, receivedData)

		response := DNSResponse(buf)

		_, err = udpConn.WriteToUDP(response, source)
		if err != nil {
			fmt.Println("Failed to send response:", err)
		}
	}
}
