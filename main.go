package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/embe-codez/DomainDetective/cli"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Prompt user to enter a domain name.
	fmt.Print("Enter a domain name: ")
	domain, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	domain = strings.TrimSpace(domain)

	// Prompt user to enter a record type (A, CNAME, TXT).
	fmt.Print("Enter a record type (A, CNAME, TXT): ")
	recordType, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	recordType = strings.TrimSpace(recordType)

	// Perform DNS lookup
	cli.DNSLookup(domain, recordType)

	// Perform WHOIS lookup
	cli.WHOISLookup(domain)

	// Retrieve certificate info
	cli.CertificateInfo(domain)
}