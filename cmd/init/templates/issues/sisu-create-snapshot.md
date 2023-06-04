---
name: "[SISU] Create Snapshot"
about: Create a snapshot form a selected branch/commit/tag
title: "[SISU] Create Snapshot"
labels: "snapshot"
assignees: ""
---

## Usage

```
Usage:
  sisu create snapshot [flags]

Aliases:
  snapshot, s

Flags:
  -e, --environment string   Select the environment (required)
  -f, --from string          Select the tag, branch or commit (required)
  -t, --to string            Select the destination branch (required)
  -h, --help                 help for snapshot

Global Flags:
      --config string       config file (default is sisu.{yml,yaml})
      --log-format string   Log format (logfmt, json, text)
  -l, --log-level string    Log level (trace, debug, info, warn, error, fatal, panic (default "info")
```

## Basic usage

```bash
/sisu create snapshot --environment <env> --from <branch/commit/tag>
```

### Write above your permutations to create a snapshot :rocket:

---

### Examples

**Create a snapshot**

```
/sisu create snapshot --environment <env> --from <branch/commit/tag> --to <branch>
```

### Further information

TODO Aqui va el enlace a la documentacion
