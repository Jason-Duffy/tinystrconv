package tinystrconv

import (
	"errors"
)

// FloatToString converts a float to its string representation with the specified precision.
func FloatToString(f float64, precision int) (string, error) {
	if f != f { // NaN
		return "", errors.New("cannot convert NaN to string")
	}

	neg := f < 0
	if neg {
		f = -f
	}

	intPart := int64(f)
	fracPart := f - float64(intPart)
	result, err := IntToString(int(intPart), 10)
	if err != nil {
		return "", err
	}

	if precision > 0 {
		result += "."
		for i := 0; i < precision; i++ {
			fracPart *= 10
			digit := int64(fracPart)
			result += string('0' + rune(digit))
			fracPart -= float64(digit)
		}
		// Perform rounding if necessary
		fracPart *= 10
		if int64(fracPart) >= 5 {
			result = roundUp(result)
		}
	} else {
		result += ".0"
	}

	if neg {
		result = "-" + result
	}

	return result, nil
}

// StringToFloat converts a string representation of a float to a float64.
func StringToFloat(s string) (float64, error) {
	if len(s) == 0 {
		return 0, errors.New("input string is empty")
	}

	neg := false
	if s[0] == '-' {
		neg = true
		s = s[1:]
	}

	intPartStr := ""
	fracPartStr := ""
	decimalPointSeen := false
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			if decimalPointSeen {
				return 0, errors.New("invalid float string")
			}
			decimalPointSeen = true
		} else if decimalPointSeen {
			fracPartStr += string(s[i])
		} else {
			intPartStr += string(s[i])
		}
	}

	intPart, err := StringToInt(intPartStr, 10)
	if err != nil {
		return 0, err
	}

	var fracPart float64
	fracDivisor := 1.0
	for i := 0; i < len(fracPartStr); i++ {
		fracPart = fracPart*10 + float64(fracPartStr[i]-'0')
		fracDivisor *= 10
	}
	fracPart /= fracDivisor

	result := float64(intPart) + fracPart
	if neg {
		result = -result
	}

	return result, nil
}

// roundUp handles rounding up the last digit if necessary.
func roundUp(s string) string {
	carry := true
	result := []byte(s)
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
