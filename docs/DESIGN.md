# AutoPilot Design (WIP)

AutoPilot is a tool to automate and document your daily commands and workflows. It is designed to be simple, flexible, extensible and easy shareable with others. You can define commands and runbooks in simple Markdown format and commit them to your project repository.

## Library

Library is a collection of commands and runbooks (workflows).
They might be stored in a different formats: Markdown, YAML, JSON, etc.
They vary in complexity and extensibility.

## Markdown Library

NOTE: Our initial implementation will focus on Markdown format only, because it is human-readable.

Markdown library items are separated by headings:
* If a heading contains a code block, it is considered a command.
* If a heading contains an ordered list, it is considered a runbook.
* Otherwise, it is ignored.

~~~markdown
# Example Library

## Test application command

This is a command to test the application.

```sh
./test.sh
```

## Build application command

This is a command to build the application.

```sh
./build.sh
```

## Deploy application command

This is a command to deploy the application.

```sh
./deploy.sh
```

## Configure Development Environment Runbook

This is a runbook for configuring the development environment.

1. Install prerequisites (manual step)
   Ensure all prerequisites are installed.
2. Run setup script (shell step)
   ```sh
   ./setup.sh
   ```
3. Run something else (manual step)
   ...

## Deploy Application Runbook

This is a runbook for deploying the application.

1. Build the application (shell step)
   ```sh
   ./build.sh
   ```
2. Deploy the application (shell step)
   ```sh
    ./deploy.sh
    ```
3. Verify the deployment (manual step)
    ...
~~~

### Commands

Commands are simple standalone commands that can be executed. Similar to makefile targets.
They can be of different types: manual, shell, and etc.

Markdown command types are separated by code block syntax highlighting.
If a code block has `sh` syntax highlighting, it is considered a shell command.
Otherwise, it is considered a manual command.

#### Manual Commands

Manual commands are commands that require manual execution.
Usually, they contain instructions on how to perform the task manually.
Useful for documenting complex tasks before automating them.

#### Shell Commands

Shell commands are commands that can be executed automatically.
They contain the actual command to be executed.

### Runbooks (Workflows)

Runbook or Workflow is a collection of steps (commands) that has a specific execution order.

Executing a runbooks means executing all steps in the order they are listed.

## Library Discovery

AutoPilot will discover the library items from the current directory and its subdirectories.
It will look for Markdown files and parse them to extract commands and runbooks.

By default, AutoPilot will look for `LIBRARY.md` file in the current directory and parent directories recursively.
Also, by default it will look for `~/LIBRARY.md` file in the user's home directory to provide global commands and runbooks.
It can be configured to look for a different file or files.

## Configuration

AutoPilot can be configured via a configuration file and environment variables.

System-wide configuration file is located at `/etc/autopilot/config.yaml`.
User-specific configuration file is located at `~/.autopilot/config.yaml`. It can be used to override system-wide settings.

Configuration file is in YAML format. See the default configuration file here: [config.yaml](../pkg/config/config.yaml).

## Command Execution

AutoPilot executes commands in the current shell environment by default.
It can be configured to execute commands in a separate (clean) shell environment or container.

AutoPilot supports environment variable substitution before executing commands.
It can substitute environment variables in command strings using the format `${VAR}`.

## Runbook Execution

AutoPilot executes runbooks by executing each step in the order they are listed.

After step execution, AutoPilot saves environment variables (context) set by the command to be used in subsequent steps.

## Command/Runbook Discovery

AutoPilot provides a command/runbook discovery feature (fuzzy find) to search for items in the library.
It uses `fzf` to provide a TUI for item discovery.
