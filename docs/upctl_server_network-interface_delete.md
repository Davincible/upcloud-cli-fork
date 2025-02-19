## upctl server network-interface delete

Delete a network interface

```
upctl server network-interface delete <UUID/Title/Hostname...> [flags]
```

### Examples

```
upctl server network-interface delete 009d7f4e-99ce-4c78-88f1-e695d4c37743 --index 1
upctl server network-interface delete my_server --index 7
```

### Options

```
      --index int   Interface index.
  -h, --help        help for delete
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

* [upctl server network-interface](upctl_server_network-interface.md)	 - Manage network interface

###### Auto generated by spf13/cobra on 5-Jul-2022
