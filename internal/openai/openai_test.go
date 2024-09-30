package openai

import (
	"ai-git-commit/config"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCommitMessage_EmptyAPIKey(t *testing.T) {
	content := `Generate a concise and meaningful commit message based on the following code changes:\n\n+ // Added JWT token validation for API endpoints\n+ func ValidateJWT(token string) bool {\n+     // Token validation logic\n+ }\n+ \n+ // Updated the login endpoint to return error messages for invalid credentials\n+ func Login(username, password string) error {\n+     if !ValidateJWT(token) {\n+         return fmt.Errorf(\"Invalid token\")\n+     }\n+     // Additional login logic\n+ }`
	_, err := GetCommitMessage("", content)

	assert.Error(t, err, "missing API key")
}

func TestGetCommitMessage_EmptyRequestBody(t *testing.T) {
	apiKey, err := config.ReadApiKey()
	if err != nil {
		panic(err)
	}
	_, err = GetCommitMessage(apiKey, "")

	assert.Error(t, err, "empty prompt")
}

func TestGetCommitMessage_Success(t *testing.T) {
	apiKey, err := config.ReadApiKey()
	if err != nil {
		panic(err)
	}

	if os.Getenv("GITHUB_ACTIONS") == "true" {
		t.Skip("Skipping integration test in GitHub Actions environment")
	}

	if apiKey == "" {
		t.Skip("Skipping integration test: OPENAI_API_KEY is not set")
	}

	content := `+ // Added JWT token validation for API endpoints...`

	message, err := GetCommitMessage(apiKey, content)

	assert.NoError(t, err, "Expected no error from the API call")
	assert.NotEmpty(t, message, "Expected a commit message from the API call")
}
