package tinystrconv

import "testing"

// TestFormat tests the Format function for formatting strings with various format specifiers.
func TestFormat(t *testing.T) {
	testCases := []struct {
		format string
		args   []interface{}
		want   string
		err    bool
	}{
		{"Hello, %s!", []interface{}{"world"}, "Hello, world!", false},                     // Test case for a string
		{"Number: %d", []interface{}{42}, "Number: 42", false},                             // Test case for an integer
		{"Float: %f", []interface{}{3.14}, "Float: 3.14", false},                           // Test case for a float
		{"Boolean: %t", []interface{}{true}, "Boolean: true", false},                       // Test case for a boolean
		{"%s %d %f %t", []interface{}{"str", 10, 3.14, false}, "str 10 3.14 false", false}, // Test case for multiple arguments
		{"%s", []interface{}{123}, "", true},                                               // Error case: wrong type for %s
		{"%d", []interface{}{"abc"}, "", true},                                             // Error case: wrong type for %d
		{"%f", []interface{}{"3.14"}, "", true},                                            // Error case: wrong type for %f
		{"%t", []interface{}{"true"}, "", true},                                            // Error case: wrong type for %t
		{"%d", []interface{}{}, "", true},                                                  // Error case: missing argument for %d
		{"%s", []interface{}{"hello", "world"}, "", true},                                  // Error case: too many arguments
	}

	for _, testCase := range testCases {
		got, err := Format(testCase.format, testCase.args...)
		if (err != nil) != testCase.err {
			t.Errorf("Format(%q, %v) error = %v, wantErr %v", testCase.format, testCase.args, err, testCase.err)
			continue
		}
		if got != testCase.want {
			t.Errorf("Format(%q, %v) = %q, want %q", testCase.format, testCase.args, got, testCase.want)
		}
	}
}
