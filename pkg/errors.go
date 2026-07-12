package pkg

import (
	"errors"
	"fmt"
	"strings"
)

var ErrFlagNotFound = errors.New("flag not found")
var ErrFlagNotType = errors.New("flag is not of the expected type")

// InvalidFlagValueError is returned when the value provided for a flag cannot
// be parsed or fails its validation function.
type InvalidFlagValueError struct {
	Flag   Flag   // the flag that received the invalid value
	Value  string // the raw value provided by the user
	Detail string // the detail shown in the message (the raw value or the validation error)
	Cause  error  // the underlying parser or validator error, if any
}

func (e *InvalidFlagValueError) Error() string {
	return fmt.Sprintf(GetLocale().ErrInvalidFlagValue, e.Flag.Signature(), e.Detail)
}
func (e *InvalidFlagValueError) Unwrap() error { return e.Cause }

// InvalidPositionalValueError is returned when the value provided for a
// positional argument cannot be parsed or fails its validation function.
type InvalidPositionalValueError struct {
	Positional Positional // the positional that received the invalid value
	Value      string     // the raw value provided by the user
	Detail     string     // the detail shown in the message (the raw value or the validation error)
	Cause      error      // the underlying parser or validator error, if any
}

func (e *InvalidPositionalValueError) Error() string {
	return fmt.Sprintf(GetLocale().ErrInvalidPositionalValue, e.Positional.Name(), e.Detail)
}
func (e *InvalidPositionalValueError) Unwrap() error { return e.Cause }

// MissingRequiredFlagError is returned when a required flag is not provided.
type MissingRequiredFlagError struct {
	Flag Flag
}

func (e *MissingRequiredFlagError) Error() string {
	return fmt.Sprintf(GetLocale().ErrMissingRequiredFlag, e.Flag.Signature())
}

// MissingRequiredPositionalError is returned when a required positional
// argument is not provided.
type MissingRequiredPositionalError struct {
	Positional Positional
}

func (e *MissingRequiredPositionalError) Error() string {
	return fmt.Sprintf(GetLocale().ErrMissingRequiredPositional, e.Positional.Name())
}

// UnknownFlagError is returned when the user provides a flag that is not
// defined in the command and extra flags are not allowed.
type UnknownFlagError struct {
	Name string // the flag name as provided by the user, without dashes
}

func (e *UnknownFlagError) Error() string {
	return fmt.Sprintf(GetLocale().ErrUnknownFlag, e.Name)
}

// RepeatedFlagError is returned when a non-repeatable flag is provided more
// than once and repeated flags are not allowed.
type RepeatedFlagError struct {
	Name string // the flag name as provided by the user, without dashes
}

func (e *RepeatedFlagError) Error() string {
	return fmt.Sprintf(GetLocale().ErrFlagSpecifiedMultiple, e.Name)
}

// MissingFlagValueError is returned when a non-boolean flag is provided
// without a value.
type MissingFlagValueError struct {
	Name string // the flag name as provided by the user, without dashes
}

func (e *MissingFlagValueError) Error() string {
	return fmt.Sprintf(GetLocale().ErrMissingValueForFlag, e.Name)
}

// UnexpectedPositionalError is returned when the user provides a positional
// argument that is not defined in the command and extra positionals are not
// allowed.
type UnexpectedPositionalError struct {
	Token string // the raw token provided by the user
}

func (e *UnexpectedPositionalError) Error() string {
	return fmt.Sprintf(GetLocale().ErrUnexpectedExtraPositional, e.Token)
}

// ExclusiveFlagsError is returned when more than one mutually exclusive flag
// is provided. See CheckExclusiveFlags.
type ExclusiveFlagsError struct {
	Flags []Flag // the conflicting flags provided by the user
}

func (e *ExclusiveFlagsError) Error() string {
	return fmt.Sprintf(GetLocale().ErrExclusiveFlags, strings.Join(flagSignatures(e.Flags), " and "))
}

// AtLeastOneFlagError is returned when none of a set of flags is provided but
// at least one is expected. See CheckAnyFlag.
type AtLeastOneFlagError struct {
	Flags []Flag // the expected flags
}

func (e *AtLeastOneFlagError) Error() string {
	return fmt.Sprintf(GetLocale().ErrAtLeastOneFlag, strings.Join(flagSignatures(e.Flags), " or "))
}

func flagSignatures(flags []Flag) []string {
	signatures := make([]string, 0, len(flags))
	for _, flag := range flags {
		signatures = append(signatures, flag.Signature())
	}
	return signatures
}
