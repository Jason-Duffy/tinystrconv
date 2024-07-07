package tinystrconv

import "testing"

// TestQuoteString tests the QuoteString function for converting strings to their quoted representations.
func TestQuoteString(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{
		{"Hello, world!", `"Hello, world!"`}, // Test case for a simple string
		{"", `""`},                           // Test case for an empty string
		{"\n", `"\n"`},                       // Test case for a newline character
		{"\t", `"\t"`},                       // Test case for a tab character
		{"\"\"", `"\"\""`},                   // Test case for double quotes
		{"\\", `"\\"`},                       // Test case for a backslash
		{"Hello\nWorld", `"Hello\nWorld"`},   // Test case for a string with newline
		{"Hello\tWorld", `"Hello\tWorld"`},   // Test case for a string with tab
		{"Hello\\World", `"Hello\\World"`},   // Test case for a string with backslash
		{"Hello\"World", `"Hello\"World"`},   // Test case for a string with double quote
		{"\a\b\f\r\v", `"\a\b\f\r\v"`},       // Test case for various escape sequences
	}

	for _, testCase := range testCases {
		got := QuoteString(testCase.input)
		if got != testCase.want {
			t.Errorf("QuoteString(%q) = %q, want %q", testCase.input, got, testCase.want)
		}
	}
}
