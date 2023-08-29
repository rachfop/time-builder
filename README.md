
<div align="center">
<h1 align="center">
<img src="https://raw.githubusercontent.com/PKief/vscode-material-icon-theme/ec559a9f6bfd399b82bb44393651661b08aaf7ba/icons/folder-markdown-open.svg" width="100" />
<br>time-builder
</h1>
<h3>◦ Build time, build dreams!</h3>
<h3>◦ Developed with the software and tools listed below.</h3>

<p align="center">
<img src="https://img.shields.io/badge/Go-00ADD8.svg?style&logo=Go&logoColor=white" alt="Go" />
<img src="https://img.shields.io/badge/HCL-006BB6.svg?style&logo=HCL&logoColor=white" alt="HCL" />
<img src="https://img.shields.io/badge/Markdown-000000.svg?style&logo=Markdown&logoColor=white" alt="Markdown" />
</p>
<img src="https://img.shields.io/github/languages/top/rachfop/time-builder?style&color=5D6D7E" alt="GitHub top language" />
<img src="https://img.shields.io/github/languages/code-size/rachfop/time-builder?style&color=5D6D7E" alt="GitHub code size in bytes" />
<img src="https://img.shields.io/github/commit-activity/m/rachfop/time-builder?style&color=5D6D7E" alt="GitHub commit activity" />
<img src="https://img.shields.io/github/license/rachfop/time-builder?style&color=5D6D7E" alt="GitHub license" />
</div>

<div align="center">
  <img src="./static/demo.gif" alt="demo tape" />
</div>

---

