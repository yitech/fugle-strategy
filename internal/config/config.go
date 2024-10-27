package config

import (
	"os"

	"github.com/yosuke-furukawa/json5/encoding/json5"
)

type Config struct {
	Port  int         `json:"port"`  // Port as an integer
	Redis RedisConfig `json:"redis"` // Redis configuration
}

type RedisConfig struct {
	Host     string `json:"host"`     // Host as a string
	Port     int    `json:"port"`     // Port as an integer
	Password string `json:"password"` // Password as a string
	DB       int    `json:"db"`       // DB as an integer
}

func ReadConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config Config
	err = json5.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
