## upctl database show

Show database details

```
upctl database show <UUID/Title...> [flags]
```

### Examples

```
upctl database show 9a8effcb-80e6-4a63-a7e5-066a6d093c14
upctl database show my-pg-database
upctl database show my-mysql-database
```

### Options

```
  -h, --help   help for show
```

### Options inherited from parent commands

```
  -t, --client-timeout duration   CLI timeout when using interactive mode on some commands
      --config string             Config file
      --debug                     Print out more verbose debug logs
      --force-colours[=true]      force coloured output despite detected terminal support
      --no-colours[=true]         disable coloured output despite detected terminal support
  -o, --output string             Output format (supported: json, yaml and human) (default "human")
```

### SEE ALSO

* [upctl database](upctl_database.md)	 - Manage databases

###### Auto generated by spf13/cobra on 5-Jul-2022
