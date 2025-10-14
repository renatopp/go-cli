package internal

import "strconv"

type Int64Flag struct {
	*BaseFlag
	Value int64
}

func NewInt64Flag(long string, short string, description string) *Int64Flag {
	return &Int64Flag{
		BaseFlag: &BaseFlag{
			long:        long,
			short:       short,
			description: description,
		},
		Value: 0,
	}
}

func (f *Int64Flag) acceptsValue() bool { return true }
func (f *Int64Flag) fromString(value string) error {
	f.BaseFlag.set = true
	val, err := strconv.ParseInt(value, 10, 64)
	f.Value = val
	return err
}
func (f *Int64Flag) WithDefault(value int64) *Int64Flag {
	f.Value = value
	return f
}
func (f *Int64Flag) AsRequired() *Int64Flag {
	f.BaseFlag.required = true
	return f
}
