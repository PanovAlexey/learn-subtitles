package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (r CommandRouter) helpCommand(inputMesage tgbotapi.Message) {
	r.logger.Info("help handler. from: " + inputMesage.From.UserName)

	msg := tgbotapi.NewMessage(
		inputMesage.Chat.ID,
		"/help - help\n"+
			"/list - list subtitles ",
	)
	r.bot.Send(msg)
}
