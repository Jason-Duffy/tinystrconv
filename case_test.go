// =============================================================================
// Project: tinystrconv
// File: case_test.go
// Description: Test suite for case conversion functions.
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

// TestToLowerCase tests the toLowerCase function for converting strings to lowercase.
func TestToLowerCase(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"", ""},                         // Test case for empty string
		{"a", "a"},                       // Test case for single lowercase letter
		{"A", "a"},                       // Test case for single uppercase letter
		{"Hello", "hello"},               // Test case for mixed case string
		{"WORLD", "world"},               // Test case for all uppercase string
		{"GoLang", "golang"},             // Test case for mixed case string
		{"12345", "12345"},               // Test case for numbers
		{"!@#$%", "!@#$%"},               // Test case for special characters
		{"HeLLo WoRLD!", "hello world!"}, // Test case for mixed characters
	}

	for _, tc := range testCases {
		got := toLowerCase(tc.input)
		if got != tc.expected {
			t.Errorf("toLowerCase(%q) = %q, want %q", tc.input, got, tc.expected)
		}
	}
}

// TestToUpperCase tests the toUpperCase function for converting strings to uppercase.
func TestToUpperCase(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"", ""},                         // Test case for empty string
		{"a", "A"},                       // Test case for single lowercase letter
		{"A", "A"},                       // Test case for single uppercase letter
		{"Hello", "HELLO"},               // Test case for mixed case string
		{"WORLD", "WORLD"},               // Test case for all uppercase string
		{"GoLang", "GOLANG"},             // Test case for mixed case string
		{"12345", "12345"},               // Test case for numbers
		{"!@#$%", "!@#$%"},               // Test case for special characters
		{"HeLLo WoRLD!", "HELLO WORLD!"}, // Test case for mixed characters
	}

	for _, tc := range testCases {
		got := toUpperCase(tc.input)
		if got != tc.expected {
			t.Errorf("toUpperCase(%q) = %q, want %q", tc.input, got, tc.expected)
		}
	}
}
