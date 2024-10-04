package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type APIKeyReader interface {
	ReadAPIKey() (string, error)
}

type FileConfigReader struct {
	Filename string
}

type OpenAIConfig struct {
	APIKey    string `yaml:"api_key"`
	Model     string `yaml:"model"`
	MaxTokens int    `yaml:"max_tokens"`
}

type Config struct {
	OpenAI OpenAIConfig `yaml:"openai"`
}

func (r *FileConfigReader) openConfigFile() (*os.File, error) {
	file, err := os.Open(r.Filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (r *FileConfigReader) decodeConfigFile(file *os.File) (Config, error) {
	var cfg Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}

func (r *FileConfigReader) getAPIKey(cfg Config) (string, error) {
	if cfg.OpenAI.APIKey != "" {
		return cfg.OpenAI.APIKey, nil
	}

	apiKey, ok := os.LookupEnv("API_KEY")
	if !ok {
		return "", os.ErrNotExist
	}
	return apiKey, nil
}

func (r *FileConfigReader) ReadAPIKey() (string, error) {
	file, err := r.openConfigFile()
	if err != nil {
		return "", err
	}
	defer file.Close()

	cfg, err := r.decodeConfigFile(file)
	if err != nil {
		return "", err
	}

	return r.getAPIKey(cfg)
}
