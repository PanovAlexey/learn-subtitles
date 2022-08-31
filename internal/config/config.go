package config

import (
	"errors"
	"os"
	"strconv"
)

var configSingleton *Config

type Config struct {
	applicationName          string
	telegramBotToken         string
	telegramBotUpdateTimeout int
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

func (c Config) GetTelegramBotUpdateTimeout() int {
	return c.telegramBotUpdateTimeout
}

func (c *Config) initConfigByEnv() error {
	c.applicationName = getEnv("APPLICATION_NAME")
	c.telegramBotToken = getEnv("TELEGRAM_BOT_TOKEN")
	telegramBotUpdateTimeout := getEnv("TELEGRAM_BOT_UPDATE_TIMEOUT")

	if len(telegramBotUpdateTimeout) > 0 {
		c.telegramBotUpdateTimeout, _ = strconv.Atoi(telegramBotUpdateTimeout)
	}

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

	if c.telegramBotUpdateTimeout == 0 {
		c.telegramBotUpdateTimeout = 60
	}
}
