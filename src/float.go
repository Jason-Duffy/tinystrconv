package tinystrconv

import (
	"errors"
)

// FloatToString converts a float to its string representation with the specified precision.
func FloatToString(value float64, precision int) (string, error) {
	if value != value { // NaN
		return "", errors.New("cannot convert NaN to string")
	}

	isNegative := value < 0
	if isNegative {
		value = -value
	}

	integerPart := int64(value)
	fractionPart := value - float64(integerPart)
	result, err := IntToString(int(integerPart), 10)
	if err != nil {
		return "", err
	}

	if precision > 0 {
		result += "."
		for i := 0; i < precision; i++ {
			fractionPart *= 10
			digit := int64(fractionPart)
			result += string('0' + rune(digit))
			fractionPart -= float64(digit)
		}
		// Perform rounding if necessary
		fractionPart *= 10
		if int64(fractionPart) >= 5 {
			result = roundUp(result)
		}
	} else {
		result += ".0"
	}

	if isNegative {
		result = "-" + result
	}

	return result, nil
}

// StringToFloat converts a string representation of a float to a float64.
func StringToFloat(input string) (float64, error) {
	if len(input) == 0 {
		return 0, errors.New("input string is empty")
	}

	isNegative := false
	if input[0] == '-' {
		isNegative = true
		input = input[1:]
	}

	integerPartStr := ""
	fractionPartStr := ""
	decimalPointSeen := false
	for i := 0; i < len(input); i++ {
		if input[i] == '.' {
			if decimalPointSeen {
				return 0, errors.New("invalid float string")
			}
			decimalPointSeen = true
		} else if decimalPointSeen {
			fractionPartStr += string(input[i])
		} else {
			integerPartStr += string(input[i])
		}
	}

	integerPart, err := StringToInt(integerPartStr, 10)
	if err != nil {
		return 0, err
	}

	var fractionPart float64
	fractionDivisor := 1.0
	for i := 0; i < len(fractionPartStr); i++ {
		fractionPart = fractionPart*10 + float64(fractionPartStr[i]-'0')
		fractionDivisor *= 10
	}
	fractionPart /= fractionDivisor

	result := float64(integerPart) + fractionPart
	if isNegative {
		result = -result
	}

	return result, nil
}

// roundUp handles rounding up the last digit if necessary.
func roundUp(input string) string {
	carry := true
	result := []byte(input)
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
