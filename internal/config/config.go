package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env     string `yaml:"env" env-required:"true"`
	Storage string `yaml:"storage_path" env-required:"true"`
	Logger  `yaml:"logger" env-required:"true"`
	TgCreds `yaml:"tg_creds" env-required:"true"`
}

type TgCreds struct {
	Token string `yaml:"api_token" env-required:"true"`
}

type Logger struct {
	LogType  string `yaml:"log_type" env-required:"true"`
	LogLevel int8   `yaml:"log_level" env-required:"true"`
}

func LoadCFG() (error, *Config) {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		return errors.New("CONFIG_PATH is not defined"), nil
	}

	// check if file exist
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return errors.New(fmt.Sprintf("config file '%s' does not exist", configPath)), nil
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return errors.New(fmt.Sprintf("cannot read config: %s", err)), nil
	}

	return nil, &cfg
}
