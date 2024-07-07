package tinystrconv

import "testing"

func TestFormat(t *testing.T) {
	testCases := []struct {
		format    string
		args      []interface{}
		want      string
		shouldErr bool
	}{
		{"Hello, %s!", []interface{}{"world"}, "Hello, world!", false},
		{"Value: %d", []interface{}{42}, "Value: 42", false},
		{"Hex: %x", []interface{}{255}, "Hex: ff", false},
		{"Binary: %b", []interface{}{7}, "Binary: 111", false},
		{"Float: %.2f", []interface{}{3.14159}, "Float: 3.14", false},
		{"Float: %.5f", []interface{}{3.14159}, "Float: 3.14159", false},
		{"Invalid: %q", []interface{}{42}, "", true},
		{"Missing arg: %d %d", []interface{}{42}, "", true},
	}

	for _, tc := range testCases {
		got, err := Format(tc.format, tc.args...)
		if (err != nil) != tc.shouldErr {
			t.Errorf("Format(%q, %v) error = %v, wantErr %v", tc.format, tc.args, err, tc.shouldErr)
			continue
		}
		if got != tc.want {
			t.Errorf("Format(%q, %v) = %q, want %q", tc.format, tc.args, got, tc.want)
		}
	}
}
