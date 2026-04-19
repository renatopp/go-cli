package internal

import (
	"fmt"
	"strconv"
)

type UintFlag struct {
	*BaseFlag
	Value uint
}

func NewUintFlag(long string, short string, description string) *UintFlag {
	return &UintFlag{
		BaseFlag: &BaseFlag{
			long:        long,
			short:       short,
			description: description,
		},
		Value: 0,
	}
}

func (f *UintFlag) acceptsValue() bool { return true }
func (f *UintFlag) fromString(value string) error {
	f.BaseFlag.set = true
	val, err := strconv.ParseUint(value, 10, 0)
	f.Value = uint(val)
	if err != nil {
		err = fmt.Errorf("invalid unsigned integer value for flag `%s`: %v", f.Signature(), value)
	}
	return err
}
func (f *UintFlag) WithDefault(value uint) *UintFlag {
	f.Value = value
	f.BaseFlag.defaultSet = true
	return f
}
func (f *UintFlag) AsRequired() *UintFlag {
	f.BaseFlag.required = true
	return f
}
func (f *UintFlag) DefaultValue() string {
	return fmt.Sprintf("%d", f.Value)
}
