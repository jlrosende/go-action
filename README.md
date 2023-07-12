# go-action

## Usage

```
Usage:
  sisu [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  deploy      Add git repository containing markdown content files
  help        Help about any command
  release     Add git repository containing markdown content files
  test        Add git repository containing markdown content files

Flags:
      --config string       config file (default is sisu.{yml,yaml})
  -h, --help                help for sisu
      --log-format string   Log format (logfmt, json, text)
  -l, --log-level string    Log level (trace, debug, info, warn, error, fatal, panic (default "info")
  -v, --version             version for sisu

Use "sisu [command] --help" for more information about a command.
```

To use this action you need to configure a file configuration file named `sisu.yaml`.

**_`sisu.yaml`_**

```yaml

```

```bash
sisu --config ./<path-to-config> --log-level <trace|debug|info|warn|error|panic> --log-format
```

### Deploy

```bash
sisu deploy --environtment <env_name> --from <tag|commit|branch>
```

```bash
sisu deploy -e <env_name> -f <tag|commit|branch>
```

### Testing

### Autocompletion

// TODO

You can add autocompletion adding one of this commands in your terminal

#### Bash

```bash
sisu completion bash
```

#### Zsh

```bash
sisu completion zsh
```

#### Powershell

```bash
sisu completion powersehll
```

#### Fish

```bash
sisu completion fish
```
