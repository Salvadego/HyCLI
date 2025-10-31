# HyCLI

<!--toc:start-->
- [HyCLI](#hycli)
  - [Features](#features)
  - [Binaries](#binaries)
  - [Quick Start](#quick-start)
    - [Create a client configuration](#create-a-client-configuration)
    - [Select active client](#select-active-client)
    - [List configured clients](#list-configured-clients)
    - [Inspect active client](#inspect-active-client)
  - [FlexibleSearch](#flexiblesearch)
  - [Impex](#impex)
  - [Scripts](#scripts)
  - [Plugins](#plugins)
  - [Configuration File](#configuration-file)
  - [Installation](#installation)
    - [Requirements](#requirements)
    - [Build](#build)
    - [Auto completion](#auto-completion)
    - [Test connectivity](#test-connectivity)
  - [Contributing](#contributing)
    - [Guidelines](#guidelines)
    - [Code Style](#code-style)
    - [Adding plugins](#adding-plugins)
    - [Submitting changes](#submitting-changes)
<!--toc:end-->

Command-line client for the **SAP Hybris Administration Console (HAC)**.

It manages HAC client configurations and provides sub-CLIs for FlexibleSearch,
Impex and Script execution. It also supports external plugins.
Meaning you can create your own commands and plugins.

## Features

* Manage multiple HAC client profiles
* Auto-export active client to environment variables
* FlexibleSearch execution (SQL and FS)
* Impex import and export
* Groovy / JavaScript / Beanshell script execution
* Plugin discovery via `hycli-*` executables

## Binaries

This repository contains three CLI tools:

| Tool     | Description                                        |
| -------- | -------------------------------------------------- |
| `hycli`  | Core CLI. Manages client configs and loads plugins |
| `flex`   | FlexibleSearch client                              |
| `impex`  | Impex import/export client                         |
| `script` | Script execution client                            |

## Quick Start

### Create a client configuration

```bash
hycli config add \
  --client local \
  --address https://localhost:9002/hac \
  --user admin \
  --password nimda
```

### Select active client

```bash
hycli config select --client local
```

### List configured clients

```bash
hycli config list
```

### Inspect active client

```bash
hycli config current
```

After selecting, environment variables are exported automatically:

* HYCLI_CLIENT_URL
* HYCLI_CLIENT_USER
* HYCLI_CLIENT_PASSWORD
* HYCLI_CLIENT_NAME

These are consumed by `flex`, `impex`, `script`, and plugins.

## FlexibleSearch

Run a FlexibleSearch query:

```bash
flex query "SELECT {pk} FROM {Product}" --max 10 --output json
```

SQL-only:

```bash
flex query "SELECT {pk} FROM {Product}" --sql
```

Run raw SQL:

```bash
flex sql "SELECT * FROM products" --max 100
```

## Impex

Import:

```bash
impex import --file sample.impex
```

Export:

```bash
impex export --file export.impex --output result.zip
```

## Scripts

Run Groovy by file:

```bash
script run --file hello.groovy
```

Run inline:

```bash
script run "print 'Hello world'"
```

Choose script type:

```bash
script run my.js --type javascript
```

## Plugins

Any executable named `hycli-*` becomes an extension command.

Examples:

* `/usr/local/bin/hycli-stats`
* `~/.local/share/hycli/plugins/hycli-check`

Plugins receive all arguments. Shell completion is supported if the plugin implements `__complete`.

## Configuration File

Stored in:

```bash
$HOME/.config/hycli/config.yaml
```

Example content:

```yaml
defaultClient: local
clients:
  local:
    address: https://localhost:9002/hac
    user: admin
    password: nimda
```

The file is auto-generated if missing.

---

## Installation

### Requirements

* Go 1.21+
* Optional: external plugins (named `hycli-*`)

### Build

Clone the repository:

```bash
git clone https://github.com/your-org/hycli.git
cd hycli
```

Build binaries:

```bash
go build -o hycli main.go
go build -o hycli-flex flex/main.go
go build -o hycli-impex impex/main.go
go build -o hycli-script script/main.go
```

Place them into PATH:

```bash
mv hycli hycli-flex hycli-impex hycli-script $GOPATH/bin
```

### Auto completion

Add to your shell configuration:

```bash
## if you use bash
source <(hycli completion bash)

## if you use zsh
source <(hycli completion zsh)

## if you use fish
hycli completion fish | source
```

Reload your shell to apply.

### Test connectivity

```bash
hycli config add --client test --address https://localhost:9002/hac --user admin --password nimda
hycli config select --client test
hycli flex query "SELECT {pk} FROM {Product}" --max 5
```

---

## Contributing

### Guidelines

* Use Go modules
* Keep commands consistent with Cobra flag naming
* Always respect environment variables exported by `hycli`
* New command sets should live in `cmd/<name>` or external plugins

### Code Style

* Run `go fmt ./...`
* Run `go vet ./...`
* Favor composable commands and shared utilities

### Adding plugins

To add a plugin:

1. Create an executable named `hycli-<name>`
2. Put it in:

   * `$HOME/.local/share/hycli/plugins`, or
   * anywhere in `$PATH`
3. HyCLI will discover it automatically:

   ```bash
   hycli <name> ...
   ```

For shell completion, implement handling of:

```bash
__complete
```

Your plugin should optionally output completion candidates line-by-line and end with an
optional directive like `:1`.


### Submitting changes

* Create a feature branch
* Open a pull request with a short description

