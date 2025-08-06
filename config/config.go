package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Server   ServerConfig   `json:"server"`
	Database DatabaseConfig `json:"database"`
}

type ServerConfig struct {
	Port    int `json:"port"`
	Timeout int `json:"timeout"`
}

type DatabaseConfig struct {
	URI        string `json:"uri"`
	Name       string `json:"name"`
	Collection string `json:"collection"`
}

func Load() Config {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}
	defer file.Close()

	var cfg Config
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		log.Fatalf("Error parsing config: %v", err)
	}

	return cfg
}
