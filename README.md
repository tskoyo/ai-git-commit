# AI Git Commit

A CLI tool written in Go that uses OpenAI's GPT-4 to generate meaningful git commit messages based on staged changes. This tool streamlines your development workflow by suggesting commit messages that accurately describe code modifications.

## Table of Contents
- [Features](#features)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [Environment Variables](#environment-variables)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Features
- Automatically generate commit messages using OpenAI's GPT-4 model.
- Reads code changes and provides concise, descriptive commit messages.
- Skips integration tests in CI/CD environments like GitHub Actions.
- Configurable through a YAML file and environment variables.

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/ai-git-commit.git
2. Change to the project directory
   ```bash
   cd ai-git-commit
3. Install the tool
   ```bash
   ./install.sh

## Configuration

1. You can set environment variables.
   ```bash
   echo API_KEY="your_api_key"
