# Temporal Python Project Generator

## Overview

This CLI tool generates a Temporal Python project with a predefined structure. It sets up workflows, activities, and necessary dependencies to help you get started quickly with Temporal.

## Features

- Generates a Python project structure for Temporal.
- Allows customization of workflow names and other settings.
- Supports the use of a configuration file for easier and repetitive generation.

## Installation

To install the CLI tool, you can build it from source:

```bash
go build -o temporal-py
```

## Usage

### Basic Command

Generate a basic Python project for Temporal.

```bash
go run main.go generate
```

Generates a Temporal Python project which connects to the Cloud.

```bash
go run main.go generate --cloud
```

Generates a Temporal Python project with 3 Signals and 2 Queries.

```bash
go run main.go generate --signal-count 3 --query-count 2
```

Generates a Temporal Python project with a child workflow with 2 Signals and 4 Queries.


```bash
go run main.go generate --child-workflow ChildWorkflowName  --signal Hello --signal-count 2 --query Bye --query-count 4
```

Generates a Temporal Python project with a child workflow and 3 Signals and 2 Queries.

```bash
go run main.go generate --child-workflow ChildWorkflowName --signal-count 3 --query-count 2
```


### Using a Configuration File

To make the generation process easier, you can use a configuration file named `config.yaml` placed in the same directory as the CLI tool. A sample configuration file (`config.yaml.sample`) is included in the project.

Here is what a typical `config.yaml` could look like:

```yaml
# Temporal Python Project Configuration

# Name of the workflow
workflowName: "MyWorkflow"

# Name of the child workflow
childWorkflowName: "MyChildWorkflow"

# Whether to use cloud (true or false)
useCloud: false
```

If a configuration file is present and the `--config` flag is set, the CLI tool will use the settings from the file, unless overridden by command-line flags.

```bash
go run main.go generate --config ./config.yaml
```

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.
