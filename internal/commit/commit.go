package commit

import (
	"ai-git-commit/config"
	"ai-git-commit/internal/git"
	"ai-git-commit/internal/openai"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func GenerateMessage() {
	apiKey, err := config.ReadApiKey()
	if err != nil {
		log.Fatalf("Error reading API key: %v", err)
	}

	for {
		commitMessage := gitDiff(apiKey)

		// Display the suggested commit message
		fmt.Printf("Suggested Commit Message: %s\n", commitMessage)

		// Ask user for confirmation
		fmt.Print("Are you satisfied with this commit message? (yes/no): ")
		reader := bufio.NewReader(os.Stdin)
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(strings.ToLower(userInput))

		if userInput == "yes" {
			if err := git.CommitMessage(commitMessage); err != nil {
				log.Fatalf("Error executing git commit: %v", err)
			}
			fmt.Println("Commit successfully created!")
			break
		} else {
			fmt.Println("Generating a new commit message...")
		}
	}
}

func gitDiff(apiKey string) string {
	err := git.IsGitRepo()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	diff, err := git.GetStagedDiff()
	if err != nil {
		log.Fatalf("Error getting staged diff: %v", err)
	}

	formattedDiff := git.FormatDiff(diff)

	fmt.Printf("Formatted diff: %v\n", formattedDiff)

	commitMessage, err := openai.GetCommitMessage(apiKey, "Generate a concise commit message for the following changes: "+formattedDiff)
	if err != nil {
		log.Fatalf("Error generating commit message: %v", err)
	}

	return commitMessage
}
