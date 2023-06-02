---
name: "[SISU] Create Release"
about: Create a release form a selected branch/commit/tag
title: "[SISU] Create Release"
labels: "release"
assignees: ""
---

## Usage

```
Usage:
  sisu create release [flags]

Aliases:
  release, r

Flags:
  -e, --environment string   Select the environment (required)
  -f, --from string          Select the tag, branch or commit (required)
  -h, --help                 help for release
  -i, --increment string     TODO (default "patch")
  -v, --version string       TODO

Global Flags:
      --config string       config file (default is sisu.{yml,yaml})
      --log-format string   Log format (logfmt, json, text)
  -l, --log-level string    Log level (trace, debug, info, warn, error, fatal, panic (default "info")
```

## Basic usage

```bash
/sisu create release --environment <env> --from <branch/commit/tag>
```

### Write above your permutations to create a release :rocket:

---

### Examples

**Create a release with automatic incremental patch**

```
/sisu create release --environment <env> --from <branch/commit/tag>
```

**Create a release with automatic incremental minor**

```
/sisu create release --environment <env> --from <branch/commit/tag> --increment minor
```

**Create a release with automatic incremental mayor**

```
/sisu create release --environment <env> --from <branch/commit/tag> --increment major
```

**Create a release and set a custom release version**

```
/sisu create release --environment <env> --from <branch/commit/tag> --version X.Y.Z
```

### Further information

TODO Aqui va el enlace a la documentacion
