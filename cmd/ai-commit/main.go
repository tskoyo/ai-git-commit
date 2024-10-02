package main

import (
	"ai-git-commit/internal/commit"
	"os"
)

func main() {
	apiKey, ok := os.LookupEnv("API_KEY")
	if !ok {
		panic("API_KEY not found")
	} else {
		commit.GenerateMessage(apiKey)
	}
}
