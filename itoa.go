package tinystrconv

import (
	"errors"
)

const digits = "0123456789abcdefghijklmnopqrstuvwxyz"

// Itoa converts an integer to its string representation in base 10.
func Itoa(n int) string {
	result, _ := ItoaBase(n, 10)
	return result
}

// ItoaBase converts an integer to its string representation in the specified base.
func ItoaBase(n int, base int) (string, error) {
	if base < 2 || base > len(digits) {
		return "", errors.New("unsupported base")
	}

	if n == 0 {
		return "0", nil
	}

	var result []byte
	negative := false
	if base == 10 && n < 0 {
		negative = true
		n = -n
	}

	for n > 0 {
		digit := n % base
		result = append([]byte{digits[digit]}, result...)
		n /= base
	}

	if negative {
		result = append([]byte{'-'}, result...)
	}

	return string(result), nil
}
