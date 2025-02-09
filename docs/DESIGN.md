# AutoPilot Design (WIP)

AutoPilot is a tool to document and automate daily commands and workflows. It is designed to be simple, flexible, extensible, and easy to share. You can define commands and runbooks in simple Markdown format and commit them to your project repository.

Runbook format is clean and simple Markdown document, so you can use it independently without AutoPilot or any other tool. AutoPilot provides additional features like parsing document and extracting commands, command discovery, environment variable substitution, and command execution. You don't need to learn a new DSL or use a complex infrastructure to start using AutoPilot.

## Idea

The core idea behind AutoPilot is to provide a simplest possible framework for gradual automation. It allows to start with simple manual workflows and gradually automate them as needed. AutoPilot is designed to be flexible and extensible, so you can define workflows in a way that makes sense for your specific use case. You can start with a simple manual workflow and incrementally automate the parts that make the most sense for your specific needs.

## Core Concepts

### Commands (Instructions)

Commands are simple standalone scripts or instructions (manual or shell). Runbooks group these commands to form workflows with a specified execution order. Steps can be manual or automated, depending on your needs.

> AutoPilot supports only Markdown format for commands and libraries at this time.
> We might add other formats (e.g., YAML/JSON) in the future.

* Manual Commands

  Manual commands are commands that require manual execution.
  Usually, they contain instructions on how to perform the task manually.
  Useful for documenting complex tasks before automating them.

* Shell Commands

  Shell commands are commands that can be executed automatically.
  They contain the actual command to be executed.

> Markdown command types are separated by code block syntax highlighting.
> If a code block has `sh` syntax highlighting, it is considered a shell command.
> Otherwise, it is considered a manual command.

### Runbooks (Workflows)

A runbook is a collection of steps (commands) that has a specific execution order.

* Manual Steps – Require user action to proceed.

    For example:
    * Manual Commands – Instructions that need to be executed manually.
    * Approval Steps – Require explicit user approval before proceeding.

* Automated Steps – Shell scripts, API calls, or external integrations that can be executed automatically.

    For example:
    * Shell Steps – Executed as shell commands.
    * API Steps – Executed as API calls.
    * Conditional Steps – Executed based on logic (e.g., if a previous step fails).

Executing a runbook means executing all steps in the order they are listed.

### Library

Library is a collection of commands and runbooks (workflows) that can be stored in various formats:

* Markdown (default) – Human-readable, simple command lists. Might be used independently.
* YAML/JSON (future?) – Structured formats for more complex automation.
* Custom DSL (future?) – A more expressive way to define workflows.

> AutoPilot initially will support only Markdown format for libraries, because it is human-readable and simple. It will lack some complex features like loops, conditionals, etc. We might add support for other formats in the future.

### Execution State & Tracking

Execution state is a way to track the progress of commands and workflows. It can be used to resume workflows after system restarts or failures. Also it allows to track the progress of long-running workflows or running multiple workflows in parallel. Another use case is ability to transfer the runbook execution to another machine or user.

Execution state is stored locally on the file system or in a database (future?).

AutoPilot tracks the execution state of commands and workflows to provide the following features:

* Resume Execution: Track progress and resume workflows after interruptions.
* Interactive Execution: Pause, modify (future?), skip, or restart steps dynamically.
* Execution History: Maintain logs of previous runs for debugging and auditing.

### Context Management

Context management is a way to store and retrieve variables during execution. It can be used to pass data between steps, share data between sub-workflows, or store intermediate results. Context management is essential for complex workflows that require data transformation, aggregation, or conditional logic.

AutoPilot supports context management by allowing to set and get environment variables during execution. Then it can use variable substitution in commands and workflows to pass data between steps. It can substitute environment variables in command strings using the format `${VAR}`. In the future, it might support more advanced features like artifacts, secrets, and other data types.

### Execution Environment

Execution environment is a way to control where and how commands and workflows are executed. It can be used to run commands in a clean environment, container, or remote machine. Execution environment is essential for ensuring consistency, security, and isolation during execution.

AutoPilot supports the following execution environments:
* Local Shell – Execute commands in the current shell environment.
* Clean Shell – Execute commands in a clean shell environment.
* Container (future?) – Execute commands in a containerized environment.
* Remote Machine (future?) – Execute commands on a remote machine.

### Library & Command Discovery & Search

Is a way to discover libraries and commands in the library. It can be used to search for commands, workflows, or documentation. Discovery feature is essential for finding and executing commands quickly without remembering the exact name or syntax.

AutoPilot provides the following discovery features:
* Library Discovery – Automatically discover libraries on the file system, database (future?), or remote repository (future?).
* Command Fuzzy Search – Search for commands and workflows in the library using fuzzy matching or other search algorithms.
* Command Ranking (future?) – Rank commands based on usage, popularity, or other metrics to improve search results.

AutoPilot discovers the library items from the current directory and its parent directories recursively.
It will look for Markdown files and parse them to extract commands and runbooks.

By default, AutoPilot will look for `LIBRARY.md` file in the current directory and parent directories recursively.
Also, by default it will look for `~/LIBRARY.md` file in the user's home directory to provide global commands and runbooks.
It can be configured to look for a different file or files.

For command discovery, AutoPilot uses `fzf` to provide a TUI for item discovery. It allows to search for items in the library and execute them interactively.

## Design details

### Markdown Library

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

### Configuration

AutoPilot can be configured via a configuration file and environment variables.

System-wide configuration file is located at `/etc/autopilot/config.yaml`.
User-specific configuration file is located at `~/.autopilot/config.yaml`. It can be used to override system-wide settings.

Configuration file is in YAML format. See the default configuration file here: [config.yaml](../pkg/config/config.yaml).
