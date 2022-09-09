package config

import (
	"errors"
	"os"
	"strconv"
	"time"
)

var configSingleton *Config

type Config struct {
	applicationName          string
	telegramBotToken         string
	telegramBotUpdateTimeout int

	databaseUser          string
	databasePassword      string
	databaseAddress       string
	databasePort          string
	databaseName          string
	maxOpenConnections    int
	maxIdleConnections    int
	connectionMaxIdleTime time.Duration
	connectionMaxLifeTime time.Duration
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

func (c Config) GetDatabaseUser() string {
	return c.databaseUser
}

func (c Config) GetDatabasePassword() string {
	return c.databasePassword
}

func (c Config) GetDatabaseAddress() string {
	return c.databaseAddress
}

func (c Config) GetDatabasePort() string {
	return c.databasePort
}

func (c Config) GetDatabaseName() string {
	return c.databaseName
}

func (c Config) GetMaxOpenConnections() int {
	return c.maxOpenConnections
}

func (c Config) GetMaxIdleConnections() int {
	return c.maxIdleConnections
}

func (c Config) GetConnectionMaxIdleTime() time.Duration {
	return c.connectionMaxIdleTime
}

func (c Config) GetConnectionMaxLifeTime() time.Duration {
	return c.connectionMaxLifeTime
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

	c.databaseUser = getEnv("DB_USER")
	c.databasePassword = getEnv("DB_PASSWORD")
	c.databaseAddress = getEnv("DB_ADDRESS")
	c.databasePort = getEnv("DB_PORT")
	c.databaseName = getEnv("DB_NAME")

	maxOpenConnections := getEnv("DB_MAX_OPEN_CONNECTIONS")

	if len(maxOpenConnections) > 0 {
		c.maxOpenConnections, _ = strconv.Atoi(maxOpenConnections)
	}

	maxIdleConnections := getEnv("DB_MAX_IDLE_CONNECTIONS")

	if len(maxIdleConnections) > 0 {
		c.maxIdleConnections, _ = strconv.Atoi(maxIdleConnections)
	}

	connectionMaxIdleTime := getEnv("DB_CONNECTION_MAX_IDLE_TIME")

	if len(connectionMaxIdleTime) > 0 {
		connectionMaxIdleTimeInt64, err := strconv.ParseInt(connectionMaxIdleTime, 10, 64)

		if err != nil {
			return err
		}

		c.connectionMaxLifeTime = time.Duration(connectionMaxIdleTimeInt64)
	}

	connectionMaxLifeTime := getEnv("DB_CONNECTION_MAX_LIFE_TIME")

	if len(connectionMaxLifeTime) > 0 {
		connectionMaxLifeTimeInt64, err := strconv.ParseInt(connectionMaxLifeTime, 10, 64)

		if err != nil {
			return err
		}

		c.connectionMaxLifeTime = time.Duration(connectionMaxLifeTimeInt64)
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
