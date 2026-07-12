// Package parsers provides the parser functions used by go-cli to convert
// user input strings into typed values. They can be used with cli.FlagFunc and
// cli.PosFunc to build custom flags and positional arguments, and composed
// with the Slice combinator for comma-separated lists.
package parsers

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// String returns the input string as-is.
func String(s string) (string, error) {
	return s, nil
}

// Bool parses a boolean value, accepting the same forms as strconv.ParseBool.
func Bool(s string) (bool, error) {
	return strconv.ParseBool(s)
}

// Int parses a signed integer of the given type, validating its bit size.
func Int[T ~int | ~int8 | ~int16 | ~int32 | ~int64](s string) (T, error) {
	var zero T
	var val int64
	var err error
	switch any(zero).(type) {
	case int:
		val, err = strconv.ParseInt(s, 10, 0)
	case int8:
		val, err = strconv.ParseInt(s, 10, 8)
	case int16:
		val, err = strconv.ParseInt(s, 10, 16)
	case int32:
		val, err = strconv.ParseInt(s, 10, 32)
	case int64:
		val, err = strconv.ParseInt(s, 10, 64)
	default:
		return zero, fmt.Errorf("unsupported integer type: %T", zero)
	}
	if err != nil {
		return zero, err
	}
	return T(val), nil
}

// Uint parses an unsigned integer of the given type, validating its bit size.
func Uint[T ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64](s string) (T, error) {
	var zero T
	var val uint64
	var err error
	switch any(zero).(type) {
	case uint:
		val, err = strconv.ParseUint(s, 10, 0)
	case uint8:
		val, err = strconv.ParseUint(s, 10, 8)
	case uint16:
		val, err = strconv.ParseUint(s, 10, 16)
	case uint32:
		val, err = strconv.ParseUint(s, 10, 32)
	case uint64:
		val, err = strconv.ParseUint(s, 10, 64)
	default:
		return zero, fmt.Errorf("unsupported unsigned integer type: %T", zero)
	}
	if err != nil {
		return zero, err
	}
	return T(val), nil
}

// Float parses a floating point number of the given type, validating its bit
// size.
func Float[T ~float32 | ~float64](s string) (T, error) {
	var zero T
	var val float64
	var err error
	switch any(zero).(type) {
	case float32:
		val, err = strconv.ParseFloat(s, 32)
	case float64:
		val, err = strconv.ParseFloat(s, 64)
	default:
		return zero, fmt.Errorf("unsupported float type: %T", zero)
	}
	if err != nil {
		return zero, err
	}
	return T(val), nil
}

// Duration parses a duration using time.ParseDuration (e.g. "1h30m").
func Duration(s string) (time.Duration, error) {
	return time.ParseDuration(s)
}

// Slice turns a parser for T into a parser for []T, splitting the input on
// commas and parsing each element. Example: Slice(Int[int]) parses "1,2,3"
// into []int{1, 2, 3}.
func Slice[T any](parser func(string) (T, error)) func(string) ([]T, error) {
	return func(s string) ([]T, error) {
		parts := strings.Split(s, ",")
		result := make([]T, len(parts))
		for i, part := range parts {
			val, err := parser(part)
			if err != nil {
				return nil, err
			}
			result[i] = val
		}
		return result, nil
	}
}
