package core

import (
	"fmt"

	"github.com/renatopp/go-cli/errors"
)

type AnyFlag interface {
	Long() string
	Short() string
	Description() string
	Env() string
	RawValue() string
	RawDefault() string
	Signature() string
	Count() int
	HasDefault() bool
	HasEnv() bool
	IsProvided() bool
	IsRequired() bool
	IsHidden() bool
	IsRepeatable() bool
	IsRepeated() bool
	IsGlobal() bool
	parse(value string) error
	onParsed()
}

// Implements a flag with parametric type. You can use this to create custom
// flags of any types you want but with default behavior.
type Flag[T any] struct {
	description      string         // description of the flag for help text
	raw              string         // the raw string value provided by the user
	env              string         // the name of the environment variable that can be used to set the flag value
	rawDefault       string         // the raw default value for the flag, used for help text and error messages
	provided         bool           // whether the flag has been provided by the user and parsed successfully
	defaulted        bool           // whether the flag has a default value
	required         bool           // whether the flag is required
	hidden           bool           // whether the flag should be hidden from help output
	long             string         // --name
	short            string         // -n
	repeatable       bool           // whether the flag can be specified multiple times
	global           bool           // whether the flag is global
	onParsedCallback func(*Flag[T]) // callback function to be called after parsing

	value     T                       // the parsed value of the flag, only set after parsing. In case of repeatable flags, this will hold the last value provided by the user, and all values will be stored in the `values` field below.
	values    []T                     // the parsed values of a repeatable flag, only set after parsing. For non-repeatable flags, this will be a slice with a single value (the one returned by Value()) or an empty slice if the flag was not provided and has no default.
	default_  T                       // the default value of the flag
	parser    func(string) (T, error) // base parser function to convert string input to the desired type
	validator func(T) error           // custom validator function, provided by the user
}

// NewFlag creates a new GenericFlag.
func NewFlag[T any](long, short, description string, parser func(string) (T, error)) *Flag[T] {
	return &Flag[T]{
		description: description,
		long:        long,
		short:       short,
		parser:      parser,
	}
}

// WithDefault sets the default value for the flag.
func (f *Flag[T]) WithDefault(value T) *Flag[T] {
	f.default_ = value
	f.rawDefault = fmt.Sprintf("%v", value)
	f.defaulted = true
	return f
}

// WithValidation allow further validation to an existing flag type. The
// validation function should return an error if the value is invalid, which
// will be used to provide better error messages to the user.
func (f *Flag[T]) WithValidation(validator func(T) error) *Flag[T] {
	f.validator = validator
	return f
}

// WithEnv allows the flag to be set from an environment variable. The environment variable
// will be used if the flag is not provided by the user. Notice that the default value will
// only be used if the flag is not provided and the environment variable is not set.
func (f *Flag[T]) WithEnv(name string) *Flag[T] {
	f.env = name
	return f
}

// OnParsed registers a callback function that will be called after all arguments have
// been parsed successfully, just before resuming the execution of the program following
// the `Parse` call.
//
// You can use this callback to define a base behavior for a global flag.
//
// The callback receives the flag itself as an argument.
func (f *Flag[T]) OnParsed(cb func(flag *Flag[T])) {
	f.onParsedCallback = cb
}

// AsRequired marks the flag as required, meaning the user must provide a value
// for it.
func (f *Flag[T]) AsRequired() *Flag[T] {
	f.required = true
	return f
}

// AsRepeatable marks the flag as repeatable, meaning the user can specify it multiple times. All values provided by the user will be stored in a slice of values of type T, which can be accessed using the Values() method. For non-repeatable flags, the Values() method will return a slice with a single value (the one returned by Value()) or an empty slice if the flag was not provided and has no default.
func (f *Flag[T]) AsRepeatable() *Flag[T] {
	f.repeatable = true
	return f
}

// AsGlobal marks the flag as global, meaning it can be used in any subcommand.
func (f *Flag[T]) AsGlobal() *Flag[T] {
	f.global = true
	return f
}

