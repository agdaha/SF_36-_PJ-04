package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	URLS        []string `json:"rss"`
	Period      int      `json:"request_period"`
	DatabaseURL string   `json:"db_url"`
	BindAddr    string   `json:"bind_address"`
}

// Парсинг файла настроек и создание структуры конфига
func New(configPath string) (*Config, error) {
	config := &Config{Period: 3, BindAddr: ":8383"}

	b, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла настроек: %v", err)
	}

	err = json.Unmarshal(b, &config)
	if err != nil {
		return nil, fmt.Errorf("ошибка парсинга файла настроек: %v", err)
	}
	return config, nil
}
