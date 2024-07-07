package tinystrconv

import (
	"errors"
)

// Format formats a string with the provided arguments, similar to fmt.Sprintf.
func Format(format string, args ...interface{}) (string, error) {
	var result []byte
	argIndex := 0

	for i := 0; i < len(format); i++ {
		if format[i] == '%' {
			if i+1 < len(format) {
				switch format[i+1] {
				case 'd':
					if argIndex >= len(args) {
						return "", errors.New("missing argument for %d")
					}
					intVal, ok := args[argIndex].(int)
					if !ok {
						return "", errors.New("argument for %d is not an int")
					}
					str, err := IntToString(intVal, 10)
					if err != nil {
						return "", err
					}
					result = append(result, []byte(str)...)
					argIndex++
				case 'f':
					if argIndex >= len(args) {
						return "", errors.New("missing argument for %f")
					}
					floatVal, ok := args[argIndex].(float64)
					if !ok {
						return "", errors.New("argument for %f is not a float64")
					}
					formattedFloat, err := FloatToString(floatVal, 2)
					if err != nil {
						return "", err
					}
					result = append(result, []byte(formattedFloat)...)
					argIndex++
				case 's':
					if argIndex >= len(args) {
						return "", errors.New("missing argument for %s")
					}
					strVal, ok := args[argIndex].(string)
					if !ok {
						return "", errors.New("argument for %s is not a string")
					}
					result = append(result, []byte(strVal)...)
					argIndex++
				case 't':
					if argIndex >= len(args) {
						return "", errors.New("missing argument for %t")
					}
					boolVal, ok := args[argIndex].(bool)
					if !ok {
						return "", errors.New("argument for %t is not a bool")
					}
					result = append(result, []byte(BoolToString(boolVal))...)
					argIndex++
				case 'x':
					if argIndex >= len(args) {
						return "", errors.New("missing argument for %x")
					}
					intVal, ok := args[argIndex].(int)
					if !ok {
						return "", errors.New("argument for %x is not an int")
					}
					str, err := IntToString(intVal, 16)
					if err != nil {
						return "", err
					}
					result = append(result, []byte(str)...)
					argIndex++
				case 'b':
					if argIndex >= len(args) {
						return "", errors.New("missing argument for %b")
					}
					intVal, ok := args[argIndex].(int)
					if !ok {
						return "", errors.New("argument for %b is not an int")
					}
					str, err := IntToString(intVal, 2)
					if err != nil {
						return "", err
					}
					result = append(result, []byte(str)...)
					argIndex++
				case 'o':
					if argIndex >= len(args) {
						return "", errors.New("missing argument for %o")
					}
					intVal, ok := args[argIndex].(int)
					if !ok {
						return "", errors.New("argument for %o is not an int")
					}
					str, err := IntToString(intVal, 8)
					if err != nil {
						return "", err
					}
					result = append(result, []byte(str)...)
					argIndex++
				default:
					return "", errors.New("unknown format specifier")
				}
				i++
			} else {
				return "", errors.New("incomplete format specifier at end of format string")
			}
		} else {
			result = append(result, format[i])
		}
	}

	if argIndex < len(args) {
		return "", errors.New("too many arguments provided")
	}

	return string(result), nil
}
