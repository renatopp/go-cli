package internal

import "fmt"

// Represents a flag or a positional argument. This is mostly used for
// shared fields between both types of arguments. Internal usage only.
type baseArgument struct {
	description string // description of the argument for help text
	raw         string // the raw string value provided by the user
	rawDefault  string // the raw default value for the argument
	parsed      bool   // whether the argument has been parsed successfully
	defaulted   bool   // whether the argument has a default value
	required    bool   // whether the argument is required
}

// newBaseArgument creates a new baseArgument with the given description.
func newBaseArgument(description string) *baseArgument {
	return &baseArgument{
		description: description,
	}
}

// Description returns the description of the argument for help text.
func (f *baseArgument) Description() string { return f.description }

// RawValue returns the raw string value provided by the user for this argument.
func (f *baseArgument) RawValue() string { return f.raw }

// IsParsed returns true if the argument has been parsed successfully.
func (f *baseArgument) IsParsed() bool { return f.parsed }

// IsRequired returns true if the argument is required.
func (f *baseArgument) IsRequired() bool { return f.required }

// HasDefault returns true if the argument has a default value.
func (f *baseArgument) HasDefault() bool { return f.defaulted }

// RawDefault returns the raw default value for the argument as a string.
// Used for help text and error messages.
func (f *baseArgument) RawDefault() string { return f.rawDefault }

// SetRawDefault sets the raw default value for the argument and marks it as
// having a default value.
func (f *baseArgument) SetRawDefault(rawDefault string) {
	f.rawDefault = rawDefault
	f.defaulted = true
}

// Represents a flag argument, which can be specified with a long name
// (e.g., --name) and/or a short name (e.g., -n).
type BaseFlag struct {
	*baseArgument
	long       string // --name
	short      string // -n
	repeatable bool   // whether the flag can be specified multiple times
}

// NewBaseFlag creates a new BaseFlag with the given long name, short name,
// and description.
func NewBaseFlag(long string, short string, description string) *BaseFlag {
	return &BaseFlag{
		baseArgument: &baseArgument{
			description: description,
		},
		long:       long,
		short:      short,
		repeatable: false,
	}
}

// Long returns the long name of the flag (e.g., "name" for --name).
func (f *BaseFlag) Long() string { return f.long }

// Short returns the short name of the flag (e.g., "n" for -n).
func (f *BaseFlag) Short() string { return f.short }

// IsRepeatable returns true if the flag can be specified multiple times.
func (f *BaseFlag) IsRepeatable() bool { return f.repeatable }

// Signature returns the flag's signature for help text, combining both long
// and short names if available.
func (f *BaseFlag) Signature() string {
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

// Represents a positional argument, which is identified by its position
// in the command's argument list.
type BasePositional struct {
	*baseArgument
	variadic bool   // whether this is a variadic positional argument (e.g., "files...")
	index    int    // the position index of the positional argument
	name     string // the name of the positional argument for help text
}

// NewBasePositional creates a new BasePositional with the given name and
// description.
func NewBasePositional(name string, description string) *BasePositional {
	return &BasePositional{
		baseArgument: &baseArgument{
			description: description,
		},
		name: name,
	}
}

// Index returns the position index of the positional argument. Starting from
// 0.
func (p *BasePositional) Index() int { return p.index }

// Name returns the name of the positional argument for help text.
func (p *BasePositional) Name() string { return p.name }

// IsVariadic returns true if this is a variadic positional argument (e.g., "files...").
func (p *BasePositional) IsVariadic() bool { return p.variadic }