// AsHidden marks the flag as hidden, so it is omitted from help output.
func (f *Flag[T]) AsHidden() *Flag[T] {
	f.hidden = true
	return f
}

// Long returns the long name of the flag (e.g., "name" for --name).
func (f *Flag[T]) Long() string { return f.long }

// Short returns the short name of the flag (e.g., "n" for -n).
func (f *Flag[T]) Short() string { return f.short }

// Description returns the description of the flag for help text.
func (f *Flag[T]) Description() string { return f.description }

// Env returns the name of the environment variable that can be used to set the flag value.
func (f *Flag[T]) Env() string { return f.env }

// RawValue returns the raw string value provided by the user for this flag.
func (f *Flag[T]) RawValue() string { return f.raw }

// RawDefault returns the raw default value for the flag as a string.
// Used for help text and error messages.
func (f *Flag[T]) RawDefault() string { return f.rawDefault }

// HasDefault returns true if the flag has a default value.
func (f *Flag[T]) HasDefault() bool { return f.defaulted }

// HasEnv returns true if the flag can be set from an environment variable.
func (f *Flag[T]) HasEnv() bool { return f.env != "" }

// IsProvided returns true if the flag has been provided by the user and parsed successfully.
func (f *Flag[T]) IsProvided() bool { return f.provided }

// IsRequired returns true if the flag is required.
func (f *Flag[T]) IsRequired() bool { return f.required }

// IsHidden returns true if the flag should be hidden from help output.
func (f *Flag[T]) IsHidden() bool { return f.hidden }

// IsRepeatable returns true if the flag can be specified multiple times.
func (f *Flag[T]) IsRepeatable() bool { return f.repeatable }

// IsGlobal returns true if the flag is global, meaning it can be used in any subcommand.
func (f *Flag[T]) IsGlobal() bool { return f.global }

// IsRepeated returns true if the flag has been specified multiple times.
func (f *Flag[T]) IsRepeated() bool { return len(f.values) > 1 }

// Signature returns the flag's signature for help text, combining both long
// and short names if available.
func (f *Flag[T]) Signature() string {
	if f.long != "" && f.short != "" {
		return fmt.Sprintf("-%s, --%s", f.short, f.long)
	}
	if f.long != "" {
		return "--" + f.long
	}
	if f.short != "" {
		return "-" + f.short
	}
	return ""
}

// Default returns the default value for the flag.
func (f *Flag[T]) Default() T {
	return f.default_
}

// Value returns the parsed value OR the default value if there is one.
// In case of repeatable flags, this will return the last value provided by the
// user, and all values will be stored in the `Values`.
func (f *Flag[T]) Value() T {
	if f.IsProvided() {
		return f.value
	}
	return f.default_
}

// Values returns the parsed values for a repeatable flag. For non-repeatable flags, this will return a slice with a single value (the one returned by Value()) or an empty slice if the flag was not provided and has no default.
func (f *Flag[T]) Values() []T {
	if f.IsProvided() {
		return f.values
	}
	if f.HasDefault() {
		return []T{f.default_}
	}
	return []T{}
}

// Count returns the number of times the flag was specified by the user.
// For non-repeatable flags, this will be either 0 or 1.
func (f *Flag[T]) Count() int { return len(f.values) }

// Parse implements the parsing logic for the generic flag.
func (f *Flag[T]) parse(value string) error {
	parsedValue, err := f.parser(value)
	if err != nil {
		return errors.NewInvalidFlagValueError(f.Signature(), value, err)
	}
	f.provided = true
	f.value = parsedValue
	f.values = append(f.values, parsedValue)
	if f.validator != nil {
		if err := f.validator(parsedValue); err != nil {
			return errors.NewInvalidFlagValueError(f.Signature(), err.Error(), err)
		}
	}

	return nil
}

func (f *Flag[T]) onParsed() {
	if f.onParsedCallback != nil {
		f.onParsedCallback(f)
	}
}
