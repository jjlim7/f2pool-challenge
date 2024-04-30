package handlers_test

import (
	"backend/internals/handlers"
	"testing"
)

func TestLookupIPv4(t *testing.T) {
	// Mock domain for testing
	domain := "google.com"

	// Call the lookupIPv4 method
	ipv4Addresses, err := handlers.LookupIPv4(domain)

	// Check if there's an error
	if err != nil {
		t.Errorf("lookupIPv4 error: %v", err)
	}

	// Assert that we received some IPv4 addresses
	if len(ipv4Addresses) == 0 {
		t.Error("lookupIPv4 did not return any IPv4 addresses")
	}
}
