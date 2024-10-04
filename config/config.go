package config

import (
	"log"
	"os"

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

func OpenConfigFile(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
		return nil, err
	}
	return file, nil
}

func DecodeConfigFile(file *os.File) (Config, error) {
	var cfg Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		log.Fatalf("Error decoding YAML: %v", err)
		return cfg, err
	}
	return cfg, nil
}

func GetAPIKey(cfg Config) (string, error) {
	if cfg.OpenAI.APIKey != "" {
		return cfg.OpenAI.APIKey, nil
	}

	apiKey, ok := os.LookupEnv("API_KEY")
	if !ok {
		panic("API_KEY not found")
	}
	return apiKey, nil
}

func ReadAPIKey() (string, error) {
	file, err := OpenConfigFile("config.yml")
	if err != nil {
		return "", err
	}
	defer file.Close()

	cfg, err := DecodeConfigFile(file)
	if err != nil {
		return "", err
	}

	return GetAPIKey(cfg)
}
