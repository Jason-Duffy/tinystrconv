package tinystrconv

import (
	"testing"
)

const floatTolerance = 1e-9 // Tolerance for floating-point comparisons

// TestFloatToString tests the FloatToString function for converting floats to strings.
func TestFloatToString(t *testing.T) {
	testCases := []struct {
		input     float64
		precision int
		want      string
		err       bool
	}{
		{0.0, 2, "0.00", false},              // Test case for zero
		{3.14, 2, "3.14", false},             // Test case for a positive float
		{-2.718, 3, "-2.718", false},         // Test case for a negative float
		{1.23456789, 8, "1.23456789", false}, // Test case for a float with more decimal places
		{1234567.89, 2, "1234567.89", false}, // Test case for a larger float
	}

	for _, testCase := range testCases {
		got, err := FloatToString(testCase.input, testCase.precision)
		if (err != nil) != testCase.err {
			t.Errorf("FloatToString(%f, %d) error = %v, wantErr %v", testCase.input, testCase.precision, err, testCase.err)
			continue
		}
		if got != testCase.want {
			t.Errorf("FloatToString(%f, %d) = %q, want %q", testCase.input, testCase.precision, got, testCase.want)
		}
	}
}

// TestStringToFloat tests the StringToFloat function for converting strings to floats.
func TestStringToFloat(t *testing.T) {
	testCases := []struct {
		input string
		want  float64
		err   bool
	}{
		{"0.0", 0.0, false},               // Test case for zero
		{"3.14", 3.14, false},             // Test case for a positive float
		{"-2.718", -2.718, false},         // Test case for a negative float
		{"1.23456789", 1.23456789, false}, // Test case for a float with more decimal places
		{"1234567.89", 1234567.89, false}, // Test case for a larger float
		{"invalid", 0, true},              // Test case for invalid input
		{"", 0, true},                     // Test case for empty string
		{"3.14.15", 0, true},              // Test case for multiple decimal points
	}

	for _, testCase := range testCases {
		got, err := StringToFloat(testCase.input)
		if (err != nil) != testCase.err {
			t.Errorf("StringToFloat(%s) error = %v, wantErr %v", testCase.input, err, testCase.err)
			continue
		}
		if testCase.err {
			continue
		}
		if !floatsEqual(got, testCase.want) {
			t.Errorf("StringToFloat(%s) = %v, want %v", testCase.input, got, testCase.want)
		}
	}
}

// floatsEqual compares two floating-point numbers for equality within a tolerance.
func floatsEqual(a, b float64) bool {
	if a == b {
		return true
	}
	diff := a - b
	return diff < floatTolerance && -diff < floatTolerance
}
