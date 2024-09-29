package openai

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

const apiURL = "https://api.openai.com/v1/chat/completions"

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChoiceMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Choice struct {
	Index        int           `json:"index"`
	Message      ChoiceMessage `json:"message"`
	Logprobs     interface{}   `json:"logprobs"`
	FinishReason string        `json:"finish_reason"`
}

type ResponseBody struct {
	Choices []Choice `json:"choices"`
}

type RequestBody struct {
	Model   string    `json:"model"`
	Message []Message `json:"messages"`
}

func GetCommitMessage(apiKey, prompt string) (string, error) {
	requestBody := RequestBody{
		Model:   "gpt-4o",
		Message: []Message{{Role: "user", Content: prompt}},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
