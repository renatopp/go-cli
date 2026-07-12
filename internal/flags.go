package internal

import "fmt"

type Flag interface {
	Long() string
	Short() string
	Description() string
	RawValue() string
	Parse(value string) error
	Count() int
	IsParsed() bool
	IsRequired() bool
	IsHidden() bool
	IsRepeatable() bool
	IsRepeated() bool
	IsGlobal() bool
	HasDefault() bool
	RawDefault() string
	Signature() string
	SetRawDefault(rawDefault string)
}

// Implements a flag with parametric type. You can use this to create custom
// flags of any types you want but with default behavior.
type GenericFlag[T any] struct {
	description string // description of the flag for help text
	raw         string // the raw string value provided by the user
	rawDefault  string // the raw default value for the flag
	parsed      bool   // whether the flag has been provided by the user and parsed successfully
	defaulted   bool   // whether the flag has a default value
	required    bool   // whether the flag is required
	hidden      bool   // whether the flag should be hidden from help output
	long        string // --name
	short       string // -n
	repeatable  bool   // whether the flag can be specified multiple times
	global      bool   // whether the flag is global

	value     T                       // the parsed value of the flag, only set after parsing. In case of repeatable flags, this will hold the last value provided by the user, and all values will be stored in the `values` field below.
	values    []T                     // the parsed values of a repeatable flag, only set after parsing. For non-repeatable flags, this will be a slice with a single value (the one returned by Value()) or an empty slice if the flag was not provided and has no default.
	default_  T                       // the default value of the flag
	parser    func(string) (T, error) // base parser function to convert string input to the desired type
	validator func(T) error           // custom validator function, provided by the user
}

// NewGenericFlag creates a new GenericFlag.
func NewGenericFlag[T any](long, short, description string, parser func(string) (T, error)) *GenericFlag[T] {
	return &GenericFlag[T]{
		description: description,
		long:        long,
		short:       short,
		parser:      parser,
	}
}

// Description returns the description of the flag for help text.
func (f *GenericFlag[T]) Description() string { return f.description }

// RawValue returns the raw string value provided by the user for this flag.
func (f *GenericFlag[T]) RawValue() string { return f.raw }

// IsParsed returns true if the flag has been provided by the user and parsed successfully.
func (f *GenericFlag[T]) IsParsed() bool { return f.parsed }

// IsRequired returns true if the flag is required.
func (f *GenericFlag[T]) IsRequired() bool { return f.required }

// IsHidden returns true if the flag should be hidden from help output.
func (f *GenericFlag[T]) IsHidden() bool { return f.hidden }

// HasDefault returns true if the flag has a default value.
func (f *GenericFlag[T]) HasDefault() bool { return f.defaulted }

// RawDefault returns the raw default value for the flag as a string.
// Used for help text and error messages.
func (f *GenericFlag[T]) RawDefault() string { return f.rawDefault }

// SetRawDefault sets the raw default value for the flag and marks it as
// having a default value.
func (f *GenericFlag[T]) SetRawDefault(rawDefault string) {
	f.rawDefault = rawDefault
	f.defaulted = true
}

// Long returns the long name of the flag (e.g., "name" for --name).
func (f *GenericFlag[T]) Long() string { return f.long }

// Short returns the short name of the flag (e.g., "n" for -n).
func (f *GenericFlag[T]) Short() string { return f.short }

// IsRepeatable returns true if the flag can be specified multiple times.
func (f *GenericFlag[T]) IsRepeatable() bool { return f.repeatable }

// IsGlobal returns true if the flag is global, meaning it can be used in any subcommand.
func (f *GenericFlag[T]) IsGlobal() bool { return f.global }

// Signature returns the flag's signature for help text, combining both long
// and short names if available.
func (f *GenericFlag[T]) Signature() string {
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

// Value returns the parsed value OR the default value if there is one.
// In case of repeatable flags, this will return the last value provided by the
// user, and all values will be stored in the `Values`.
func (f *GenericFlag[T]) Value() T {
	if f.IsParsed() {
		return f.value
	}
	return f.default_
}

// Values returns the parsed values for a repeatable flag. For non-repeatable flags, this will return a slice with a single value (the one returned by Value()) or an empty slice if the flag was not provided and has no default.
func (f *GenericFlag[T]) Values() []T {
	if f.IsParsed() {
		return f.values
	}
	if f.HasDefault() {
		return []T{f.default_}
	}
	return []T{}
}

// Count returns the number of times the flag was specified by the user.
// For non-repeatable flags, this will be either 0 or 1.
func (f *GenericFlag[T]) Count() int { return len(f.values) }

// WithDefault sets the default value for the flag.
func (f *GenericFlag[T]) WithDefault(value T) *GenericFlag[T] {
	f.default_ = value
	f.SetRawDefault(fmt.Sprintf("%v", value))
	return f
}

// WithValidation allow further validation to an existing flag type. The
// validation function should return an error if the value is invalid, which
// will be used to provide better error messages to the user.
func (f *GenericFlag[T]) WithValidation(validator func(T) error) *GenericFlag[T] {
	f.validator = validator
	return f
}

// AsRequired marks the flag as required, meaning the user must provide a value
// for it.
func (f *GenericFlag[T]) AsRequired() *GenericFlag[T] {
	f.required = true
	return f
}

// IsRepeated returns true if the flag has been specified multiple times.
func (f *GenericFlag[T]) IsRepeated() bool { return len(f.values) > 1 }

// AsRepeatable marks the flag as repeatable, meaning the user can specify it multiple times. All values provided by the user will be stored in a slice of values of type T, which can be accessed using the Values() method. For non-repeatable flags, the Values() method will return a slice with a single value (the one returned by Value()) or an empty slice if the flag was not provided and has no default.
func (f *GenericFlag[T]) AsRepeatable() *GenericFlag[T] {
	f.repeatable = true
	return f
}

// AsGlobal marks the flag as global, meaning it can be used in any subcommand.
func (f *GenericFlag[T]) AsGlobal() *GenericFlag[T] {
	f.global = true
	return f
}

// AsHidden marks the flag as hidden, so it is omitted from help output.
func (f *GenericFlag[T]) AsHidden() *GenericFlag[T] {
	f.hidden = true
	return f
}

// Parse implements the parsing logic for the generic flag.
func (f *GenericFlag[T]) Parse(value string) error {
	parsedValue, err := f.parser(value)
	if err != nil {
		return &InvalidFlagValueError{Flag: f, Value: value, Detail: value, Cause: err}
	}
	f.parsed = true
	f.value = parsedValue
	f.values = append(f.values, parsedValue)
	if f.validator != nil {
		if err := f.validator(parsedValue); err != nil {
			return &InvalidFlagValueError{Flag: f, Value: value, Detail: err.Error(), Cause: err}
		}
	}

	return nil
}
