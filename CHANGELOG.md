# 0.2.0 ()

- Adding Print and Error functions to use stdout and stderr.

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

