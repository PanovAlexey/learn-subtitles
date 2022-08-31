package telegram

import (
	"github.com/PanovAlexey/learn-subtitles/internal/controller/bots/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type server struct {
	router           telegram.CommandRouter
	botUpdateTimeout int
}

func NewServer(router telegram.CommandRouter, botUpdateTimeout int) server {
	return server{
		router:           router,
		botUpdateTimeout: botUpdateTimeout,
	}
}

func (s server) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = s.botUpdateTimeout
	updates := s.router.GetBot().GetUpdatesChan(u)

	for update := range updates {
		s.router.HandleUpdate(update)
	}
}
