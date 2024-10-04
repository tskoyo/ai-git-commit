package openai

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Mock for APIKeyReader
type MockAPIKeyReader struct {
	APIKey string
	Err    error
}

func (m *MockAPIKeyReader) ReadAPIKey() (string, error) {
	return m.APIKey, m.Err
}

// Mock for CommitMessageGenerator
type MockCommitMessageGenerator struct {
	Message string
	Err     error
}

func (m *MockCommitMessageGenerator) GetCommitMessage(apiKey, content string) (string, error) {
	if apiKey == "" {
		return "", errors.New("missing API key")
	}
	if content == "" {
		return "", errors.New("empty prompt")
	}
	return m.Message, m.Err
}

func TestGetCommitMessage_EmptyAPIKey(t *testing.T) {
	mockGenerator := &MockCommitMessageGenerator{Err: errors.New("missing API key")}
	_, err := mockGenerator.GetCommitMessage("", "some content")

	assert.Error(t, err, "missing API key")
}

func TestGetCommitMessage_EmptyRequestBody(t *testing.T) {
	mockAPIKeyReader := &MockAPIKeyReader{APIKey: "mock-api-key"}
	mockGenerator := &MockCommitMessageGenerator{Err: errors.New("empty prompt")}

	apiKey, err := mockAPIKeyReader.ReadAPIKey()
	assert.NoError(t, err, "Expected no error when reading API key")

	_, err = mockGenerator.GetCommitMessage(apiKey, "")

	assert.Error(t, err, "empty prompt")
}

func TestGetCommitMessage_Success(t *testing.T) {
	mockAPIKeyReader := &MockAPIKeyReader{APIKey: "mock-api-key"}
	mockGenerator := &MockCommitMessageGenerator{Message: "Mock commit message"}

	apiKey, err := mockAPIKeyReader.ReadAPIKey()
	assert.NoError(t, err, "Expected no error when reading API key")

	content := `+ // Added JWT token validation for API endpoints...`
	message, err := mockGenerator.GetCommitMessage(apiKey, content)

	assert.NoError(t, err, "Expected no error from the mock API call")
	assert.Equal(t, "Mock commit message", message, "Expected the mock commit message")
}
