package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	TestsPath string `json:"tests_path"`
}

func New() Config {
	return Config{}
}

func (c *Config) Load(path string) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(file, c)
}
