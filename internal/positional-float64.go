package internal

import (
	"fmt"
	"strconv"
)

type FloatPositional struct {
	*BasePositional
	Value float64
}

func NewFloatPositional(index int, name string, description string) *FloatPositional {
	return &FloatPositional{
		BasePositional: &BasePositional{
			index:       index,
			name:        name,
			description: description,
		},
		Value: 0,
	}
}

func (f *FloatPositional) fromString(value string) error {
	f.BasePositional.set = true
	val, err := strconv.ParseFloat(value, 64)
	f.Value = val
	if err != nil {
		err = fmt.Errorf("invalid float64 value for positional argument '%s': %v", f.Name(), value)
	}
	return err
}
func (f *FloatPositional) WithDefault(value float64) *FloatPositional {
	f.Value = value
	return f
}
func (f *FloatPositional) AsRequired() *FloatPositional {
	f.BasePositional.required = true
	return f
}
