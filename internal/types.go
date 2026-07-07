package internal

type Flag interface {
	Long() string
	Short() string
	Description() string
	RawValue() string
	Parse(value string) error
	Count() int
	IsParsed() bool
	IsRequired() bool
	IsHidden() bool
	IsRepeatable() bool
	IsRepeated() bool
	IsGlobal() bool
	HasDefault() bool
	RawDefault() string
	Signature() string
	SetRawDefault(rawDefault string)
}

type Positional interface {
	Name() string
	Description() string
	RawValue() string
	Parse(value string) error
	IsParsed() bool
	IsRequired() bool
	IsHidden() bool
	IsVariadic() bool
	HasDefault() bool
	RawDefault() string
	SetRawDefault(rawDefault string)
}
