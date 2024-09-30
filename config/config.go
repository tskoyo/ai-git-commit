package config

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type OpenAIConfig struct {
	APIKey    string `yaml:"api_key"`
	Model     string `yaml:"model"`
	MaxTokens int    `yaml:"max_tokens"`
}

type Config struct {
	OpenAI OpenAIConfig `yaml:"openai"`
}

func ReadApiKey() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	projectRoot := filepath.Dir(filepath.Dir(cwd))

	configPath := filepath.Join(projectRoot, "config.yml")

	file, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
		return "", err
	}
	defer file.Close()

	var cfg Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		log.Fatalf("Error decoding YAML: %v", err)
		return "", err
	}

	return cfg.OpenAI.APIKey, nil
}
