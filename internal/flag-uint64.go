package internal

import (
	"fmt"
	"strconv"
)

type Uint64Flag struct {
	*BaseFlag
	Value uint64
}

func NewUint64Flag(long string, short string, description string) *Uint64Flag {
	return &Uint64Flag{
		BaseFlag: &BaseFlag{
			long:        long,
			short:       short,
			description: description,
		},
		Value: 0,
	}
}

func (f *Uint64Flag) acceptsValue() bool { return true }
func (f *Uint64Flag) fromString(value string) error {
	f.BaseFlag.set = true
	val, err := strconv.ParseUint(value, 10, 64)
	f.Value = val
	if err != nil {
		err = fmt.Errorf("invalid uint64 value for flag `%s`: %v", f.Signature(), value)
	}
	return err
}
func (f *Uint64Flag) WithDefault(value uint64) *Uint64Flag {
	f.Value = value
	f.BaseFlag.defaultSet = true
	return f
}
func (f *Uint64Flag) AsRequired() *Uint64Flag {
	f.BaseFlag.required = true
	return f
}
func (f *Uint64Flag) DefaultValue() string {
	return fmt.Sprintf("%d", f.Value)
}
