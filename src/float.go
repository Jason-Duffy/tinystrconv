package tinystrconv

// FloatToString converts a float to its string representation with the specified precision.
func FloatToString(f float64, precision int) (string, error) {
	if f != f {
		return "NaN", nil
	}
	if f > 1.7976931348623157e+308 {
		return "+Inf", nil
	}
	if f < -1.7976931348623157e+308 {
		return "-Inf", nil
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
