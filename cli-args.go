package cli

// NArgs returns the number of positional arguments provided by the user.
// Should be used only after Parse() is called, otherwise it will return 0.
func NArgs() int {
	if !app.IsParsed() {
		return 0
	}
	return len(app.Arguments().Args)
}

// Arg retrieves the value of a positional argument by its index.
// Should be used only after Parse() is called, otherwise it will return an
// empty string.
func Arg(index int) string {
	if !app.IsParsed() {
		return ""
	}
	args := app.Arguments().Args
	if index < 0 || index >= len(args) {
		return ""
	}
	return args[index]
}

// Args retrieves all positional arguments provided by the user.
// Should be used only after Parse() is called, otherwise it will return an
// empty slice.
func Args() []string {
	if !app.IsParsed() {
		return []string{}
	}
	return app.Arguments().Args
}

// ExtraArgs retrieves all extra positional arguments provided by the user, i.e.,
// those that are not defined in the command. Should be used only after Parse() is
// called, otherwise it will return an empty slice.
func ExtraArgs() []string {
	if !app.IsParsed() {
		return []string{}
	}
	return app.Arguments().ExtraArgs
}
