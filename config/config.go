package config

import (
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Logging Logger `yaml:"logging"`
	Server  Server `yaml:"server"`
}

type Logger struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func MustInitEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func InitConfig() (*Config, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
