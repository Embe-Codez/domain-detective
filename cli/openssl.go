package cli

import (
	"crypto/tls"
	"fmt"
	"log"
)

type CertificateVars struct {
	Domain    string
	Subject   string
	Issuer    string
	NotBefore string
	NotAfter  string
}

func NewCertificateVars(domain string) *CertificateVars {
	return &CertificateVars{
		Domain: domain,
	}
}

func (c *CertificateVars) GatherInfo() error {
	conn, err := tls.Dial("tcp", c.Domain+":443", nil)
	if err != nil {
		return fmt.Errorf("failed to connect to %s: %w", c.Domain, err)
	}
	defer conn.Close()

	cert := conn.ConnectionState().PeerCertificates[0]

	c.Subject = cert.Subject.CommonName
	c.Issuer = cert.Issuer.CommonName
	c.NotBefore = cert.NotBefore.String()
	c.NotAfter = cert.NotAfter.String()

	return nil
}

func CertificateInfo(domain string) {
	c := NewCertificateVars(domain)
	err := c.GatherInfo()
	if err != nil {
		log.Fatal("Failed to gather certificate information:", err)
	}

	fmt.Println("Certificate information:")
	fmt.Println("Domain:", c.Domain)
	fmt.Println("Subject:", c.Subject)
	fmt.Println("Issuer:", c.Issuer)
	fmt.Println("Not Before:", c.NotBefore)
	fmt.Println("Not After:", c.NotAfter)
}