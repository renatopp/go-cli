package internal

import "fmt"

type BaseFlag struct {
	long        string
	short       string
	description string
	raw         string
	parsed      bool
	required    bool
	set         bool
}

func (f *BaseFlag) acceptsValue() bool  { return true }
func (f *BaseFlag) isParsed() bool      { return f.parsed }
func (f *BaseFlag) setParsed()          { f.parsed = true }
func (f *BaseFlag) Long() string        { return f.long }
func (f *BaseFlag) Short() string       { return f.short }
func (f *BaseFlag) Description() string { return f.description }
func (f *BaseFlag) IsSet() bool         { return f.set }
func (f *BaseFlag) Raw() string         { return f.raw }
func (f *BaseFlag) IsRequired() bool    { return f.required }
func (f *BaseFlag) Signature() string {
	if f.Long() != "" && f.Short() != "" {
		return fmt.Sprintf("--%s, -%s", f.Long(), f.Short())
	}
	if f.Long() != "" {
		return fmt.Sprintf("--%s", f.Long())
	}
	if f.Short() != "" {
		return fmt.Sprintf("-%s", f.Short())
	}
	return ""
}

type Flag interface {
	acceptsValue() bool
	fromString(value string) error
	isParsed() bool
	setParsed()
	Long() string
	Short() string
	Description() string
	IsSet() bool
	Raw() string
	IsRequired() bool
	Signature() string
}
