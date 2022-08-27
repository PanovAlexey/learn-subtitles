package main

import (
	"github.com/PanovAlexey/learn-subtitles/internal/config"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("error loading env variables: %s", err.Error())
	}

	config, err := config.NewConfig()

	if err != nil {
		// @ToDo: add logging
	}
}

func startTelegramBotServer() {
	 
}
