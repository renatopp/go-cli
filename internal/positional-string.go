package internal

type StringPositional struct {
	*BasePositional
	Value string
}

func NewStringPositional(index int, name string, description string) *StringPositional {
	return &StringPositional{
		BasePositional: &BasePositional{
			index:       index,
			name:        name,
			description: description,
		},
		Value: "",
	}
}

func (f *StringPositional) fromString(value string) error {
	f.BasePositional.set = true
	f.Value = value
	return nil
}
func (f *StringPositional) WithDefault(value string) *StringPositional {
	f.Value = value
	return f
}
func (f *StringPositional) AsRequired() *StringPositional {
	f.BasePositional.required = true
	return f
}
