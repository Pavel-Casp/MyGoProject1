package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`
	Logger struct {
		Level string `mapstructure:"level"`
	} `mapstructure:"logger"`
}

func Load() (*Config, error) {
	v := viper.New()

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "./configs"
	}

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(configPath)

	if err := v.ReadInConfig(); err != nil {
		log.Printf("Config file not found at %s, using env vars. Error: %v",
			filepath.Join(configPath, "config.yaml"), err)
	}

	v.AutomaticEnv()
	v.SetEnvPrefix("APP")
	v.BindEnv("server.port", "APP_SERVER_PORT")
	v.BindEnv("logger.level", "APP_LOGGER_LEVEL")

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	// Устанавливаем значения по умолчанию
	if cfg.Server.Port == "" {
		cfg.Server.Port = "8080"
	}
	if cfg.Logger.Level == "" {
		cfg.Logger.Level = "info"
	}

	return &cfg, nil
}
