package fossil

import "testing"

//***** Testing *****//

func TestNewClient(t *testing.T) {
	// Panic check
	NewClient("www.example.com", "TESTTOKEN")
}

func TestNewApplication(t *testing.T) {
	// Panic check
	NewApplication("www.example.com", "TESTTOKEN")
}
