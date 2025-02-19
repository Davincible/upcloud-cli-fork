## upctl database start

Start on a managed database

```
upctl database start <UUID/Title...> [flags]
```

### Examples

```
upctl database start b0952286-1193-4a81-a1af-62efc014ae4b
upctl database start b0952286-1193-4a81-a1af-62efc014ae4b 666bcd3c-5c63-428d-a4fd-07c27469a5a6
upctl database start pg-1x1xcpu-2gb-25gb-pl-waw1
```

### Options

```
  -h, --help   help for start
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
