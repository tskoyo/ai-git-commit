package main

import (
	"ai-git-commit/config"
	"ai-git-commit/internal/commit"
)

func main() {
	configReader := &config.FileConfigReader{Filename: "config.yml"}

	apiKey, err := configReader.ReadAPIKey()
	if err != nil {
		panic(err)
	}

	commit.GenerateMessage(apiKey)
}
