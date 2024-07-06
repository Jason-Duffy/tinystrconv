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
