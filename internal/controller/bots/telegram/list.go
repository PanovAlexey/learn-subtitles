package telegram

import (
	"github.com/PanovAlexey/learn-subtitles/internal/domain/dto"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func (r CommandRouter) listCommand(inputMessage tgbotapi.Message, user dto.UserDatabaseDto) {
	r.logger.Info("list handler. from: " + inputMessage.From.UserName)

	dialog, _ := r.userStateService.GetUserDialog(strconv.FormatInt(user.Id.Int64, 10))

	dialog.SetHasSubtitlesListState()

	listMsg := "<strong>Pulp Fiction script</strong> \nType:<strong>subtitles</strong>\nGet: /download1001 \n\n" +
		"<strong>Harry Potter and chamber of secrets</strong> \nType:<strong>book</strong>\nGet: /download1002 \n\n" //ToDo: change for real data
	count := 2 //ToDo: change for real data

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Your subtitles list (<strong>"+strconv.Itoa(count)+"</strong> pieces): \n\n"+listMsg)
	msg.ParseMode = tgbotapi.ModeHTML
	r.bot.Send(msg)
}
