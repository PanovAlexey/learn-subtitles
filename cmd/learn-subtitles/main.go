package main

import (
	"github.com/PanovAlexey/learn-subtitles/internal/config"
	"github.com/PanovAlexey/learn-subtitles/pkg/logging"
	"github.com/joho/godotenv"
	"log"
)

type Logger interface {
	Error(args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("error loading env variables: %s", err.Error())
	}

	config, err := config.NewConfig()

	if err != nil {
		log.Fatalf("can't initialize config: %v", err)
	}

	logger, err := logging.GetLogger(config.GetApplicationName())

	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

func startTelegramBotServer() {
	 
}
