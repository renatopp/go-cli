# Shell Completion

Complete subcommands, flags and values on TAB in bash and zsh.

## Enabling Completion

Enable it on the root command:

```go
func main() {
    cli.Name("ship")
    cli.AutoCompletion(true)

    cli.Command("deploy", "Deploy a service", func() {
        cli.Flag("env", "e", "Target environment")
        cli.Parse()
    })

    cli.Parse()
}
```

This adds a hidden `completion <shell>` command that prints the script your
users install once:

```sh
# bash (~/.bashrc)
source <(ship completion bash)

# zsh (~/.zshrc, after compinit)
source <(ship completion zsh)
# or install it in your fpath
ship completion zsh > "${fpath[1]}/_ship"
```

Package managers can ship the scripts in the shells' autoload directories
(e.g. `/usr/share/bash-completion/completions/ship` and
`/usr/share/zsh/site-functions/_ship`), so completion works right after install.

Out of the box, it completes:

- Subcommand names (hidden commands are skipped)
- Flag names, skipping flags already provided (unless repeatable)
- `true`/`false` for bool flags and positionals
- File names for any other flag value or positional

## Dynamic Values

Use `WithCompletion` on flags and positionals to suggest values. The function
receives the partial value typed so far. Add a description after a tab:

```go
cli.Flag("env", "e", "Target environment").
    WithCompletion(func(prefix string) []string {
        return []string{"dev\tDevelopment", "prod\tProduction"}
    })

cli.Pos("service", "Service name").
    WithCompletion(func(prefix string) []string {
        return listServices(prefix) // e.g. query an API
    })
```

## How It Works

The script calls your program back on every TAB with a hidden argument, e.g.
`ship __complete deploy --env p`. Your commands are executed as usual until
the target command reaches `Parse()`, which prints the candidates instead of
parsing and exits. Because of that:

- **Avoid side effects or output before `Parse()`**, they run on every TAB.
- Upgrading your CLI does not require reinstalling the script.

You can test it the same way:

```go
cli.Stdout(&buf)
cli.UsePanic(true)
cli.AutoCompletion(true)
cli.ParseArgs([]string{"__complete", "deploy", "--env", ""})
```
