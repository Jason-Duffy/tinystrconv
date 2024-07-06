package tinystrconv

import "testing"

// TestFormatFloat tests the FormatFloat function for converting floats to strings.
func TestFormatFloat(t *testing.T) {
	testCases := []struct {
		input     float64
		precision int
		want      string
	}{
		{0.0, 2, "0.00"},              // Test case for zero
		{3.14, 2, "3.14"},             // Test case for a positive float
		{-2.718, 3, "-2.718"},         // Test case for a negative float
		{1.23456789, 8, "1.23456789"}, // Test case for a float with more decimal places
		{1234567.89, 2, "1234567.89"}, // Test case for a larger float
	}

	for _, testCase := range testCases {
		got := FormatFloat(testCase.input, testCase.precision)
		if got != testCase.want {
			t.Errorf("FormatFloat(%f, %d) = %q, want %q", testCase.input, testCase.precision, got, testCase.want)
		}
	}
}
