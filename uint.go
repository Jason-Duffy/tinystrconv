// =============================================================================
// Project: tinystrconv
// File: uint.go
// Description: Functions for unsigned integer conversion.
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

// -------------------------------------------------------------------------- //
//                              Public Functions                              //
// -------------------------------------------------------------------------- //

// UintToString converts an unsigned integer to its string representation in the specified base.
func UintToString(number uint, base int) (string, error) {
	return toStringHelper(uint64(number), base, false)
}

// StringToUint converts a string representation of an unsigned integer in the specified base to a uint.
func StringToUint(input string, base int) (uint, error) {
	number, err := stringToNumberHelper(input, base)
	if err != nil {
		return 0, err
	}
	return uint(number), nil
}
