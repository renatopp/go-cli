package internal

import (
	"fmt"
	"strconv"
)

type FloatFlag struct {
	*BaseFlag
	Value float64
}

func NewFloatFlag(long string, short string, description string) *FloatFlag {
	return &FloatFlag{
		BaseFlag: &BaseFlag{
			long:        long,
			short:       short,
			description: description,
		},
		Value: 0,
	}
}

func (f *FloatFlag) acceptsValue() bool { return true }
func (f *FloatFlag) fromString(value string) error {
	f.BaseFlag.set = true
	val, err := strconv.ParseFloat(value, 64)
	f.Value = val
	if err != nil {
		err = fmt.Errorf("invalid float value for flag `%s`: %v", f.Signature(), value)
	}
	return err
}
func (f *FloatFlag) WithDefault(value float64) *FloatFlag {
	f.Value = value
	return f
}
func (f *FloatFlag) AsRequired() *FloatFlag {
	f.BaseFlag.required = true
	return f
}
