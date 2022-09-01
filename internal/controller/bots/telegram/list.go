package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func (r CommandRouter) listCommand(inputMessage tgbotapi.Message) {
	r.logger.Info("list handler. from: " + inputMessage.From.UserName)

	dialog, _ := r.userStateService.GetUserDialog(strconv.FormatInt(inputMessage.Chat.ID, 10))

	dialog.SetHasSubtitlesListState()

	listMsg := "<strong>Pulp Fiction script</strong> \nType:<strong>subtitles</strong>\nGet: /download1001 \n\n" +
		"<strong>Harry Potter and chamber of secrets</strong> \nType:<strong>book</strong>\nGet: /download1002 \n\n" //ToDo: change for real data
	count := 2 //ToDo: change for real data

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Your subtitles list (<strong>"+strconv.Itoa(count)+"</strong> pieces): \n\n"+listMsg)
	msg.ParseMode = tgbotapi.ModeHTML
	r.bot.Send(msg)
}
