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

func ReadApiKey() string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(filepath.Join(filepath.Dir(cwd), "..", "config.yml"))
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}
	defer file.Close()

	var cfg Config

	decoder := yaml.NewDecoder(file)
	if err != decoder.Decode(&cfg) {
		log.Fatalf("Error decoding YAML: %v", err)
	}

	return cfg.OpenAI.APIKey
}
