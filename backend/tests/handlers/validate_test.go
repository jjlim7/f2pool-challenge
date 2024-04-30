package handlers_test

import (
	"testing"

	"backend/internals/handlers"
)

func TestIsValidIPv4(t *testing.T) {
	// Test valid IPv4 address
	validIP := "192.168.1.1"
	isValid := handlers.IsValidIPv4(validIP)
	if !isValid {
		t.Errorf("%s is a valid IPv4 address, but isValidIPv4 returned false", validIP)
	}

	// Test invalid IPv4 address
	invalidIP := "invalid_ip_address"
	isValid = handlers.IsValidIPv4(invalidIP)
	if isValid {
		t.Errorf("%s is an invalid IPv4 address, but isValidIPv4 returned true", invalidIP)
	}
}
