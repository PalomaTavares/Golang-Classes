package addresses

import (
	"strings"
)

// address type verifies if the addres has a valid type
func AddressType(address string) string {
	validTypes := []string{"street", "avenue", "road", "highway"}

	addressLower := strings.ToLower(address)

	words := strings.Split(addressLower, " ")

	// Check if the address has at least 2 words
	if len(words) < 2 {
		return "invalid type"
	}

	addressFirstWord := words[1]
	isAddressValid := false

	for _, vType := range validTypes {
		if vType == addressFirstWord {
			isAddressValid = true
		}
	}

	if isAddressValid {
		return strings.Title(addressFirstWord)
	}

	return "invalid type"
}
