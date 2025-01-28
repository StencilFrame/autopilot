# DoNot - [Do Not]hing Scripting Automation.

DoNot is a tool designed to gradually automate repetitive tasks by defining them in a runbook. A runbook is a series of steps that can be executed manually or automatically. This tool supports defining runbooks in both simple Markdown and extensible YAML formats.

Inspired by [Do Nothing Scripting](https://blog.danslimmon.com/2019/07/15/do-nothing-scripting-the-key-to-gradual-automation) by Dan Slimmon.

If you want to read more about the design and idea behind DoNot, check out the [design document](docs/DESIGN.md) and [idea document](docs/IDEA.md).

## Supported features

- Define runbooks in Markdown or YAML
- Supports manual and shell steps

## Roadmap

This is an early MVP version of DoNot. The following features are planned for future releases:

- Keep execution track to resume from the last step
- Context management for storing and retrieving variables during execution
- Logging of step execution
- Support for more complex step types (e.g., input, conditional, nested steps, etc.)
- Support Runbook type overrides (e.g., `--type=markdown` or `--type=yaml`)
- Support complex executors (e.g., Docker, API calls, etc.)
- Better CLI interface for managing runbooks and execution
- Web UI for managing runbooks and execution
- Support different distributions (e.g., Homebrew, APT, etc.)
- Support for different platforms (e.g., Windows, macOS, Linux, etc.)
- Add SDK support for runbooks
- Add support for plugins, notifications and scheduling

If you have any feature requests or suggestions, please open an issue on GitHub. Pull requests are also welcome!

If you want us to prioritize a feature, please thumbs up the issue or comment on it.

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

1. Initialize the environment (this is a manual step)
   Ensure all prerequisites are installed.
2. Run setup script (this is an automatic step)
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

Runbook types are automatically detected based on the file extension:
- `.md` files are considered Markdown runbooks.
- `.yml` or `.yaml` files are considered YAML runbooks.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request on GitHub.

## Support

If you have any questions or need help, please open an issue on GitHub or find us on [Slack](https://join.slack.com/t/stencilframesupport/shared_invite/zt-2ynp05the-4~kanvoSa~HTHxZCUDuKEg)

## Sponsoring

If you find this project helpful, please consider sponsoring it. Your support helps us maintain and improve the project.
