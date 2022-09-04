package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (r CommandRouter) defaultBehavior(inputMessage tgbotapi.Message) {
	r.logger.Info("default handler. from: " + inputMessage.From.UserName + ". text:" + inputMessage.Text)

	dialog, _ := r.userStateService.GetUserDialog(strconv.FormatInt(inputMessage.Chat.ID, 10))

	resultText := ""
	info, err := dialog.TryToHandleUserData(inputMessage.Text)

	if err != nil {
		r.logger.Debug("tg bot default handler: ", err)
		resultText = "You wrote:\n <i>" + inputMessage.Text + "</i>.\nInvalid input. Please repeat.\n\n" /* + r.getAvailableCommandListString()*/
	} else {
		resultText = info
	}

	msg := tgbotapi.NewMessage(inputMesage.Chat.ID, "You wrote: "+inputMesage.Text)
	msg.ReplyToMessageID = inputMesage.MessageID
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		resultText,
	)
	msg.ParseMode = tgbotapi.ModeHTML
	msg.ReplyToMessageID = inputMessage.MessageID

	r.bot.Send(msg)
}
