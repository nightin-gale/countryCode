package countryCode

import (
	"fmt"
	"strings"

	"github.com/nightin-gale/countryCode/data"
)

const version = "1.0.2"

// GetCountryCode returns the country code for the given country name.
// Returns an error if the country name is not found.
func GetCountryCode(name string) (string, error) {
	key := strings.ToUpper(name)
	res, ok := data.NameToCode[key]
	if ok {
		return res, nil
	} else {
		return "", fmt.Errorf("Country Code for %s not found", key)
	}
}

// GetCountryName returns the country name for the given country code.
// Returns an error if the country code is not found.
func GetCountryName(code string) (string, error) {
	if len(code) != 2 {
		return "", fmt.Errorf("Country code should be 2 characters")
	}

	key := strings.ToUpper(code)
	res, ok := data.CodeToName[key]
	if ok {
		return res, nil
	} else {
		return "UNKNOWN", nil
	}
}
