#!/bin/bash

BINARY_NAME="ai-git-commit"

echo "Building the Go program..."
go build -o $BINARY_NAME main.go

if [ $? -ne 0 ]; then
    echo "Build failed. Please make sure Go is installed and the code compiles successfully."
    exit 1
fi

echo "Installing the binary to /usr/local/bin..."
sudo mv $BINARY_NAME /usr/local/bin/

read -p "Do you want to set up the Git alias 'git ai-commit'? (y/n) " answer
if [ "$answer" != "${answer#[Yy]}" ]; then
    git config --global alias.ai-commit '!ai-git-commit'
    echo "Git alias 'ai-commit' has been set up."
fi

echo "Installation complete! You can now use 'ai-git-commit' or 'git ai-commit' (if you set up the alias)"
