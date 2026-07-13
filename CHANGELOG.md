# 0.6.0 (2026-07-13)

Major overhaul of the API, file structure, and internal implementation.

- Removing many functions out of the scope of the library, such as Print, Error and Shell.
- Pruning the API surface to remove uncommon flag and positional types 
- Adding support for custom flag and positional types.
- Adding formatters for help and error messages.
- Improving error handling.
- Adding ENV support for flags and positionals.
- Adding OnParsed callback for global flags.
- Improving help and error styles.
- Improving docs.

# 0.5.0 (2026-07-07)

- Replacing stdout and stderr by io.Writers.
- Adding global flags with initial typing exposure.
- Adding locale support.
- Adding Count() on flags and positionals.
- Fixing validation error feedback.
- Fixing shell timeout.
- Fixing exposure of string slices.

# 0.4.0 (2026-06-14)

- Adding SetArgs().
- App now reflects the CLI interface and can be instantiated.

# 0.3.0 (2026-06-04)

- Adding Fatal and FatalIf functions.

# 0.2.0 (2026-06-01)

- Adding Print and Error functions to use stdout and stderr.
- Adding Shell as convenience wrapper to exec.Command.
- Adding AsHidden options on commands, flags and positionals.

# 0.1.1 (2026-04-22)

- Re-adding FlagFloat and PosFloat.

# 0.1.0 (2026-04-21)

- Improving API ergonomics.
  - "Strict mode" is default now but user can disable only the features she wants: extra positionals, extra flags, repeated flags.
  - "Panic or Exit" to stop the flow of the cli, so it can be used in tests and for other edge cases.
  - Custom "Stdout" and "stderr" instead of a single output.
- Adding option for "repeated" flags so one can use `--include a --include b` or `-vvv`.
- Adding option for "variadic" positional, capturing every positional argument.
- Adding "auto version"
- Adding "custom validation" functions for positionals and flags.
- Adding several new positional and flag types.
- Adding helper functions to check exclusive flags.
- Adding test cases.
- Adding examples.

Notice that this version is still in the development releases (0.x.x) and it's API can change in future.

# 0.0.0 (2025-10-14)

- First version.
