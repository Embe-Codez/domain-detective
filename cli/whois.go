package cli

import (
	"fmt"

	whois "github.com/undiabler/golang-whois"
)

type WHOISClient struct {
	Domain string
	Result string
}

func NewWHOISClient(domain string) *WHOISClient {
	return &WHOISClient{
		Domain: domain,
	}
}

func (wc *WHOISClient) PerformWHOISLookup() error {
	var err error
	wc.Result, err = whois.GetWhois(wc.Domain)
	if err != nil {
		return err
	}
	return nil
}

func (wc *WHOISClient) PrintWHOISResult() {
	fmt.Println("WHOIS Result for", wc.Domain)
	fmt.Println(wc.Result)
}

func WHOISLookup(domain string) {
	client := NewWHOISClient(domain)
	err := client.PerformWHOISLookup()
	if err != nil {
		fmt.Println("WHOIS lookup failed:", err)
		return
	}

	client.PrintWHOISResult()
}