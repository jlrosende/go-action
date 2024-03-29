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
      --force-rebuild        Force build and compilation of the artifact.
  -f, --from string          Select the origin tag, branch or commit (required)
  -h, --help                 help for snapshot
  -t, --to string            Select the destination branch (required)

Global Flags:
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

**Create a snapshot and force rebuild artifact**

```
/sisu create snapshot --environment <env> --from <branch/commit/tag> --to <branch> --force-rebuild
```

### Further information

TODO Aqui va el enlace a la documentacion
