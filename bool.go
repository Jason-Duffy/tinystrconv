// =============================================================================
// Project: tinystrconv
// File: bool.go
// Description: Conversion functions for boolean values.
// Datasheet/Docs:
//
// Author: Jason Duffy
// Created on: 06/07/2024
//
// Copyright: (C) 2024, Jason Duffy
// License: See LICENSE file in the project root for full license information.
// Disclaimer: See DISCLAIMER file in the project root for full disclaimer.
// =============================================================================

// -------------------------------------------------------------------------- //
//                               Import Statement                             //
// -------------------------------------------------------------------------- //

package tinystrconv

import (
	"errors"
)

// -------------------------------------------------------------------------- //
//                              Public Functions                              //
// -------------------------------------------------------------------------- //

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
