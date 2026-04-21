package v2

type Flag interface {
	Long() string
	Short() string
	Description() string
	RawValue() string
	Parse(value string) error
	IsParsed() bool
	IsRequired() bool
	IsRepeatable() bool
	IsRepeated() bool
	HasDefault() bool
	RawDefault() string
	Signature() string
	SetRawDefault(rawDefault string)
}

type Positional interface {
	Index() int
	Name() string
	Description() string
	RawValue() string
	Parse(value string) error
	IsParsed() bool
	IsRequired() bool
	IsVariadic() bool
	HasDefault() bool
	RawDefault() string
	SetRawDefault(rawDefault string)
}
