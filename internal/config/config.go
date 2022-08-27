package config

import (
	"errors"
	"os"
)

var configSingleton *Config

type Config struct {
	telegramBotToken string
}

func NewConfig() (*Config, error) {
	if configSingleton == nil {
		config := &Config{}
		err := initConfigByEnv(config)

		if err != nil {
			return nil, err
		}

		configSingleton = config
	}

	return configSingleton, nil
}

func (c Config) GetTelegramBotToken() string {
	return c.telegramBotToken
}

func initConfigByEnv(config *Config) error {
	config.telegramBotToken = getEnv("TELEGRAM_BOT_TOKEN")

	if len(config.telegramBotToken) < 1 {
		return errors.New("telegram bot token is empty")
	}

	return nil
}

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return ""
}
