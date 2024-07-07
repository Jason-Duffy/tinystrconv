// =============================================================================
// Project: tinystrconv
// File: case.go
// Description: Functions for case conversion.
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
//                          Private Helper Functions                          //
// -------------------------------------------------------------------------- //

// toLowerCase converts a string to lowercase.
func toLowerCase(input string) string {
	result := make([]byte, len(input))
	for index := 0; index < len(input); index++ {
		character := input[index]
		if 'A' <= character && character <= 'Z' {
			character += 'a' - 'A'
		}
		result[index] = character
	}
	return string(result)
}

// toUpperCase converts a string to uppercase.
func toUpperCase(input string) string {
	result := make([]byte, len(input))
	for index := 0; index < len(input); index++ {
		character := input[index]
		if 'a' <= character && character <= 'z' {
			character -= 'a' - 'A'
		}
		result[index] = character
	}
	return string(result)
}
