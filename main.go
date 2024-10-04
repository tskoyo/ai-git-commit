package main

import (
	"ai-git-commit/config"
	"ai-git-commit/internal/commit"
)

func main() {
	apiKey, err := config.ReadAPIKey()
	if err != nil {
		panic(err)
	}

	commit.GenerateMessage(apiKey)
}
