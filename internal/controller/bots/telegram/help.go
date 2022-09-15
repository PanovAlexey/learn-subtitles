package telegram

import (
	"github.com/PanovAlexey/learn-subtitles/internal/domain/dto"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func (r CommandRouter) helpCommand(inputMessage tgbotapi.Message, user dto.UserDatabaseDto) {
	r.logger.Info("help handler. from: " + inputMessage.From.UserName)

	dialog, _ := r.userStateService.GetUserDialog(strconv.FormatInt(user.Id.Int64, 10))
	dialog.SetRestState()

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		r.getAvailableCommandListString(),
	)

	msg.ParseMode = tgbotapi.ModeHTML
	r.bot.Send(msg)
}
