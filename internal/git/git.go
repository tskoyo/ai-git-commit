package git

import (
	"ai-git-commit/internal/openai"
	"bytes"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

var execCommand = exec.Command

func GitDiff(apiKey string) {
	err := isGitRepo()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	diff, err := getStagedDiff()
	if err != nil {
		log.Fatalf("Error getting staged diff: %v", err)
	}

	formattedDiff := formatDiff(diff)

	commitMessage, err := openai.GetCommitMessage(apiKey, "Generate a concise commit message for the following changes: "+formattedDiff)
	if err != nil {
		log.Fatalf("Error generating commit message: %v", err)
	}

	fmt.Printf("Suggested Commit Message: %s\n", commitMessage)
}

func isGitRepo() error {
	cmd := execCommand("git", "rev-parse", "--is-inside-work-tree")
	if err := cmd.Run(); err != nil {
		return errors.New("current directory is not a git repository")
	}
	return nil
}

func getStagedDiff() (string, error) {
	cmd := execCommand("git", "diff", "--cached")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "", err
	}
	return out.String(), nil
}

func formatDiff(diff string) string {
	lines := strings.Split(diff, "\n")
	for _, line := range lines {
		fmt.Printf("line %s\n", line)
	}

	var formattedLines []string

	for _, line := range lines {
		if line == "" || strings.HasPrefix(line, "diff") || strings.HasPrefix(line, "index") {
			continue
		}
		formattedLines = append(formattedLines, strings.TrimSpace(line))
	}

	return strings.Join(formattedLines, " ")
}
