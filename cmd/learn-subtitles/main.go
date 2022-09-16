package main

import (
	"github.com/PanovAlexey/learn-subtitles/internal/application/service/phrase"
	"github.com/PanovAlexey/learn-subtitles/internal/application/service/subtitles"
	"github.com/PanovAlexey/learn-subtitles/internal/application/service/user"
	"github.com/PanovAlexey/learn-subtitles/internal/config"
	"github.com/PanovAlexey/learn-subtitles/internal/controller/bots/telegram"
	"github.com/PanovAlexey/learn-subtitles/internal/infrastructure/repository"
	"github.com/PanovAlexey/learn-subtitles/internal/infrastructure/service/bot_state_machine"
	"github.com/PanovAlexey/learn-subtitles/internal/infrastructure/service/database/postgresql"
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

	postgresConnector, err := postgresql.GetPostgresConnector(
		config.GetDatabaseUser(),
		config.GetDatabasePassword(),
		config.GetDatabaseAddress(),
		config.GetDatabasePort(),
		config.GetDatabaseName(),
		config.GetMaxOpenConnections(),
		config.GetMaxIdleConnections(),
		config.GetConnectionMaxIdleTime(),
		config.GetConnectionMaxLifeTime(),
	)

	if err != nil {
		logger.Panic(err)
	}

	defer postgresConnector.DB.Close()

	subtitleRepository := repository.NewSubtitleRepository(postgresConnector)
	userRepository := repository.NewUserRepository(postgresConnector)
	phraseRepository := repository.NewPhraseRepository(postgresConnector)

	subtitlesService := subtitles.NewSubtitlesService(subtitleRepository)
	userService := user.NewUserService(userRepository)
	phraseService := phrase.NewPhraseService(phraseRepository)

	err = startTelegramBotServer(*config, logger, subtitlesService, userService, phraseService)

	if err != nil {
		logger.Panic(err)
	}
}

func startTelegramBotServer(
	config config.Config,
	logger loggerInterface.Logger,
	subtitlesService subtitles.SubtitlesService,
	userService user.UserService,
	phraseService phrase.PhraseService,
) error {
	botResolver := telegramService.NewBotResolver(config.GetTelegramBotToken(), logger)
	bot, err := botResolver.GetTelegramBot()

	if err != nil {
		return err
	}

	userStatesService := bot_state_machine.NewUserStatesService()

	telegramRouter := telegram.NewRouter(bot,
		logger,
		subtitlesService,
		phraseService,
		&userStatesService,
		userService)
	telegramServer := telegramServer.NewServer(telegramRouter, config.GetTelegramBotUpdateTimeout())
	telegramServer.Start()

	return nil
}
