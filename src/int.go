package tinystrconv

import (
	"errors"
)

const digitChars = "0123456789abcdefghijklmnopqrstuvwxyz"

// IntToString converts an integer to its string representation in the specified base.
func IntToString(number int, base int) (string, error) {
	if base < 2 || base > len(digitChars) {
		return "", errors.New("unsupported base")
	}

	if number == 0 {
		return "0", nil
	}

	var result []byte
	isNegative := false
	if number < 0 {
		if base != 10 {
			return "", errors.New("negative numbers are not supported for non-decimal bases")
		}
		isNegative = true
		number = -number
	}

	for number > 0 {
		digit := number % base
		result = append(result, digitChars[digit])
		number /= base
	}

	if isNegative {
		result = append(result, '-')
	}

	// Reverse the result slice
	for left, right := 0, len(result)-1; left < right; left, right = left+1, right-1 {
		result[left], result[right] = result[right], result[left]
	}

	return string(result), nil
}

// StringToInt converts a string representation of an integer in the specified base to an int.
func StringToInt(s string, base int) (int, error) {
	if base < 2 || base > len(digitChars) {
		return 0, errors.New("unsupported base")
	}

	isNegative := false
	if s[0] == '-' {
		if base != 10 {
			return 0, errors.New("negative numbers are not supported for non-decimal bases")
		}
		isNegative = true
		s = s[1:]
	}

	var result int
	for _, c := range s {
		digit := -1
		for i := 0; i < base; i++ {
			if c == rune(digitChars[i]) {
				digit = i
				break
			}
		}
		if digit == -1 {
			return 0, errors.New("invalid character in input string")
		}
		result = result*base + digit
	}

	if isNegative {
		result = -result
	}

	return result, nil
}
