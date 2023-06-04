---
name: "[SISU] Run Deployment"
about: Run deployment pipeline to provision a environment form a selected branch/commit/tag
title: "[SISU] Run Deployment"
labels: "deployment"
assignees: ""
---

## Usage

```
Usage:
  sisu deploy [flags]

Aliases:
  deploy, d

Flags:
  -c, --cloud string         Regex to select which matrix options are select by cloud (default ".*")
  -e, --environment string   Select the environment to be matrixed (required)
  -f, --from string          Select the tag, branch or commit to be matrixed (required)
  -h, --help                 help for matrix
  -n, --name string          Regex to select the matrix options from the list (default ".*")
  -r, --region string        Regex to select which matrix options are select by region (default ".*")

Global Flags:
      --config string       config file (default is sisu.{yml,yaml})
      --log-format string   Log format (logfmt, json, text)
  -l, --log-level string    Log level (trace, debug, info, warn, error, fatal, panic (default "info")
```

## Basic usage

```bash
/sisu deploy --environment <env> --from <branch/commit/tag>
```

### Write above your permutations to run a deployment :rocket:

---

### Examples

**Run a deployment form a code source, in a selected environment**

```
/sisu deploy --environment <env> --from <branch/commit/tag>
```

**Run a deployment form a code source, in a selected environment, filter functions by region**

```
/sisu deploy --environment <env> --from <branch/commit/tag> --region <region>
```

**Run a deployment form a code source, in a selected environment, filter functions by name**

```
/sisu deploy --environment <env> --from <branch/commit/tag> --name <name>
```

**Run a deployment form a code source, in a selected environment, filter functions by cloud**

```
/sisu deploy --environment <env> --from <branch/commit/tag> --cloud <cloud>
```

**Run a deployment form a code source, in a selected environment, filter functions by multiple params**

```
/sisu deploy --environment <env> --from <branch/commit/tag> --name <name> --region <region>
```

### Further information

TODO Aqui va el enlace a la documentacion
