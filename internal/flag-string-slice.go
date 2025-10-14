package internal

import "strings"

type StringSliceFlag struct {
	*BaseFlag
	Value []string
}

func NewStringSliceFlag(long string, short string, description string) *StringSliceFlag {
	return &StringSliceFlag{
		BaseFlag: &BaseFlag{
			long:        long,
			short:       short,
			description: description,
		},
		Value: []string{},
	}
}

func (f *StringSliceFlag) acceptsValue() bool { return true }
func (f *StringSliceFlag) fromString(value string) error {
	f.BaseFlag.set = true
	f.Value = strings.Split(value, ",")
	return nil
}
func (f *StringSliceFlag) WithDefault(value []string) *StringSliceFlag {
	f.Value = value
	return f
}
func (f *StringSliceFlag) AsRequired() *StringSliceFlag {
	f.BaseFlag.required = true
	return f
}
