package tinystrconv

// toLowerCase converts a string to lowercase.
func toLowerCase(input string) string {
	result := make([]byte, len(input))
	for index := 0; index < len(input); index++ {
		character := input[index]
		if 'A' <= character && character <= 'Z' {
			character += 'a' - 'A'
		}
		result[index] = character
	}
	return string(result)
}

// toUpperCase converts a string to uppercase.
func toUpperCase(input string) string {
	result := make([]byte, len(input))
	for index := 0; index < len(input); index++ {
		character := input[index]
		if 'a' <= character && character <= 'z' {
			character -= 'a' - 'A'
		}
		result[index] = character
	}
	return string(result)
}
