package internal

import (
	"fmt"
	"strconv"
)

type IntPositional struct {
	*BasePositional
	Value int
}

func NewIntPositional(index int, name string, description string) *IntPositional {
	return &IntPositional{
		BasePositional: &BasePositional{
			index:       index,
			name:        name,
			description: description,
		},
		Value: 0,
	}
}

func (f *IntPositional) fromString(value string) error {
	f.BasePositional.set = true
	val, err := strconv.Atoi(value)
	f.Value = val
	if err != nil {
		err = fmt.Errorf("invalid integer value for positional argument '%s': %v", f.Name(), value)
	}
	return err
}
func (f *IntPositional) WithDefault(value int) *IntPositional {
	f.Value = value
	return f
}
func (f *IntPositional) AsRequired() *IntPositional {
	f.BasePositional.required = true
	return f
}
