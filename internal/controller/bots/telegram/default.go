package telegram

import (
	"encoding/json"
	"github.com/PanovAlexey/learn-subtitles/internal/domain/dto"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func (r CommandRouter) defaultBehavior(inputMessage tgbotapi.Message, user dto.UserDatabaseDto) {
	r.logger.Info("default handler. from: " + inputMessage.From.UserName + ". text:" + inputMessage.Text)

	dialog, _ := r.userStateService.GetUserDialog(strconv.FormatInt(user.Id.Int64, 10))

	resultText := ""
	info, buttons, err := dialog.TryToHandleUserData(inputMessage.Text)

	if err != nil {
		r.logger.Debug("tg bot default handler: ", err)
		resultText = "You wrote:\n <i>" + inputMessage.Text + "</i>.\nInvalid input. Please repeat.\n\n" /* + r.getAvailableCommandListString()*/
	} else {
		resultText = info
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		resultText,
	)
	msg.ParseMode = tgbotapi.ModeHTML
	msg.ReplyToMessageID = inputMessage.MessageID

	if len(buttons) > 0 {
		jsonButtons := []tgbotapi.InlineKeyboardButton{}

		for _, button := range buttons {
			jsonButton, _ := json.Marshal(button)

			inlineKeyboardButtons := tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(button.Text, string(jsonButton)),
			)

			for _, v := range inlineKeyboardButtons {
				jsonButtons = append(jsonButtons, v)
			}
		}

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(jsonButtons)
	}

	r.bot.Send(msg)
}
