package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	path string
	Name string `json:"name"`
}

func Load() (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory:", err)
		return nil, err
	}

	path := filepath.Join(homeDir, ".config", "irkit", "config.toml")

	// Check if config file exists, returning an empty config if the file cannot be found
	_, err = os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		// Config doesn't exist
		return &Config{path: path}, nil
	} else if err != nil {
		// Shouldn't occur
		return nil, err
	}

	// Read file
	contents, err := os.ReadFile(path)
	if err != nil {
		return &Config{}, nil
	}

	// Unmarshal JSON
	var cfg Config
	if err := json.Unmarshal([]byte(contents), &cfg); err != nil {
		return nil, err
	}
	cfg.path = path

	return &cfg, nil
}

func (cfg *Config) Save() error {
	// Create directory if it doesn't exist
	dir := filepath.Dir(cfg.path)
	if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	// Marshal JSON to file
	jsonBytes, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(cfg.path, jsonBytes, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Saved config to ", cfg.path)
	return nil
}

func (cfg *Config) Dump() error {
	jsonBytes, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(jsonBytes))
	return nil
}
