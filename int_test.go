// =============================================================================
// Project: tinystrconv
// File: int_test.go
// Description: Test suite for integer conversion functions.
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
	"testing"
)

// -------------------------------------------------------------------------- //
//                              Public Functions                              //
// -------------------------------------------------------------------------- //

// TestIntToString tests the IntToString function for converting integers to strings in various bases.
func TestIntToString(t *testing.T) {
	testCases := []struct {
		input int
		base  int
		want  string
		err   bool
	}{
		{0, 10, "0", false},           // Test case for zero in decimal
		{1, 10, "1", false},           // Test case for small positive integer in decimal
		{9, 10, "9", false},           // Test case for single digit positive integer in decimal
		{10, 10, "10", false},         // Test case for two-digit positive integer in decimal
		{123, 10, "123", false},       // Test case for a positive integer in decimal
		{-123, 10, "-123", false},     // Test case for a negative integer in decimal
		{255, 16, "0xff", false},      // Test case for positive integer in hexadecimal
		{255, 2, "0b11111111", false}, // Test case for positive integer in binary
		{64, 8, "0o100", false},       // Test case for positive integer in octal
		{255, 37, "", true},           // Test case for unsupported base
		{255, 1, "", true},            // Test case for unsupported base
	}

	for _, testCase := range testCases {
		got, err := IntToString(testCase.input, testCase.base)
		if (err != nil) != testCase.err {
			t.Errorf("IntToString(%d, %d) error = %v, wantErr %v", testCase.input, testCase.base, err, testCase.err)
			continue
		}
		if got != testCase.want {
			t.Errorf("IntToString(%d, %d) = %q, want %q", testCase.input, testCase.base, got, testCase.want)
		}
	}
}

// TestStringToInt tests the StringToInt function for converting strings to integers in various bases.
func TestStringToInt(t *testing.T) {
	testCases := []struct {
		input string
		base  int
		want  int
		err   bool
	}{
		{"0", 10, 0, false},           // Test case for zero in decimal
		{"1", 10, 1, false},           // Test case for small positive integer in decimal
		{"9", 10, 9, false},           // Test case for single digit positive integer in decimal
		{"10", 10, 10, false},         // Test case for two-digit positive integer in decimal
		{"123", 10, 123, false},       // Test case for a positive integer in decimal
		{"-123", 10, -123, false},     // Test case for a negative integer in decimal
		{"0xff", 16, 255, false},      // Test case for positive integer in hexadecimal
		{"0b11111111", 2, 255, false}, // Test case for positive integer in binary
		{"0o100", 8, 64, false},       // Test case for positive integer in octal
		{"255", 37, 0, true},          // Test case for unsupported base
		{"255", 1, 0, true},           // Test case for unsupported base
		{"1g", 16, 0, true},           // Test case for invalid character
	}

	for _, testCase := range testCases {
		got, err := StringToInt(testCase.input, testCase.base)
		if (err != nil) != testCase.err {
			t.Errorf("StringToInt(%s, %d) error = %v, wantErr %v", testCase.input, testCase.base, err, testCase.err)
			continue
		}
		if got != testCase.want {
			t.Errorf("StringToInt(%s, %d) = %d, want %d", testCase.input, testCase.base, got, testCase.want)
		}
	}
}
