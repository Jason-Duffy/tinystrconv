package tinystrconv

// FormatBool converts a boolean to its string representation.
func FormatBool(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
