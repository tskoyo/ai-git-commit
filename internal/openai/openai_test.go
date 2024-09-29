package openai

import (
	"ai-git-commit/config"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCommitMessage_EmptyAPIKey(t *testing.T) {
	apiURL := "https://api.openai.com/v1/chat/completions"
	requestBody := RequestBody{
		Model:   "gpt-4o",
		Message: []Message{{Role: "user", Content: "Test prompt"}},
	}

	jsonData, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer server.Close()

	content := `Generate a concise and meaningful commit message based on the following code changes:\n\n+ // Added JWT token validation for API endpoints\n+ func ValidateJWT(token string) bool {\n+     // Token validation logic\n+ }\n+ \n+ // Updated the login endpoint to return error messages for invalid credentials\n+ func Login(username, password string) error {\n+     if !ValidateJWT(token) {\n+         return fmt.Errorf(\"Invalid token\")\n+     }\n+     // Additional login logic\n+ }`
	_, err := GetCommitMessage("", content)

	assert.Error(t, err, "missing API key")
}

func TestGetCommitMessage_EmptyRequestBody(t *testing.T) {
	apiURL := "https://api.openai.com/v1/chat/completions"
	requestBody := RequestBody{}

	jsonData, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	_, err := GetCommitMessage(config.ReadApiKey(), "")

	assert.Error(t, err, "empty prompt")
}

func TestGetCommitMessage_Integration(t *testing.T) {
	apiKey := config.ReadApiKey()
	if os.Getenv("GITHUB_ACTIONS") == "true" {
		t.Skip("Skipping integration test in GitHub Actions environment")
	}

	if apiKey == "" {
		t.Skip("Skipping integration test: OPENAI_API_KEY is not set")
	}

	content := `+ // Added JWT token validation for API endpoints...`

	message, err := GetCommitMessage(apiKey, content)

	assert.NoError(t, err, "Expected no error from the real API call")
	assert.NotEmpty(t, message, "Expected a commit message from the real API call")
}
