package internal

import "fmt"

type BoolFlag struct {
	*BaseFlag
	Value bool
}

func NewBoolFlag(long string, short string, description string) *BoolFlag {
	return &BoolFlag{
		BaseFlag: &BaseFlag{
			long:        long,
			short:       short,
			description: description,
		},
		Value: false,
	}
}

func (f *BoolFlag) acceptsValue() bool { return false }
func (f *BoolFlag) fromString(value string) error {
	f.BaseFlag.set = true
	if value == "true" || value == "1" {
		f.Value = true
		return nil
	}
	if value == "false" || value == "0" {
		f.Value = false
		return nil
	}
	return fmt.Errorf("invalid boolean value for flag `%s`: %v", f.Signature(), value)
}
func (f *BoolFlag) WithDefault(value bool) *BoolFlag {
	f.Value = value
	return f
}
func (f *BoolFlag) AsRequired() *BoolFlag {
	f.BaseFlag.required = true
	return f
}
