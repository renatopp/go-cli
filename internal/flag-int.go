package internal

import (
	"fmt"
	"strconv"
)

type IntFlag struct {
	*BaseFlag
	Value int
}

func NewIntFlag(long string, short string, description string) *IntFlag {
	return &IntFlag{
		BaseFlag: &BaseFlag{
			long:        long,
			short:       short,
			description: description,
		},
		Value: 0,
	}
}

func (f *IntFlag) acceptsValue() bool { return true }
func (f *IntFlag) fromString(value string) error {
	f.BaseFlag.set = true
	val, err := strconv.Atoi(value)
	f.Value = val
	if err != nil {
		err = fmt.Errorf("invalid integer value for flag `%s`: %v", f.Signature(), value)
	}
	return err
}
func (f *IntFlag) WithDefault(value int) *IntFlag {
	f.Value = value
	return f
}
func (f *IntFlag) AsRequired() *IntFlag {
	f.BaseFlag.required = true
	return f
}
