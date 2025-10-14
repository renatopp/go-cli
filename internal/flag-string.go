package internal

type StringFlag struct {
	*BaseFlag
	Value string
}

func NewStringFlag(long string, short string, description string) *StringFlag {
	return &StringFlag{
		BaseFlag: &BaseFlag{
			long:        long,
			short:       short,
			description: description,
		},
		Value: "",
	}
}

func (f *StringFlag) acceptsValue() bool { return true }
func (f *StringFlag) fromString(value string) error {
	f.BaseFlag.set = true
	f.Value = value
	return nil
}
func (f *StringFlag) WithDefault(value string) *StringFlag {
	f.Value = value
	return f
}
func (f *StringFlag) AsRequired() *StringFlag {
	f.BaseFlag.required = true
	return f
}
