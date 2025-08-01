/*
Copyright © 2025 Abinand P <abinand0911@gmail.com>
*/

package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	RedisAddr              string `json:"redis_addr"`
	GEMINI_API_KEY         string `json:"gemini_api_key"`
	GEMINI_EMBEDDING_MODEL string `json:"gemini_embedding_model"`
	GEMINI_CHAT_MODEL      string `json:"gemini_chat_model"`
}

func getConfigPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "PesudoCLI", "config.json")
}

func LoadConfig() (*Config, error) {
	configPath := getConfigPath()
	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {

			defaultCfg := &Config{
				RedisAddr:              "localhost:6379",
				GEMINI_API_KEY:         "",
				GEMINI_EMBEDDING_MODEL: "gemini-embedding-001",
				GEMINI_CHAT_MODEL:      "gemini-2.0-flash",
			}
			if err := SaveConfig(defaultCfg); err != nil {
				return nil, err
			}
			return defaultCfg, nil
		}
		return nil, err
	}
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func SaveConfig(config *Config) error {
	configPath := getConfigPath()
	if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(configPath, data, 0644)
}
