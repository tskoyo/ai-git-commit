#!/bin/bash

BINARY_NAME="ai-git-commit"
INSTALL_PATH="/usr/local/bin/$BINARY_NAME"

echo "Building the Go program..."
go build -o $BINARY_NAME main.go

if [ $? -ne 0 ]; then
    echo "Build failed. Please make sure Go is installed and the code compiles successfully."
    exit 1
fi

if [ -f "$INSTALL_PATH" ]; then
    echo "$BINARY_NAME is already installed."

    if cmp -s "$BINARY_NAME" "$INSTALL_PATH"; then
        echo "The existing binary is up-to-date. No update necessary."
        rm $BINARY_NAME
        exit 0
    else
        echo "Updating the installed binary..."
    fi
else
    echo "No existing installation found. Installing the binary..."
fi

sudo mv $BINARY_NAME "$INSTALL_PATH"

read -p "Do you want to set up the Git alias 'git ai-commit'? (y/n) " answer
if [ "$answer" != "${answer#[Yy]}" ]; then
    git config --global alias.ai-commit '!ai-git-commit'
    echo "Git alias 'ai-commit' has been set up."
fi

echo "Installation complete! You can now use 'ai-git-commit' or 'git ai-commit' (if you set up the alias)"
