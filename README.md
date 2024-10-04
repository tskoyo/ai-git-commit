# AI Git Commit

A CLI tool written in Go that uses OpenAI's GPT-4 to generate meaningful git commit messages based on staged changes. This tool streamlines your development workflow by suggesting commit messages that accurately describe code modifications.

## Table of Contents
- [Features](#features)
- [Installation](#installation)
- [Configuration](#configuration)
- [Requirements](#requirements)

## Features
- Automatically generate commit messages using OpenAI's GPT-4 model.
- Reads code changes and provides concise, descriptive commit messages.
- Configurable through a YAML file and environment variables.

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/ai-git-commit.git
2. Change to the project directory
   ```bash
   cd ai-git-commit
3. Install dependencies
   ```bash
   go mod tidy
4. Install the app
  ```bash
  ./install.sh (chmod +x ./install.sh if necessary)

## Configuration

1. Create a config.yml file in the root of the project with the following structure:
   ```yml
   openai:
    api_key: "YOUR_OPENAI_API_KEY"
    model: "gpt-4"
    max_tokens: 100
2. Alternatively, you can set an environment variable API_KEY.
