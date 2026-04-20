package v2

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ParseInt[T ~int | ~int8 | ~int16 | ~int32 | ~int64](s string) (T, error) {
	var zero T
	var val int64
	var err error
	if zeroType := any(zero); zeroType != nil {
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
	return zero, fmt.Errorf("unsupported integer type: %T", zero)
}

func ParseUint[T ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64](s string) (T, error) {
	var zero T
	var val uint64
	var err error
	if zeroType := any(zero); zeroType != nil {
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
	return zero, fmt.Errorf("unsupported unsigned integer type: %T", zero)
}

func ParseString(s string) (string, error) {
	return s, nil
}

func ParseBool(s string) (bool, error) {
	return strconv.ParseBool(s)
}

func ParseFloat[T ~float32 | ~float64](s string) (T, error) {
	var zero T
	var val float64
	var err error
	if zeroType := any(zero); zeroType != nil {
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
	return zero, fmt.Errorf("unsupported float type: %T", zero)
}

func ParseDuration(s string) (time.Duration, error) {
	return time.ParseDuration(s)
}

func ParseStringSlice(s string) ([]string, error) {
	return strings.Split(s, ","), nil
}
