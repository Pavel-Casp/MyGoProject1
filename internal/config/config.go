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
	} `mapstructure:"logger"`
}

func Load() (*Config, error) {
	v := viper.New()

	// 1. Чтение из файла
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("../configs")

	if err := v.ReadInConfig(); err != nil {
		log.Printf("Config file not found, using env vars. Error: %v", err)
	}

	// 2. Чтение из переменных окружения
	v.AutomaticEnv()
	v.SetEnvPrefix("APP")
	v.BindEnv("server.port", "SERVER_PORT")
	v.BindEnv("logger.level", "LOGGER_LEVEL")

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
