## upctl storage backup create

Create backup of a storage

```
upctl storage backup create <UUID/Title...> [flags]
```

### Examples

```
upctl storage backup create 01cbea5e-eb5b-4072-b2ac-9b635120e5d8 --title "first backup"
upctl storage backup create "My Storage" --title second_backup
```

### Options

```
      --title string   A short, informational description.
  -h, --help           help for create
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

* [upctl storage backup](upctl_storage_backup.md)	 - Manage backups

###### Auto generated by spf13/cobra on 5-Jul-2022
