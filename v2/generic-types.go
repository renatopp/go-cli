package v2

import "fmt"

// Implements a flag with parametric type. You can use this to create custom
// flags of any types you want but with default behavior.
type GenericFlag[T any] struct {
	*BaseFlag
	value     T                       // the parsed value of the flag, only set after parsing
	default_  T                       // the default value of the flag
	parser    func(string) (T, error) // base parser function to convert string input to the desired type
	validator func(T) error           // custom validator function, provided by the user
}

// NewGenericFlag creates a new GenericFlag.
func NewGenericFlag[T any](long, short, description string, parser func(string) (T, error)) *GenericFlag[T] {
	return &GenericFlag[T]{
		BaseFlag: &BaseFlag{
			baseArgument: newBaseArgument(description),
			long:         long,
			short:        short,
		},
		parser: parser,
	}
}

// Value returns the parsed value OR the default value if there is one.
func (f *GenericFlag[T]) Value() T {
	if f.IsParsed() {
		return f.value
	}
	return f.default_
}

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
	f.BaseFlag.required = true
	return f
}

// Parse implements the parsing logic for the generic flag.
func (f *GenericFlag[T]) Parse(value string) error {
	parsedValue, err := f.parser(value)
	if err != nil {
		return fmt.Errorf("invalid value for flag `%s`: %v", f.Signature(), err)
	}

	f.value = parsedValue
	if f.validator != nil {
		if err := f.validator(parsedValue); err != nil {
			return fmt.Errorf("invalid value for flag `%s`: %v", f.Signature(), err)
		}
	}

	return nil
}

// Implements a positional argument with parametric type. You can use this to
// create custom positional arguments of any types you want but with default
// behavior.
type GenericPositional[T any] struct {
	*BasePositional
	value     T                       // the parsed value of the positional argument, only set after parsing
	default_  T                       // the default value of the positional argument
	parser    func(string) (T, error) // base parser function to convert string input to the desired type
	validator func(T) error           // custom validator function, provided by the user
}

// NewGenericPositional creates a new GenericPositional.
func NewGenericPositional[T any](name, description string, parser func(string) (T, error)) *GenericPositional[T] {
	return &GenericPositional[T]{
		BasePositional: &BasePositional{
			baseArgument: newBaseArgument(description),
			name:         name,
		},
		parser: parser,
	}
}

// Value returns the parsed value OR the default value if there is one.
func (f *GenericPositional[T]) Value() T {
	if f.IsParsed() {
		return f.value
	}
	return f.default_
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
	f.BasePositional.required = true
	return f
}

// Parse implements the parsing logic for the generic positional argument.
func (f *GenericPositional[T]) Parse(value string) error {
	parsedValue, err := f.parser(value)
	if err != nil {
		return fmt.Errorf("invalid value for positional argument `%s`: %v", f.Name(), err)
	}

	f.value = parsedValue
	if f.validator != nil {
		if err := f.validator(parsedValue); err != nil {
			return fmt.Errorf("invalid value for positional argument `%s`: %v", f.Name(), err)
		}
	}

	return nil
}
