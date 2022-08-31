package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (r CommandRouter) defaultBehavior(inputMesage tgbotapi.Message) {
	r.logger.Info("default handler. from: " + inputMesage.From.UserName + ". text:" + inputMesage.Text)

	msg := tgbotapi.NewMessage(inputMesage.Chat.ID, "You wrote: "+inputMesage.Text)
	msg.ReplyToMessageID = inputMesage.MessageID

	r.bot.Send(msg)
}