## 📒 Table of Contents
- [📒 Table of Contents](#-table-of-contents)
- [📍 Overview](#-overview)
- [⚙️ Features](#️-features)
- [🚀 Getting Started](#-getting-started)
- [Installation](#installation)
  - [🎮 Using time-builder](#-using-time-builder)
  - [Using a Configuration File](#using-a-configuration-file)
- [📂 Project Structure](#-project-structure)
- [🧩 Modules](#-modules)
- [🤝 Contributing](#-contributing)
- [📄 License](#-license)

---


## 📍 Overview

The "time-builder" project is a command-line tool that generates Python code for Temporal Workflows. It allows users to provide input flags such as workflow name, child workflow name, and activity name, and generates corresponding Python code files based on these inputs. The generated code includes functionality to connect to a Temporal server, execute workflows, signal workflows with arguments, query workflows, and print the results. Overall, the project aims to simplify the process of generating Python code for Temporal Workflows, saving time and effort for developers.

---

## ⚙️ Features

| Feature                | Description                           |
| ---------------------- | ------------------------------------- |
| **⚙️ Architecture**     | The codebase follows a modular architecture where different components are organized into separate files. It uses a command-line interface (CLI) approach to generate Python code based on templates. The main package in `main.go` imports the `Execute` function from the `cmd` package. The `cmd` package contains two files: `generate.go` and `root.go`. The `generate.go` file implements the command-line tool that generates Python code based on various input flags and template files. The `root.go` file defines the root command using the Cobra library and sets up the flags for the `time-builder` command. Overall, the codebase follows a clean and organized architectural design.    |
| **📖 Documentation**   | The codebase contains documentation in the form of code comments, README files, and configuration files. The `README.template` file provides an overview and instructions for setting up the project. The Python template files also include comments explaining the purpose and usage of the code. While the documentation is present, it could be improved by adding more detailed explanations and examples to make it more comprehensive and user-friendly.   |
| **🔗 Dependencies**    | The codebase has external dependencies specified in the `go.mod` file. These dependencies include packages related to configuration parsing, file system operations, flag parsing, and environment variable handling. The `cmd` package imports the Cobra library for command-line interface functionality. The Python code generation templates rely on the Temporal.io workflow engine, which is being used to generate Python code for Temporal workflows. Overall, the codebase relies on these external libraries and systems to provide additional functionality and support for the required features.    |
| **🧩 Modularity**      | The codebase demonstrates good modularity by organizing different functionalities into separate files. Each template file in the `python_templates` directory serves a specific purpose, such as defining workflows, activities, and configuration files. The `cmd` package separates the command-line tool functionality into two files: `generate.go` and `root.go`. This modular structure allows for easy understanding, maintenance, and extension of the codebase.   |
| **✔️ Testing**          | The codebase does not include explicit testing files or tests as part of the repository. This could be a potential area for improvement, as proper testing is crucial to ensure the reliability and stability of the code. Implementing testing strategies and tools such as unit tests, integration tests, and end-to-end tests would greatly enhance the overall quality of the codebase.   |
| **⚡️ Performance**      | The performance of the codebase is subjective and dependent on the usage of the generated Python code rather than the codebase itself. However, the codebase itself does not contain any obvious performance bottlenecks or inefficient operations. It leverages the Temporal.io

---
## 🚀 Getting Started

## Installation

Clone the time-builder repository:

```bash
git clone https://github.com/rachfop/time-builder
```

Change to the project directory:

```bash
cd time-builder`
```

---


To install the CLI tool, you can build it from source:

```bash
go build -o time-builder
```

### 🎮 Using time-builder

Generate a basic Python project for Temporal.

```bash
./time-builder generate
```

Generates a Temporal Python project which connects to the Cloud.

```bash
./time-builder generate --cloud
```

Generates a Temporal Python project with 3 Signals and 2 Queries.

```bash
./time-builder generate --signal-count 3 --query-count 2
```

Generates a Temporal Python project with a child workflow with 2 Signals and 4 Queries.

```bash
./time-builder generate --child-workflow ChildWorkflowName  --signal Hello --signal-count 2 --query Bye --query-count 4
```

Generates a Temporal Python project with a child workflow and 3 Signals and 2 Queries.

```bash
./time-builder generate --child-workflow ChildWorkflowName --signal-count 3 --query-count 2
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
./time-builder generate --config ./config.yaml
```

---

## 📂 Project Structure


```bash
repo
├── LICENSE
├── README.md
├── cmd
│   ├── generate.go
│   └── root.go
├── config.yaml
├── go.mod
├── go.sum
├── main.go
└── python_templates
    ├── README.template
    ├── __init__.template
    ├── activities.template
    ├── pyproject.template
    ├── run_worker.template
    ├── run_workflow.template
    └── workflows.template

3 directories, 15 files
```

---

## 🧩 Modules

<details closed><summary>Root</summary>

| File                                                                 | Summary                                                                                                                                                                                                                                                                        |
| ---                                                                  | ---                                                                                                                                                                                                                                                                            |
| [go.mod](https://github.com/rachfop/time-builder/blob/main/go.mod)   | The code snippet is a module dependency file that specifies the required external packages for the "time-builder" module. It includes various packages related to configuration parsing and handling, file system operations, flag parsing, and environment variable handling. |
| [main.go](https://github.com/rachfop/time-builder/blob/main/main.go) | The code snippet is a main package that imports a command package from the "github.com/rachfop/time-builder/cmd" module. It runs the Execute function from the command package when the main function is called.                                                               |

</details>

<details closed><summary>Cmd</summary>

| File                                                                             | Summary                                                                                                                                                                                                                                                                                                                                                                               |
| ---                                                                              | ---                                                                                                                                                                                                                                                                                                                                                                                   |
| [generate.go](https://github.com/rachfop/time-builder/blob/main/cmd/generate.go) | This code snippet is a command-line tool that generates Python code based on templates. It takes various input flags such as workflow name, child workflow name, local activity name, signal name, query name, signal count, query count, useCloud flag, and config path. It reads template files, populates them with the input data, and generates corresponding Python code files. |
| [root.go](https://github.com/rachfop/time-builder/blob/main/cmd/root.go)         | The provided code snippet is a CLI tool that generates Python code for Temporal Workflows. It uses the Cobra library to define the root command and add flags. The "time-builder" command can be executed to generate the Python code.                                                                                                                                                |

</details>

<details closed><summary>Python_templates</summary>

| File                                                                                                              | Summary                                                                                                                                                                                                                                                                                      |
| ---                                                                                                               | ---                                                                                                                                                                                                                                                                                          |
| [run_workflow.template](https://github.com/rachfop/time-builder/blob/main/python_templates/run_workflow.template) | The provided code snippet connects to a temporal service, executes a workflow (with optional child workflow), signals the workflow with arguments, queries the workflow, and prints the result.                                                                                              |
| [run_worker.template](https://github.com/rachfop/time-builder/blob/main/python_templates/run_worker.template)     | This code snippet establishes a connection to a temporal server and starts a worker to execute workflows and activities. It supports cloud and local deployments, and can include child workflows if specified.                                                                              |
| [pyproject.template](https://github.com/rachfop/time-builder/blob/main/python_templates/pyproject.template)       | This code snippet is a configuration file for a Python SDK generator for the Temporal.io workflow engine. It specifies the project's dependencies, development tools, and tasks such as formatting, linting, and testing.                                                                    |
| [workflows.template](https://github.com/rachfop/time-builder/blob/main/python_templates/workflows.template)       | This code snippet defines a workflow that can either execute a child workflow, a local activity, or a regular activity using the specified parameters. It also includes signal and query methods for the workflow. The child workflow is defined separately and executes a regular activity. |
| [activities.template](https://github.com/rachfop/time-builder/blob/main/python_templates/activities.template)     | The code snippet defines a dataclass called YourParams with two attributes: greeting and name. It also defines an activity function called say_hello that takes an instance of YourParams as input and returns a string by concatenating the greeting and name attributes.                   |
| [README.template](https://github.com/rachfop/time-builder/blob/main/python_templates/README.template)             | The provided code snippet sets up a Temporal project. It starts the Temporal server, runs a worker, and executes a workflow.                                                                                                                                                                 |
| [__init__.template](https://github.com/rachfop/time-builder/blob/main/python_templates/__init__.template)         | The code snippet provides a function that calculates the average value of a list of numbers.                                                                                                                                                                                                 |
| [.env.template](https://github.com/rachfop/time-builder/blob/main/python_templates/.env.template)                 | The code snippet sets configuration variables for connecting to Temporal server using mutual TLS authentication. It specifies the path to TLS certificate and key files, Temporal server URL, and the default namespace.                                                                     |

</details>

---

---



---

## 🤝 Contributing

Contributions are always welcome! Please follow these steps:
1. Fork the project repository. This creates a copy of the project on your account that you can modify without affecting the original project.
2. Clone the forked repository to your local machine using a Git client like Git or GitHub Desktop.
3. Create a new branch with a descriptive name (e.g., `new-feature-branch` or `bugfix-issue-123`).
```sh
git checkout -b new-feature-branch
```
4. Make changes to the project's codebase.
5. Commit your changes to your local branch with a clear commit message that explains the changes you've made.
```sh
git commit -m 'Implemented new feature.'
```
6. Push your changes to your forked repository on GitHub using the following command
```sh
git push origin new-feature-branch
```
7. Create a new pull request to the original project repository. In the pull request, describe the changes you've made and why they're necessary.
The project maintainers will review your changes and provide feedback or merge them into the main branch.

---

## 📄 License

This project is licensed under the `ℹ️  MIT` License.
See the [LICENSE](./LICENSE) file for additional info.