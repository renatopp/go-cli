package pkg

import (
	"fmt"

	cerrors "github.com/renatopp/go-cli/pkg/errors"
)

type Positional interface {
	Name() string
	Description() string
	RawValue() string
	Parse(value string) error
	IsParsed() bool
	IsRequired() bool
	IsHidden() bool
	IsVariadic() bool
	HasDefault() bool
	RawDefault() string
	SetRawDefault(rawDefault string)
}

// Implements a positional argument with parametric type. You can use this to
// create custom positional arguments of any types you want but with default
// behavior.
type GenericPositional[T any] struct {
	description string // description of the positional argument for help text
	raw         string // the raw string value provided by the user
	rawDefault  string // the raw default value for the positional argument
	parsed      bool   // whether the positional argument has been provided by the user and parsed successfully
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

// NewGenericPositional creates a new GenericPositional.
func NewGenericPositional[T any](name, description string, parser func(string) (T, error)) *GenericPositional[T] {
	return &GenericPositional[T]{
		description: description,
		name:        name,
		parser:      parser,
	}
}

// WithDefault sets the default value for the positional argument.
func (f *GenericPositional[T]) WithDefault(value T) *GenericPositional[T] {
	f.default_ = value
	f.SetRawDefault(fmt.Sprintf("%v", value))
	return f
}

// WithValidation allow further validation to an existing positional argument
// type. The validation function should return an error if the value is invalid,
// which will be used to provide better error messages to the user.
func (f *GenericPositional[T]) WithValidation(validator func(T) error) *GenericPositional[T] {
	f.validator = validator
	return f
}

// AsRequired marks the positional argument as required, meaning the user must
// provide a value for it.
func (f *GenericPositional[T]) AsRequired() *GenericPositional[T] {
	f.required = true
	return f
}

// AsVariadic marks the positional argument as variadic, meaning it can accept
// multiple values. This should only be used for the last positional argument of
// a command, and it will collect all remaining positional arguments provided by
// the user into a slice of values of type T.
func (f *GenericPositional[T]) AsVariadic() *GenericPositional[T] {
	f.variadic = true
	return f
}

// AsHidden marks the positional argument as hidden, so it is omitted from help output.
func (f *GenericPositional[T]) AsHidden() *GenericPositional[T] {
	f.hidden = true
	return f
}

// Name returns the name of the positional argument for help text.
func (f *GenericPositional[T]) Name() string { return f.name }

// Description returns the description of the positional argument for help text.
func (f *GenericPositional[T]) Description() string { return f.description }

// RawValue returns the raw string value provided by the user for this positional argument.
func (f *GenericPositional[T]) RawValue() string { return f.raw }

// RawDefault returns the raw default value for the positional argument as a
// string. Used for help text and error messages.
func (f *GenericPositional[T]) RawDefault() string { return f.rawDefault }

// HasDefault returns true if the positional argument has a default value.
func (f *GenericPositional[T]) HasDefault() bool { return f.defaulted }

// IsParsed returns true if the positional argument has been provided by the user and parsed successfully.
func (f *GenericPositional[T]) IsParsed() bool { return f.parsed }

// IsRequired returns true if the positional argument is required.
func (f *GenericPositional[T]) IsRequired() bool { return f.required }

// IsHidden returns true if the positional argument should be hidden from help output.
func (f *GenericPositional[T]) IsHidden() bool { return f.hidden }

// IsVariadic returns true if this is a variadic positional argument (e.g., "files...").
func (f *GenericPositional[T]) IsVariadic() bool { return f.variadic }

// Value returns the parsed value OR the default value if there is one.
// In case of variadic positionals, this will return the last value provided by
// the user, and all values will be stored in the `values` field.
func (f *GenericPositional[T]) Value() T {
	if f.IsParsed() {
		return f.value
	}
	return f.default_
}

// Values returns the parsed values for a variadic positional argument. For
// non-variadic positionals, this will return a slice with a single value
// (the one returned by Value()) or an empty slice
func (f *GenericPositional[T]) Values() []T {
	if f.IsParsed() {
		return f.values
	}
	if f.HasDefault() {
		return []T{f.default_}
	}
	return []T{}
}

// Count returns the number of times the positional argument was specified by
// the user. For non-variadic positionals, this will be either 0 or 1.
func (f *GenericPositional[T]) Count() int { return len(f.values) }

// Parse implements the parsing logic for the generic positional argument.
func (f *GenericPositional[T]) Parse(value string) error {
	parsedValue, err := f.parser(value)
	if err != nil {
		return cerrors.NewInvalidPosValueError(f.Name(), value, err)
	}

	f.value = parsedValue
	f.values = append(f.values, parsedValue)
	f.parsed = true
	if f.validator != nil {
		if err := f.validator(parsedValue); err != nil {
			return cerrors.NewInvalidPosValueError(f.Name(), err.Error(), err)
		}
	}

	return nil
}

// SetRawDefault sets the raw default value for the positional argument and
// marks it as having a default value.
func (f *GenericPositional[T]) SetRawDefault(rawDefault string) {
	f.rawDefault = rawDefault
	f.defaulted = true
}
