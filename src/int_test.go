package tinystrconv

import "testing"

// TestIntToString tests the IntToString function for converting integers to strings in various bases.
func TestIntToString(t *testing.T) {
	testCases := []struct {
		input     int
		base      int
		expected  string
		expectErr bool
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

	for _, tc := range testCases {
		got, err := IntToString(tc.input, tc.base)
		if (err != nil) != tc.expectErr {
			t.Errorf("IntToString(%d, %d) error = %v, expectErr %v", tc.input, tc.base, err, tc.expectErr)
			continue
		}
		if got != tc.expected {
			t.Errorf("IntToString(%d, %d) = %q, want %q", tc.input, tc.base, got, tc.expected)
		}
	}
}
