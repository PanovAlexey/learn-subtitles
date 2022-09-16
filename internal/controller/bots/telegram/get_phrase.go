package telegram

import (
	"github.com/PanovAlexey/learn-subtitles/internal/domain/dto"
	infrastructureDto "github.com/PanovAlexey/learn-subtitles/internal/infrastructure/dto"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func (r CommandRouter) getPhraseCommand(
	message tgbotapi.Message,
	user dto.UserDatabaseDto,
	commandButton infrastructureDto.CommandButton,
) {
	r.logger.Info("get phrase by subtitles handler. from: ", user.Id.Int64)

	dialog, _ := r.userStateService.GetUserDialog(strconv.FormatInt(user.Id.Int64, 10))
	dialog.SetSelectedSubtitlesState()

	phrase, buttons, err := dialog.GetRandomPhraseByCurrentSubtitles()

	if err != nil {
		r.showUnexpectedError(message, err)
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, phrase.Text.String+"\n")
	msg.ParseMode = tgbotapi.ModeHTML

	if len(buttons) > 0 {
		msg = r.addButtonsToMsg(msg, buttons)
	}

	r.bot.Send(msg)
}
