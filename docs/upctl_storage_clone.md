## upctl storage clone

Clone a storage

```
upctl storage clone <UUID/Title...> [flags]
```

### Examples

```
upctl storage clone 015899e0-0a68-4949-85bb-261a99de5fdd --title my_storage_clone --zone fi-hel1
upctl storage clone 015899e0-0a68-4949-85bb-261a99de5fdd --title my_storage_clone2 --zone pl-waw1  --tier maxiops
upctl storage clone "My Storage" --title my_storage_clone3 --zone pl-waw1  --tier maxiops
```

### Options

```
      --tier string    The storage tier to use. (default "hdd")
      --title string   A short, informational description.
      --zone string    The zone in which the storage will be created, e.g. fi-hel1.
  -h, --help           help for clone
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

* [upctl storage](upctl_storage.md)	 - Manage storages

###### Auto generated by spf13/cobra on 5-Jul-2022
