// Package errors defines the structured error type used by go-cli for
// parsing and validation failures. A CliError carries a Code identifying the
// kind of failure and the Parameters needed to render a message, but this
// package has no notion of localization or formatting — see the locales
// package for rendering a CliError with a specific Locale.
package errors

import (
	"fmt"
	"strings"
)

// ErrorCode identifies the kind of error a CliError represents.
type ErrorCode string

const (
	ErrUnknownFlag         ErrorCode = "unknown_flag"
	ErrMissingRequiredFlag ErrorCode = "missing_required_flag"
	ErrMissingRequiredPos  ErrorCode = "missing_required_positional"
	ErrInvalidFlagValue    ErrorCode = "invalid_flag_value"
	ErrInvalidPosValue     ErrorCode = "invalid_positional_value"
	ErrRepeatedFlag        ErrorCode = "repeated_flag"
	ErrMissingFlagValue    ErrorCode = "missing_flag_value"
	ErrUnexpectedPos       ErrorCode = "unexpected_positional"
	ErrExclusiveFlags      ErrorCode = "exclusive_flags"
	ErrAtLeastOneFlag      ErrorCode = "at_least_one_flag"
)

// CliError represents a CLI parsing or execution error with a code and
// parameters. Error() returns a generic, unlocalized message; use
// locales.Locale.LocalizedError to render the message in a specific locale.
type CliError struct {
	Code       ErrorCode
	Parameters []any
	Cause      error
}

func (e *CliError) Error() string {
	return fmt.Sprintf("%s: %v", e.Code, e.Parameters)
}

func (e *CliError) Unwrap() error {
	return e.Cause
}

// NewCliError creates a custom CLI error with the provided code and parameters.
func NewCliError(code ErrorCode, params ...any) *CliError {
	return &CliError{
		Code:       code,
		Parameters: params,
	}
}

// NewCliErrorWithCause creates a custom CLI error with a cause.
func NewCliErrorWithCause(code ErrorCode, cause error, params ...any) *CliError {
	return &CliError{
		Code:       code,
		Parameters: params,
		Cause:      cause,
	}
}

// Error constructors for common errors

func NewUnknownFlagError(name string) *CliError {
	return NewCliError(ErrUnknownFlag, name)
}

func NewMissingRequiredFlagError(signature string) *CliError {
	return NewCliError(ErrMissingRequiredFlag, signature)
}

func NewMissingRequiredPosError(name string) *CliError {
	return NewCliError(ErrMissingRequiredPos, name)
}

func NewInvalidFlagValueError(signature, detail string, cause error) *CliError {
	return NewCliErrorWithCause(ErrInvalidFlagValue, cause, signature, detail)
}

func NewInvalidPosValueError(name, detail string, cause error) *CliError {
	return NewCliErrorWithCause(ErrInvalidPosValue, cause, name, detail)
}

func NewRepeatedFlagError(name string) *CliError {
	return NewCliError(ErrRepeatedFlag, name)
}

func NewMissingFlagValueError(name string) *CliError {
	return NewCliError(ErrMissingFlagValue, name)
}

func NewUnexpectedPosError(token string) *CliError {
	return NewCliError(ErrUnexpectedPos, token)
}

func NewExclusiveFlagsError(signatures []string) *CliError {
	return NewCliError(ErrExclusiveFlags, strings.Join(signatures, " & "))
}

func NewAtLeastOneFlagError(signatures []string) *CliError {
	return NewCliError(ErrAtLeastOneFlag, strings.Join(signatures, " | "))
}
