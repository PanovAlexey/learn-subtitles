package telegram

import (
	loggerInterface "github.com/PanovAlexey/learn-subtitles/internal/infrastructure/service/logging"
	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotResolver struct {
	token  string
	logger loggerInterface.Logger
}

func NewBotResolver(token string, logger loggerInterface.Logger) BotResolver {
	return BotResolver{
		token:  token,
		logger: logger,
	}
}

func (s BotResolver) GetTelegramBot() (*tgBotApi.BotAPI, error) {
	bot, err := tgBotApi.NewBotAPI(s.token)

	if err != nil {
		return nil, err
	}

	// debug mode ON
	bot.Debug = true
	s.logger.Info("Successful authorized on telegram bot account:" + bot.Self.UserName)

	return bot, nil
}
