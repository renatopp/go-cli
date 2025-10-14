package internal

import "strconv"

type Int64Positional struct {
	*BasePositional
	Value int64
}

func NewInt64Positional(index int, name string, description string) *Int64Positional {
	return &Int64Positional{
		BasePositional: &BasePositional{
			index:       index,
			name:        name,
			description: description,
		},
		Value: 0,
	}
}

func (f *Int64Positional) fromString(value string) error {
	f.BasePositional.set = true
	val, err := strconv.ParseInt(value, 10, 64)
	f.Value = val
	return err
}
func (f *Int64Positional) WithDefault(value int64) *Int64Positional {
	f.Value = value
	return f
}
func (f *Int64Positional) AsRequired() *Int64Positional {
	f.BasePositional.required = true
	return f
}
