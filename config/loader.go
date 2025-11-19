package config

import (
	"encoding/json"
	"os"
)

func LoadConfig(filename string) (*AppConfig, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config AppConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}