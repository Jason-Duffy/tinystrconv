// =============================================================================
// Project: tinystrconv
// File: int.go
// Description: Functions for integer conversion.
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

// IntToString converts an integer to its string representation in the specified base.
func IntToString(number int, base int) (string, error) {
	isNegative := number < 0
	if isNegative {
		number = -number
	}
	return toStringHelper(uint64(number), base, isNegative)
}

// StringToInt converts a string representation of an integer in the specified base to an int.
func StringToInt(input string, base int) (int, error) {
	isNegative := false
	if input[0] == '-' {
		if base != 10 {
			return 0, errors.New("negative numbers are not supported for non-decimal bases")
		}
		isNegative = true
		input = input[1:]
	}

	number, err := stringToNumberHelper(input, base)
	if err != nil {
		return 0, err
	}

	if isNegative {
		return -int(number), nil
	}
	return int(number), nil
}
