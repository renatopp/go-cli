package core

import (
	"fmt"

	"github.com/renatopp/go-cli/errors"
)

type AnyPositional interface {
	Name() string
	Description() string
	Env() string
	RawDefault() string
	RawValue() string
	HasDefault() bool
	HasEnv() bool
	IsProvided() bool
	IsRequired() bool
	IsHidden() bool
	IsVariadic() bool
	parse(value string) error
}

// Implements a positional argument with parametric type. You can use this to
// create custom positional arguments of any types you want but with default
// behavior.
type Positional[T any] struct {
	description string // description of the positional argument for help text
	raw         string // the raw string value provided by the user
	env         string // the name of the environment variable that can be used to set the flag value
	rawDefault  string // the raw default value for the positional argument
	provided    bool   // whether the positional argument has been provided by the user and parsed successfully
	defaulted   bool   // whether the positional argument has a default value
	required    bool   // whether the positional argument is required
	hidden      bool   // whether the positional argument should be hidden from help output
	variadic    bool   // whether this is a variadic positional argument (e.g., "files...")
	name        string // the name of the positional argument for help text

	value     T                       // the parsed value of the positional argument, only set after parsing. In case of variadic positionals, this will hold the last value provided by the user, and all values will be stored in the `values` field below.
	values    []T                     // the parsed values of a variadic positional argument, only set after parsing
	default_  T                       // the default value of the positional argument
	parser    func(string) (T, error) // base parser function to convert string input to the desired type
	validator func(T) error           // custom validator function, provided by the user
}

// NewPositional creates a new GenericPositional.
func NewPositional[T any](name, description string, parser func(string) (T, error)) *Positional[T] {
	return &Positional[T]{
		description: description,
		name:        name,
		parser:      parser,
	}
}

// WithDefault sets the default value for the positional argument.
func (f *Positional[T]) WithDefault(value T) *Positional[T] {
	f.default_ = value
	f.rawDefault = fmt.Sprintf("%v", value)
	f.defaulted = true
	return f
}

// WithValidation allow further validation to an existing positional argument
// type. The validation function should return an error if the value is invalid,
// which will be used to provide better error messages to the user.
func (f *Positional[T]) WithValidation(validator func(T) error) *Positional[T] {
	f.validator = validator
	return f
}

// WithEnv sets the name of the environment variable that can be used to set the
// positional argument value. If the user does not provide a value for this
// positional argument, the parser will check if the environment variable is set
// and use its value instead.
func (f *Positional[T]) WithEnv(env string) *Positional[T] {
	f.env = env
	return f
}

// AsRequired marks the positional argument as required, meaning the user must
// provide a value for it.
func (f *Positional[T]) AsRequired() *Positional[T] {
	f.required = true
	return f
}

// AsVariadic marks the positional argument as variadic, meaning it can accept
// multiple values. This should only be used for the last positional argument of
// a command, and it will collect all remaining positional arguments provided by
// the user into a slice of values of type T.
func (f *Positional[T]) AsVariadic() *Positional[T] {
	f.variadic = true
	return f
}

// AsHidden marks the positional argument as hidden, so it is omitted from help output.
func (f *Positional[T]) AsHidden() *Positional[T] {
	f.hidden = true
	return f
}

// Name returns the name of the positional argument for help text.
func (f *Positional[T]) Name() string { return f.name }

// Description returns the description of the positional argument for help text.
func (f *Positional[T]) Description() string { return f.description }

// Env returns the name of the environment variable that can be used to set the
// positional argument value.
func (f *Positional[T]) Env() string { return f.env }

// RawValue returns the raw string value provided by the user for this positional argument.
func (f *Positional[T]) RawValue() string { return f.raw }

// RawDefault returns the raw default value for the positional argument as a
// string. Used for help text and error messages.
func (f *Positional[T]) RawDefault() string { return f.rawDefault }

// HasDefault returns true if the positional argument has a default value.
func (f *Positional[T]) HasDefault() bool { return f.defaulted }

// HasEnv returns true if the positional argument can be set from an environment variable.
func (f *Positional[T]) HasEnv() bool { return f.env != "" }

// IsProvided returns true if the positional argument has been provided by the user and parsed successfully.
func (f *Positional[T]) IsProvided() bool { return f.provided }

// IsRequired returns true if the positional argument is required.
func (f *Positional[T]) IsRequired() bool { return f.required }

// IsHidden returns true if the positional argument should be hidden from help output.
func (f *Positional[T]) IsHidden() bool { return f.hidden }

// IsVariadic returns true if this is a variadic positional argument (e.g., "files...").
func (f *Positional[T]) IsVariadic() bool { return f.variadic }

// Default returns the default value for the positional argument.
func (f *Positional[T]) Default() T { return f.default_ }

// Value returns the parsed value OR the default value if there is one.
// In case of variadic positionals, this will return the last value provided by
// the user, and all values will be stored in the `values` field.
func (f *Positional[T]) Value() T {
	if f.IsProvided() {
		return f.value
	}
	return f.default_
}

// Values returns the parsed values for a variadic positional argument. For
// non-variadic positionals, this will return a slice with a single value
// (the one returned by Value()) or an empty slice
func (f *Positional[T]) Values() []T {
	if f.IsProvided() {
		return f.values
	}
	if f.HasDefault() {
		return []T{f.default_}
	}
	return []T{}
}

// Count returns the number of times the positional argument was specified by
// the user. For non-variadic positionals, this will be either 0 or 1.
func (f *Positional[T]) Count() int { return len(f.values) }

// Parse implements the parsing logic for the generic positional argument.
func (f *Positional[T]) parse(value string) error {
	parsedValue, err := f.parser(value)
	if err != nil {
		return errors.NewInvalidPosValueError(f.Name(), value, err)
	}

	f.value = parsedValue
	f.values = append(f.values, parsedValue)
	f.provided = true
	if f.validator != nil {
		if err := f.validator(parsedValue); err != nil {
			return errors.NewInvalidPosValueError(f.Name(), err.Error(), err)
		}
	}

	return nil
}
