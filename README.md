# DoNot - [Do Not]hing Scripting Automation.

DoNot is a tool designed to gradually automate repetitive tasks by defining them in a runbook. A runbook is a series of steps that can be executed manually or automatically. This tool supports defining runbooks in both simple Markdown and extensible YAML formats.

Inspired by [Do Nothing Scripting](https://blog.danslimmon.com/2019/07/15/do-nothing-scripting-the-key-to-gradual-automation) by Dan Slimmon.

## Features

- Define runbooks in Markdown or YAML
- Supports manual and shell steps

## Installation

To install DoNot, clone the repository and build the project:

```sh
git clone https://github.com/stencilframe/donot.git
cd donot
go build -o donot ./pkg/cmd/donot
```

## Usage

### Define a Runbook

Create a runbook file in Markdown YAML or format.

#### Markdown Runbooks

Markdown runbooks are defined by listing ordered steps with instructions and code blocks for shell steps.
DoNot extracts the first ordered list in the file as the runbook steps. If there are multiple ordered lists, only the first one is considered.

DoNot supports and detects two types of steps: manual and shell (automatic).
If a step has a code block, it's considered a shell step; otherwise, it's a manual step.

##### Example

~~~markdown
# Example Runbook

1. Initialize the environment
   Ensure all prerequisites are installed.
2. Run setup script
   ```
   ./setup.sh
   ```
~~~

#### YAML Runbooks

Schema:
- `name`: Runbook name
- `steps`: List of steps
  - `id`: Step ID
  - `type`: Step type (`manual`, `shell`)
  - `name`: Step name
  - Additional fields based on step type (see below)

Additional fields for manual steps:
  - `instructions`: Step instructions (for manual steps)

Additional fields for shell steps:
  - `command`: Step command (for shell steps)

##### Example

```yaml
name: Example Runbook
steps:
  - id: step-1
    type: manual
    name: Initialize the environment
    instructions: |
      Ensure all prerequisites are installed.

  - id: step-2
    type: shell
    name: Run setup script
    command: ./setup.sh
```

### Execute a Runbook

To execute a runbook, use the `donot` command followed by the runbook file:

```sh
./donot run runbook.md
```

Runbook types are detected based on the file extension:
- `.md` files are considered Markdown runbooks.
- `.yml` or `.yaml` files are considered YAML runbooks.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request on GitHub.
