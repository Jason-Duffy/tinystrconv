// =============================================================================
// Project: tinystrconv
// File: helpers.go
// Description: Helper functions for string conversion.
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
//               Private Consts, Structs & Variable Definitions               //
// -------------------------------------------------------------------------- //

const digitCharacters = "0123456789abcdefghijklmnopqrstuvwxyz"

// -------------------------------------------------------------------------- //
//                          Private Helper Functions                          //
// -------------------------------------------------------------------------- //

// toStringHelper is a helper function that converts an unsigned integer to its
// string representation in the specified base.
func toStringHelper(number uint64, base int, isNegative bool) (string, error) {
	if base < 2 || base > len(digitCharacters) {
		return "", errors.New("unsupported base")
	}

	if number == 0 {
		return "0", nil
	}

	var result []byte
	for number > 0 {
		digit := number % uint64(base)
		result = append(result, digitCharacters[digit])
		number /= uint64(base)
	}

	if isNegative {
		result = append(result, '-')
	}

	// Reverse the result slice
	for leftIndex, rightIndex := 0, len(result)-1; leftIndex < rightIndex; leftIndex, rightIndex = leftIndex+1, rightIndex-1 {
		result[leftIndex], result[rightIndex] = result[rightIndex], result[leftIndex]
	}

	// Add base prefix
	prefix := ""
	switch base {
	case 2:
		prefix = "0b"
	case 8:
		prefix = "0o"
	case 16:
		prefix = "0x"
	}

	return prefix + string(result), nil
}

// stringToNumberHelper is a helper function that converts a string
// representation of an unsigned integer in the specified base to a uint64.
func stringToNumberHelper(input string, base int) (uint64, error) {
	if base < 2 || base > len(digitCharacters) {
		return 0, errors.New("unsupported base")
	}

	// Remove base prefix
	switch base {
	case 2:
		if len(input) > 2 && input[:2] == "0b" {
			input = input[2:]
		}
	case 8:
		if len(input) > 2 && input[:2] == "0o" {
			input = input[2:]
		}
	case 16:
		if len(input) > 2 && input[:2] == "0x" {
			input = input[2:]
		}
	}

	var result uint64
	for _, character := range input {
		digit := -1
		for index := 0; index < base; index++ {
			if character == rune(digitCharacters[index]) {
				digit = index
				break
			}
		}
		if digit == -1 {
			return 0, errors.New("invalid character in input string")
		}
		result = result*uint64(base) + uint64(digit)
	}

	return result, nil
}
