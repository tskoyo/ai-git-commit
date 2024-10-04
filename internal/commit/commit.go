package commit

import (
	"ai-git-commit/internal/git"
	"ai-git-commit/internal/openai"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func GenerateMessage(apiKey string) {
	for {
		commitMessage := gitDiff(apiKey)

		fmt.Printf("Suggested Commit Message: %s\n", commitMessage)

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

	commitMessageGenerator := &openai.OpenAIGenerator{}
	commitMessage, err := commitMessageGenerator.GetCommitMessage(apiKey, "Generate a concise commit message for the following changes: "+formattedDiff)
	if err != nil {
		log.Fatalf("Error generating commit message: %v", err)
	}

	return commitMessage
}
