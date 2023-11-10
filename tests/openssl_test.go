package main

import "testing"

func TestCertificateInfo(t *testing.T) {
	// Test case 1: Valid certificate
	info, err := CertificateInfo("example.com")
	if err != nil {
		t.Errorf("CertificateInfo() returned an error: %v", err)
	}
	if info == "" {
		t.Error("CertificateInfo() returned an empty string")
	}

	// Test case 2: Invalid certificate
	info, err = CertificateInfo("nonexistentdomain")
	if err == nil {
		t.Error("CertificateInfo() did not return an error for an invalid certificate")
	}
	if info != "" {
		t.Error("CertificateInfo() returned non-empty info for an invalid certificate")
	}
}