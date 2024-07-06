package tinystrconv

import "testing"

// TestFormatBool tests the FormatBool function for converting booleans to strings.
func TestFormatBool(t *testing.T) {
	testCases := []struct {
		input bool
		want  string
	}{
		{true, "true"},   // Test case for true
		{false, "false"}, // Test case for false
	}

	for _, testCase := range testCases {
		got := FormatBool(testCase.input)
		if got != testCase.want {
			t.Errorf("FormatBool(%t) = %q, want %q", testCase.input, got, testCase.want)
		}
	}
}
