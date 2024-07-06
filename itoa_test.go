package tinystrconv

import "testing"

// TestItoa tests the Itoa function for converting integers to strings in base 10.
func TestItoa(t *testing.T) {
	testCases := []struct {
		input int
		want  string
	}{
		{0, "0"},       // Test case for zero in decimal
		{1, "1"},       // Test case for small positive integer in decimal
		{9, "9"},       // Test case for single digit positive integer in decimal
		{10, "10"},     // Test case for two-digit positive integer in decimal
		{123, "123"},   // Test case for a positive integer in decimal
		{-123, "-123"}, // Test case for a negative integer in decimal
	}

	for _, testCase := range testCases {
		got := Itoa(testCase.input)
		if got != testCase.want {
			t.Errorf("Itoa(%d) = %q, want %q", testCase.input, got, testCase.want)
		}
	}
}

// TestItoaBase tests the ItoaBase function for converting integers to strings in various bases.
func TestItoaBase(t *testing.T) {
	testCases := []struct {
		input int
		base  int
		want  string
		err   bool
	}{
		{0, 10, "0", false},         // Test case for zero in decimal
		{1, 10, "1", false},         // Test case for small positive integer in decimal
		{9, 10, "9", false},         // Test case for single digit positive integer in decimal
		{10, 10, "10", false},       // Test case for two-digit positive integer in decimal
		{123, 10, "123", false},     // Test case for a positive integer in decimal
		{-123, 10, "-123", false},   // Test case for a negative integer in decimal
		{255, 16, "ff", false},      // Test case for positive integer in hexadecimal
		{255, 2, "11111111", false}, // Test case for positive integer in binary
		{64, 8, "100", false},       // Test case for positive integer in octal
		{8, 3, "22", false},         // Test case for positive integer in ternary
		{1295, 36, "zz", false},     // Test case for positive integer in base 36
		{255, 37, "", true},         // Test case for unsupported base
		{255, 1, "", true},          // Test case for unsupported base
	}

	for _, testCase := range testCases {
		got, err := ItoaBase(testCase.input, testCase.base)
		if (err != nil) != testCase.err {
			t.Errorf("ItoaBase(%d, %d) error = %v, wantErr %v", testCase.input, testCase.base, err, testCase.err)
			continue
		}
		if got != testCase.want {
			t.Errorf("ItoaBase(%d, %d) = %q, want %q", testCase.input, testCase.base, got, testCase.want)
		}
	}
}
