package cli

import (
	"fmt"
	"net"
)

type DNSClient struct {
	Domain      string
	RecordType  string
	Records     []string
}

func NewDNSClient(domain, recordType string) *DNSClient {
	return &DNSClient{
		Domain:     domain,
		RecordType: recordType,
	}
}

func (dc *DNSClient) PerformDNSLookup() error {
	switch dc.RecordType {
	case "A":
		ips, err := net.LookupIP(dc.Domain)
		if err != nil {
			return err
		}
		dc.Records = make([]string, len(ips))
		for i, ip := range ips {
			dc.Records[i] = ip.String()
		}
	case "CNAME":
		cname, err := net.LookupCNAME(dc.Domain)
		if err != nil {
			return err
		}
		dc.Records = []string{cname}
	case "TXT":
		txts, err := net.LookupTXT(dc.Domain)
		if err != nil {
			return err
		}
		dc.Records = txts
	default:
		return fmt.Errorf("unsupported record type: %s", dc.RecordType)
	}
	return nil
}

func DNSLookup(domain, recordType string) {
	client := NewDNSClient(domain, recordType)
	err := client.PerformDNSLookup()
	if err != nil {
		fmt.Println("DNS lookup failed:", err)
		return
	}

	client.PrintDNSResult()
}

func (dc *DNSClient) PrintDNSResult() {
	fmt.Printf("%s Records for %s:\n", dc.RecordType, dc.Domain)
	for _, record := range dc.Records {
		fmt.Println(record)
	}
}