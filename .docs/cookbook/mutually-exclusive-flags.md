# Mutually Exclusive Flags

Enforce that only one of several flags can be provided.

## Using CheckExclusiveFlags

Define flags and validate after parsing:

```go
func main() {
    json := cli.FlagBool("json", "j", "Output as JSON")
    yaml := cli.FlagBool("yaml", "y", "Output as YAML")
    csv := cli.FlagBool("csv", "c", "Output as CSV")
    
    cli.Parse()
    
    cli.CheckExclusiveFlags(json, yaml, csv)  // At most one allowed
    
    if json.IsProvided() {
        // Handle JSON output
    } else if yaml.IsProvided() {
        // Handle YAML output
    } else if csv.IsProvided() {
        // Handle CSV output
    }
}
```

`CheckExclusiveFlags()` exits with an error if more than one flag is provided. It works with any flag type.

## Require At Least One

Use `CheckAnyFlag()` to require one of several flags:

```go
func main() {
    create := cli.FlagBool("create", "", "Create resource")
    delete := cli.FlagBool("delete", "", "Delete resource")
    update := cli.FlagBool("update", "", "Update resource")
    
    cli.Parse()
    
    cli.CheckAnyFlag(create, delete, update)  // At least one required
}
```

This exits with an error if none of the flags are provided.
