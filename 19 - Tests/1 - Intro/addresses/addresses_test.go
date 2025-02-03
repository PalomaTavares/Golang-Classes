// UNIT TEST
package addresses_test

import (
	"testeIntro/addresses"
	"testing"
)

type testCenary struct {
	receivedAddress string
	expectedReturn  string
}

func TestAddressType(t *testing.T) {
	//Test funcName		ponter in testing

	t.Parallel()

	testCenaries := []testCenary{
		{"Ocean Avenue", "Avenue"},
		{"Fools Street", "Street"},
		{"Hell's highway", "Highway"},
		{"Broken Dreams's Boulevard", "invalid type"},
		{"MIDNIGHT HIGHWAY", "Highway"},
		{"abey road", "Road"},
		{"", "invalid type"}, //if commented 92.3% coverage go test --cover or go test --coverprofile result.txt
		//to read better go tool cover --func=result.txt
	}

	for _, cenary := range testCenaries {
		receivedAddressType := addresses.AddressType(cenary.receivedAddress)
		if receivedAddressType != cenary.expectedReturn {
			t.Errorf("type %s is different from expected: %s", receivedAddressType, cenary.expectedReturn)
		}
	}
}

func TestFoo(t *testing.T) {
	t.Parallel()
	if 1 > 2 {
		t.Errorf("Broke :()")
	}
}
