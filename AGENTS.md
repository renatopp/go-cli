# go-cli

- If asked for plan, answer with the plan, dont change code
- If asked for review, analysis or general question, never change the code
- Always look for the least amount of code changes to achieve the goal
- Never commit or change remote

Philosophy:
- This is a command-line interface (CLI) library that parses arguments such as positionals and flags
- It provides help message as convenience
- Works as state machine, where each state is a command or subcommand, each can have their own flags and positionals
- Interface must be declarative, direct and intuitive
- Goal is user can design CLIs without boilerplate, structs or anything more complex than a simple function
- Strict by default, but flexible

Internal Coding Rules:
- No global state apart from the default app
- No external dependencies
- Public methods before private methods
- Public functions before private functions
- Types before functions, but struct declaraction should be followed by its factory function
- App should have the same interface (or superset) as @cli.go
