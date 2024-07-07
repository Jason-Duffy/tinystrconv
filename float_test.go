package tinystrconv

import (
	"math"
	"testing"
)

const epsilon = 0.0000001 // Tolerance for floating-point comparison

func TestFloatToString(t *testing.T) {
	testCases := []struct {
		input     float64
		precision int
		want      string
		shouldErr bool
	}{
		{0.0, 2, "0.00", false},              // Test formatting zero with precision 2
		{3.14, 2, "3.14", false},             // Test formatting positive float with precision 2
		{-2.718, 3, "-2.718", false},         // Test formatting negative float with precision 3
		{1.23456789, 8, "1.23456789", false}, // Test formatting float with precision 8
		{1234567.89, 2, "1234567.89", false}, // Test formatting large float with precision 2
		{math.NaN(), 2, "", true},            // Test formatting NaN
		{0.0, 0, "0", false},                 // Test formatting zero with precision 0
		{123.456, 0, "123", false},           // Test formatting float with precision 0
	}

	for _, testCase := range testCases {
		got, err := FloatToString(testCase.input, testCase.precision)
		if (err != nil) != testCase.shouldErr {
			t.Errorf("FloatToString(%f, %d) error = %v, wantErr %v", testCase.input, testCase.precision, err, testCase.shouldErr)
			continue
		}
		if got != testCase.want {
			t.Errorf("FloatToString(%f, %d) = %q, want %q", testCase.input, testCase.precision, got, testCase.want)
		}
	}
}

func TestStringToFloat(t *testing.T) {
	testCases := []struct {
		input     string
		want      float64
		shouldErr bool
	}{
		{"0.00", 0.0, false},              // Test parsing zero with precision 2
		{"3.14", 3.14, false},             // Test parsing positive float
		{"-2.718", -2.718, false},         // Test parsing negative float
		{"1.23456789", 1.23456789, false}, // Test parsing float with precision 8
		{"1234567.89", 1234567.89, false}, // Test parsing large float
		{"NaN", 0.0, true},                // Test parsing NaN
		{"", 0.0, true},                   // Test parsing empty string
		{"-123.456", -123.456, false},     // Test parsing negative float
		{"123", 123.0, false},             // Test parsing integer string
	}

	for _, testCase := range testCases {
		got, err := StringToFloat(testCase.input)
		if (err != nil) != testCase.shouldErr {
			t.Errorf("StringToFloat(%s) error = %v, wantErr %v", testCase.input, err, testCase.shouldErr)
			continue
		}
		if math.Abs(got-testCase.want) > epsilon {
			t.Errorf("StringToFloat(%s) = %f, want %f", testCase.input, got, testCase.want)
		}
	}
}
