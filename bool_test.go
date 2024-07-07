package tinystrconv

import "testing"

// TestBoolToString tests the BoolToString function for converting booleans to strings.
func TestBoolToString(t *testing.T) {
	testCases := []struct {
		input bool
		want  string
	}{
		{true, "true"},   // Test case for true
		{false, "false"}, // Test case for false
	}

	for _, testCase := range testCases {
		got := BoolToString(testCase.input)
		if got != testCase.want {
			t.Errorf("BoolToString(%t) = %q, want %q", testCase.input, got, testCase.want)
		}
	}
}

// TestStringToBool tests the StringToBool function for converting strings to booleans.
func TestStringToBool(t *testing.T) {
	testCases := []struct {
		input string
		want  bool
		err   bool
	}{
		{"true", true, false},          // Test case for true
		{"false", false, false},        // Test case for false
		{"TRUE", true, false},          // Test case for uppercase true
		{"FALSE", false, false},        // Test case for uppercase false
		{"True", true, false},          // Test case for mixed case true
		{"False", false, false},        // Test case for mixed case false
		{"1", true, false},             // Test case for numerical representation of true
		{"0", false, false},            // Test case for numerical representation of false
		{"yes", false, true},           // Test case for invalid string
		{"no", false, true},            // Test case for invalid string
		{"not_a_boolean", false, true}, // Test case for invalid string
	}

	for _, testCase := range testCases {
		got, err := StringToBool(testCase.input)
		if (err != nil) != testCase.err {
			t.Errorf("StringToBool(%q) error = %v, wantErr %v", testCase.input, err, testCase.err)
			continue
		}
		if got != testCase.want {
			t.Errorf("StringToBool(%q) = %t, want %t", testCase.input, got, testCase.want)
		}
	}
}
