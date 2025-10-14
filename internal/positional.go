package internal

type BasePositional struct {
	index       int
	name        string
	description string
	raw         string
	parsed      bool
	required    bool
	set         bool
}

func (a *BasePositional) isParsed() bool      { return a.parsed }
func (a *BasePositional) setParsed()          { a.parsed = true }
func (a *BasePositional) Name() string        { return a.name }
func (a *BasePositional) Index() int          { return a.index }
func (a *BasePositional) Description() string { return a.description }
func (a *BasePositional) IsRequired() bool    { return a.required }
func (a *BasePositional) IsSet() bool         { return a.set }
func (a *BasePositional) Raw() string         { return a.raw }

type Positional interface {
	fromString(value string) error
	isParsed() bool
	setParsed()
	Name() string
	Index() int
	Description() string
	IsRequired() bool
	IsSet() bool
	Raw() string
}
