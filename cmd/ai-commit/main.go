package main

import (
	"ai-git-commit/config"
	"ai-git-commit/internal/openai"
	"fmt"
)

func main() {
	content := `Generate a concise and meaningful commit message based on the following code changes:\n\n+ // Added JWT token validation for API endpoints\n+ func ValidateJWT(token string) bool {\n+     // Token validation logic\n+ }\n+ \n+ // Updated the login endpoint to return error messages for invalid credentials\n+ func Login(username, password string) error {\n+     if !ValidateJWT(token) {\n+         return fmt.Errorf(\"Invalid token\")\n+     }\n+     // Additional login logic\n+ }`

	body, err := openai.GetCommitMessage(config.ReadApiKey(), content)

	if err != nil {
		panic(err)
	}

	fmt.Println(body)
}
