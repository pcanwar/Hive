package middleware

import (
	"regexp"
	"strings"
)

// https://goethereumbook.org/en/address-check/

func isAddressValid(address string) bool {
	if !strings.HasPrefix(address, "0x") {
		return false
	}
	match, err := regexp.MatchString("^0x[0-9a-fA-F]{40}$", address)
	if err != nil || !match {
		return false
	}
	// if reAddress.MatchString(address) {
	// 	return false
	// }
	return true
}
