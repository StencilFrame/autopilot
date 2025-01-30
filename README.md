# AutoPilot - A system for gradual runbook automation

AutoPilot is a tool designed to gradually automate repetitive tasks by defining them in a runbook. Unlike traditional automation tools that require complete automation from the outset, AutoPilot enables a step-by-step approach, making it easier to transition from manual to automated processes without overwhelming your team or infrastructure.

Build what matters!

Inspired by [Do Nothing Scripting](https://blog.danslimmon.com/2019/07/15/do-nothing-scripting-the-key-to-gradual-automation) by Dan Slimmon.

If you want to read more about the design and idea behind AutoPilot, check out the [idea document](docs/IDEA.md).

![demo](./examples/demo.gif)

## Purpose and Real-Life Examples

### Purpose

AutoPilot is built to address the challenges of full-scale automation, which can be overwhelming and rigid. By enabling gradual automation, AutoPilot allows teams to:

- **Start Small:** Begin with manual workflows and identify the most impactful steps to automate.
- **Maintain Control:** Automate processes at your own pace, ensuring reliability and reducing errors.
- **Enhance Flexibility:** Adapt and evolve your workflows as your needs change, without being locked into a rigid automation structure.

### Real-Life Examples

1. System Maintenance:
  - **Manual Workflow:**  
    Every morning, an engineer manually checks system health, rotates logs, and applies security patches.
  - **Runbook steps:**
    - Check system health
    - Rotate logs
    - Apply security patches
  - **Gradual Automation:**  
    You start by automating the most critical step for your application. For example, you automate the log rotation process, because if not done correctly, it can lead to disk space issues and cause the application to crash. Next time you run the runbook, you can skip the log rotation step as it's already automated. The runbook will tell you what to do next, so you don't have to remember.

    Having a runbook allows you to share the process with other team members and ensure consistency across the team.

2. DevOps Pipelines:
  - **Manual Workflow:**  
    CI/CD pipeline is not ready yet, so developers manually build and deploy the application
  - **Runbook steps:**
    - Run tests
    - Build the application
    - Deploy to staging
    - Check the deployment
  - **Gradual Automation:**  
    You decide that the most time-consuming and error-prone step is the deployment to staging. You start by automating this step, so developers can focus on writing code and running tests. Once the deployment step is automated, you move on to the other steps in the pipeline.

    Having a runbook makes sure that everyone follows the same steps and that the process is consistent.

3. Developer Experience (New User Onboarding):
  - **Manual Workflow:**  
    New user onboarding involves setting up local development environments, running build scripts, and deploying to staging.
  - **Runbook steps:**
    - Instructions for seting up local development environment
    - Instructions for running build scripts
    - Instructions for deploying to staging
  - **Gradual Automation:**
    You start by automating the local development environment setup or sub-step, as it's a time-consuming and error-prone process. Once that's automated, you move on to the build scripts and deployment steps. This way, new users can get started quickly, and you can ensure consistency across the team.

    Having a runbook ensure that they follow the same steps every time and that the process is consistent.

## Supported features

- **Define Runbooks in Markdown or YAML:** Choose the format that best fits your workflow documentation needs.
- **Supports Manual and Shell Steps:** Start with fully manual workflows and incrementally automate specific steps.
- **Gradual Automation:** Automate one step at a time, ensuring each part works perfectly before moving on.
- **User-Friendly CLI:** Intuitive command-line interface for executing runbooks
- **Cross-Platform Support:** Compatibile with Windows, macOS, and Linux

## Roadmap

This is an early MVP version of AutoPilot. The following features are planned for future releases:

- **Execution Tracking:** Keep track of runbook executions to resume from the last step
- **Context Management:** Store and retrieve variables during execution
- **Comprehensive Logging:** Detailed logs of step executions for auditing and troubleshooting
- **Advanced Step Types:** Input, conditional, nested steps, and more
- **Runbook Type Overrides:** Explicitly specify runbook formats (e.g., --type=markdown or --type=yaml)
- **Enhanced Executors:** Support Docker, API calls, and other complex execution environments
- **User-Friendly CLI:** Intuitive command-line interface for managing runbooks
- **Web UI:** Visual dashboard for managing runbooks and monitoring executions
- **Plugin and SDK Support:** Extend functionality with plugins and provide SDKs for developers
- **Notifications and Scheduling:** Support notification API and allow runbook scheduling
- **Distribution Support:** Provide installation packages for different distributions (e.g., ~~Homebrew~~, APT, etc.)
- **Security and Access Control:** Role-based access control and secrets management
- **Extensible Architecture:** Easily add new step types and integrations through plugins and APIs
- **Plugin Marketplace:** Central repository for sharing and discovering plugins and automation scripts/steps.

If you have any feature requests or suggestions, please open an issue on GitHub. Pull requests are also welcome!

If you want us to prioritize a feature, please thumbs up the issue or comment on it.

## Installation

### From Source

To install AutoPilot, clone the repository and build the project:

```sh
git clone https://github.com/stencilframe/autopilot.git
cd autopilot
go build -o autopilot ./pkg/cmd/autopilot
```

### Homebrew (macOS)

```sh
brew tap stencilframe/autopilot
brew install autopilot
```

or just

```sh
brew install stencilframe/autopilot/autopilot
```

### Download Binary from Releases

Download the latest release from the [releases page](https://github.com/StencilFrame/autopilot/releases)

## Usage

### Define a Runbook

A runbook is a series of steps that can be executed manually or automatically. This tool supports defining runbooks in both simple Markdown and extensible YAML formats.

#### Markdown Runbooks

Markdown runbooks are defined by listing ordered steps with instructions and code blocks for shell steps.
AutoPilot extracts the first ordered list in the file as the runbook steps. If there are multiple ordered lists, only the first one is considered.

AutoPilot supports and detects two types of steps: manual and shell (automatic).
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

To execute a runbook, use the `autopilot` command followed by the runbook file:

```sh
./autopilot run runbook.md
```

Runbook types are automatically detected based on the file extension:
- `.md` files are considered Markdown runbooks.
- `.yml` or `.yaml` files are considered YAML runbooks.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request on GitHub.

## Support

If you have any questions or need help, please open an issue on GitHub, start a [discussion](https://github.com/StencilFrame/autopilot/discussions) or find us on [Slack](https://join.slack.com/t/stencilframesupport/shared_invite/zt-2ynp05the-4~kanvoSa~HTHxZCUDuKEg)

## Sponsoring

If you find this project helpful, please consider sponsoring it. Your support helps us maintain and improve the project.
