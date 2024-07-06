package tinystrconv

import (
	"math"
)

// FloatToString converts a float to its string representation with the specified precision.
func FloatToString(floatValue float64, precision int) string {
	if math.IsNaN(floatValue) {
		return "NaN"
	}
	if math.IsInf(floatValue, 1) {
		return "+Inf"
	}
	if math.IsInf(floatValue, -1) {
		return "-Inf"
	}

	isNegative := floatValue < 0
	if isNegative {
		floatValue = -floatValue
	}

	integerPart := int64(floatValue)
	fractionalPart := floatValue - float64(integerPart)
	result, _ := IntToString(int(integerPart), 10)

	if precision > 0 {
		result += "."
		for i := 0; i < precision; i++ {
			fractionalPart *= 10
			digit := int64(fractionalPart)
			result += string('0' + rune(digit))
			fractionalPart -= float64(digit)
		}
		// Perform rounding if necessary
		fractionalPart *= 10
		if int64(fractionalPart) >= 5 {
			result = roundUp(result)
		}
	} else {
		result += ".0"
	}

	if isNegative {
		result = "-" + result
	}

	return result
}

// roundUp handles rounding up the last digit if necessary.
func roundUp(numberString string) string {
	carry := true
	result := []byte(numberString)
	for i := len(result) - 1; i >= 0 && carry; i-- {
		if result[i] == '.' {
			continue
		}
		if result[i] == '9' {
			result[i] = '0'
		} else {
			result[i]++
			carry = false
		}
	}
	if carry {
		result = append([]byte{'1'}, result...)
	}
	return string(result)
}
