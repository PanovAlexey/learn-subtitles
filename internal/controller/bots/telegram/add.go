package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func (r CommandRouter) addCommand(inputMessage tgbotapi.Message) {
	r.logger.Info("list handler. from: " + inputMessage.From.UserName)

	dialog, _ := r.userStateService.GetUserDialog(strconv.FormatInt(inputMessage.Chat.ID, 10))

	dialog.SetReadyToAddSubtitlesNameState()

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Please enter the title of new text:\n")
	msg.ParseMode = tgbotapi.ModeHTML
	r.bot.Send(msg)
}
