package main

import "testing"

func TestWHOISLookup(t *testing.T) {
	// Test case 1: Valid domain name
	result, err := WHOISLookup("example.com")
	if err != nil {
		t.Errorf("WHOISLookup() returned an error: %v", err)
	}
	if result == "" {
		t.Error("WHOISLookup() returned an empty result")
	}

	// Test case 2: Invalid domain name
	result, err = WHOISLookup("nonexistentdomain")
	if err == nil {
		t.Error("WHOISLookup() did not return an error for an invalid domain name")
	}
	if result != "" {
		t.Error("WHOISLookup() returned non-empty result for an invalid domain name")
	}
}