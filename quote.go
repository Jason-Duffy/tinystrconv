package tinystrconv

// QuoteString converts a string to its quoted representation.
func QuoteString(input string) string {
	var result []byte
	result = append(result, '"')
	for i := 0; i < len(input); i++ {
		char := input[i]
		switch char {
		case '\a':
			result = append(result, '\\', 'a')
		case '\b':
			result = append(result, '\\', 'b')
		case '\f':
			result = append(result, '\\', 'f')
		case '\n':
			result = append(result, '\\', 'n')
		case '\r':
			result = append(result, '\\', 'r')
		case '\t':
			result = append(result, '\\', 't')
		case '\v':
			result = append(result, '\\', 'v')
		case '\\':
			result = append(result, '\\', '\\')
		case '"':
			result = append(result, '\\', '"')
		default:
			result = append(result, char)
		}
	}
	result = append(result, '"')
	return string(result)
}
