# Contributing to go-cache

Thank you for your interest in contributing to go-cache! We appreciate all contributions whether it’s new features or bug fixes. This guide outlines the steps and standards to follow when contributing.

## Table of Contents

1. [Getting Started](#getting-started)
2. [How to Contribute](#how-to-contribute)
   - [Fork the Repository](#fork-the-repository)
   - [Clone Your Fork](#clone-your-fork)
   - [Create a Feature Branch](#create-a-feature-branch)
   - [Write Unit Tests](#write-unit-tests)
   - [Use Makefile Commands](#use-makefile-commands)
3. [Pull Request Process](#pull-request-process)
4. [License](#license)

## Getting Started

Before you begin contributing, make sure you have the following tools installed:

- **Git**
- **Go**: Ensure you are using the correct version of Go specified in the project’s `go.mod` file.
- **Make**

Once your development environment is set up, follow the steps below to contribute.

## How to Contribute

### Fork the Repository

Start by forking the go-cache repository to your own GitHub account.

### Clone Your Fork

Next, clone the forked repository to your local machine.

### Create a Feature Branch

Before making any changes, create a new branch for your feature or bug fix. Branches should be named according to the feature you're working on.

- **feature/my-awesome-feature**
- **bugfix/bug-name**

### Write Unit Tests

Ensure that every new feature or bug fix is accompanied by relevant unit tests. Tests should be placed in the appropriate test files.

### Use Makefile Commands

This repository includes a Makefile to help you streamline development tasks.

- Running tests:

```bash
    make test
```

- Running quality control checks:

```bash
    make audit
```

- Format code and tidy modfile:

```bash
    make tidy
```

### Pull Request Process

1. **Create a pull request**.
2. **Review process**: Your pull request will be reviewed by the maintainer. Please be responsive to any feedback.
3. **Approval**: Once all checks (CI workflow) are passed and the pull request is approved by the maintainer, it will be merged into the main branch.
4. **Celebrate**: 🎉 Thank you for contributing to go-cache!

### License

By contributing, you agree that your contributions will be licensed under the [MIT License](LICENSE).
