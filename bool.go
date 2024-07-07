package tinystrconv

import (
	"errors"
)

// BoolToString converts a boolean to its string representation.
func BoolToString(boolValue bool) string {
	if boolValue {
		return "true"
	}
	return "false"
}

// StringToBool converts a string (e.g., "true", or "1") to its boolean representation.
func StringToBool(input string) (bool, error) {
	lowercaseInput := toLowerCase(input)
	switch lowercaseInput {
	case "true", "1":
		return true, nil
	case "false", "0":
		return false, nil
	default:
		return false, errors.New("invalid boolean string")
	}
}
