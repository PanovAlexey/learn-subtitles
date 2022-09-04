package main

import (
	"github.com/PanovAlexey/learn-subtitles/internal/application/service/phrase"
	"github.com/PanovAlexey/learn-subtitles/internal/application/service/subtitles"
	"github.com/PanovAlexey/learn-subtitles/internal/config"
	"github.com/PanovAlexey/learn-subtitles/internal/controller/bots/telegram"
	"github.com/PanovAlexey/learn-subtitles/internal/infrastructure/service/bot_state_machine"
	loggerInterface "github.com/PanovAlexey/learn-subtitles/internal/infrastructure/service/logging"
	telegramService "github.com/PanovAlexey/learn-subtitles/internal/infrastructure/service/telegram"
	telegramServer "github.com/PanovAlexey/learn-subtitles/internal/server/bots/telegram"
	"github.com/PanovAlexey/learn-subtitles/pkg/logging"
	"github.com/joho/godotenv"
	"log"
)

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

	defer logger.Sync()

	err = startTelegramBotServer(*config, logger)

	if err != nil {
		logger.Panic(err)
	}
}

func startTelegramBotServer(config config.Config, logger loggerInterface.Logger) error {
	botResolver := telegramService.NewBotResolver(config.GetTelegramBotToken(), logger)
	bot, err := botResolver.GetTelegramBot()

	if err != nil {
		return err
	}

	subtitlesService := subtitles.NewSubtitlesService(config)
	phraseService := phrase.NewPhraseService()
	userStatesService := bot_state_machine.NewUserStatesService()

	telegramRouter := telegram.NewRouter(bot, logger, subtitlesService, phraseService, &userStatesService)
	telegramServer := telegramServer.NewServer(telegramRouter, config.GetTelegramBotUpdateTimeout())
	telegramServer.Start()

	return nil
}
