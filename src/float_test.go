package tinystrconv

import (
	"math"
	"testing"
)

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
		{math.Inf(1), 2, "+Inf", false},      // Test case for positive infinity
		{math.Inf(-1), 2, "-Inf", false},     // Test case for negative infinity
		{math.NaN(), 2, "NaN", false},        // Test case for NaN
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
