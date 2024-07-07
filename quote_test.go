// =============================================================================
// Project: tinystrconv
// File: quote_test.go
// Description: Test suite for quote conversion functions.
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

import "testing"

// -------------------------------------------------------------------------- //
//                              Public Functions                              //
// -------------------------------------------------------------------------- //

// TestQuoteString tests the QuoteString function for converting strings to their quoted representations.
func TestQuoteString(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{
		{"Hello, world!", `"Hello, world!"`}, // Test case for a simple string
		{"", `""`},                           // Test case for an empty string
		{"\n", `"\n"`},                       // Test case for a newline character
		{"\t", `"\t"`},                       // Test case for a tab character
		{"\"\"", `"\"\""`},                   // Test case for double quotes
		{"\\", `"\\"`},                       // Test case for a backslash
		{"Hello\nWorld", `"Hello\nWorld"`},   // Test case for a string with newline
		{"Hello\tWorld", `"Hello\tWorld"`},   // Test case for a string with tab
		{"Hello\\World", `"Hello\\World"`},   // Test case for a string with backslash
		{"Hello\"World", `"Hello\"World"`},   // Test case for a string with double quote
		{"\a\b\f\r\v", `"\a\b\f\r\v"`},       // Test case for various escape sequences
	}

	for _, testCase := range testCases {
		got := QuoteString(testCase.input)
		if got != testCase.want {
			t.Errorf("QuoteString(%q) = %q, want %q", testCase.input, got, testCase.want)
		}
	}
}
