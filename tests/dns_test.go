package main

import (
	"net"
	"testing"
)

func TestDNSLookup(t *testing.T) {
	// Test case 1: Valid domain name
	ip, err := DNSLookup("example.com")
	if err != nil {
		t.Errorf("DNSLookup() returned an error: %v", err)
	}
	if ip == nil {
		t.Error("DNSLookup() returned a nil IP")
	}

	// Test case 2: Invalid domain name
	ip, err = DNSLookup("nonexistentdomain")
	if err == nil {
		t.Error("DNSLookup() did not return an error for an invalid domain name")
	}
	if ip != nil {
		t.Error("DNSLookup() returned a non-nil IP for an invalid domain name")
	}
}