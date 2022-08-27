package config

import (
	"errors"
	"os"
)

var configSingleton *Config

type Config struct {
	applicationName  string
	telegramBotToken string
}

func NewConfig() (*Config, error) {
	if configSingleton == nil {
		config := &Config{}
		err := config.initConfigByEnv()

		if err != nil {
			return nil, err
		}

		config.initConfigByDefault()
		configSingleton = config
	}

	return configSingleton, nil
}

func (c Config) GetApplicationName() string {
	return c.applicationName
}

func (c Config) GetTelegramBotToken() string {
	return c.telegramBotToken
}

func (c *Config) initConfigByEnv() error {
	c.applicationName = getEnv("APPLICATION_NAME")
	c.telegramBotToken = getEnv("TELEGRAM_BOT_TOKEN")

	if len(c.telegramBotToken) < 1 {
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

func (c *Config) initConfigByDefault() {
	if len(c.applicationName) < 1 {
		c.applicationName = "Learn subtitles"
	}
}
