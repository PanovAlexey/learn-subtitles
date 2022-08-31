package telegram

import (
	"encoding/json"
	"fmt"
	"github.com/PanovAlexey/learn-subtitles/internal/application/service/phrase"
	"github.com/PanovAlexey/learn-subtitles/internal/application/service/subtitles"
	loggerInterface "github.com/PanovAlexey/learn-subtitles/internal/infrastructure/service/logging"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CommandRouter struct {
	bot              *tgbotapi.BotAPI
	logger           loggerInterface.Logger
	subtitlesService subtitles.SubtitlesService
	phraseService    phrase.PhraseService
}

// is needed to conveniently receive values passed by the user by pressing a button, rather than manually entering text.
type CommandData struct {
	Offset int `json:"offset"`
}

func NewRouter(
	bot *tgbotapi.BotAPI,
	logger loggerInterface.Logger,
	subtitlesService subtitles.SubtitlesService,
	phraseService phrase.PhraseService,
) CommandRouter {
	return CommandRouter{
		bot:              bot,
		logger:           logger,
		subtitlesService: subtitlesService,
		phraseService:    phraseService,
	}
}

func (r CommandRouter) GetBot() *tgbotapi.BotAPI {
	return r.bot
}

func (r CommandRouter) HandleUpdate(update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			r.logger.Error("telegram bot recovered from panic: ", panicValue)
		}
	}()

	if update.CallbackQuery != nil {
		parsedData := CommandData{}
		json.Unmarshal([]byte(update.CallbackQuery.Data), &parsedData)
		msg := tgbotapi.NewMessage(
			update.CallbackQuery.Message.Chat.ID,
			"Data from button: "+update.CallbackQuery.Data+
				fmt.Sprintf("Parsed: %+v\n", parsedData.Offset),
		)
		r.bot.Send(msg)
		return
	}

	if update.Message == nil {
		return
	}

	switch update.Message.Command() {
	case "help":
		r.helpCommand(*update.Message)
	case "list":
		r.listCommand(*update.Message)
	default:
		r.defaultBehavior(*update.Message)
	}
}
