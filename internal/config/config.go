package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`
	Logger struct {
		Level string `mapstructure:"level"`
		File  string `mapstructure:"file"`
	} `mapstructure:"logger"`
}

func Load() (*Config, error) {
	v := viper.New()

	// 1. Чтение из файла `configs/config.yaml`
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("../configs") // Путь относительно `internal/config/`

	if err := v.ReadInConfig(); err != nil {
		log.Printf("Config file not found, using env vars. Error: %v", err)
	}

	// 2. Чтение переменных окружения (приоритет)
	v.AutomaticEnv()
	v.SetEnvPrefix("APP")
	v.BindEnv("server.port", "SERVER_PORT")
	v.BindEnv("logger.level", "LOGGER_LEVEL")
	v.BindEnv("logger.file", "LOGGER_FILE")

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	// Значения по умолчанию
	if cfg.Server.Port == "" {
		cfg.Server.Port = "8080"
	}
	if cfg.Logger.Level == "" {
		cfg.Logger.Level = "info"
	}

	return &cfg, nil
}
