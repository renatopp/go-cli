package internal

import (
	"fmt"
)

var app *App

func init() {
	Clear()
}

func Clear() {
	app = &App{
		cmd:    NewCommand(),
		strict: false,
		printf: func(format string, a ...any) {
			fmt.Printf(format+"\n", a...)
		},
	}
}

func GetApp() *App {
	return app
}

func GetCmd() *Command {
	return app.cmd
}
