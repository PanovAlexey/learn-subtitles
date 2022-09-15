package telegram

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PanovAlexey/learn-subtitles/internal/application/service/phrase"
	"github.com/PanovAlexey/learn-subtitles/internal/application/service/subtitles"
	service "github.com/PanovAlexey/learn-subtitles/internal/application/service/user"
	"github.com/PanovAlexey/learn-subtitles/internal/domain/dto"
	"github.com/PanovAlexey/learn-subtitles/internal/domain/entity"
	"github.com/PanovAlexey/learn-subtitles/internal/infrastructure/service/bot_state_machine"
	loggerInterface "github.com/PanovAlexey/learn-subtitles/internal/infrastructure/service/logging"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

type CommandRouter struct {
	bot              *tgbotapi.BotAPI
	logger           loggerInterface.Logger
	subtitlesService subtitles.SubtitlesService
	phraseService    phrase.PhraseService
	userStateService *bot_state_machine.UserStatesService
	userService      service.UserService
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
	userStateService *bot_state_machine.UserStatesService,
	userService service.UserService,
) CommandRouter {
	return CommandRouter{
		bot:              bot,
		logger:           logger,
		subtitlesService: subtitlesService,
		phraseService:    phraseService,
		userStateService: userStateService,
		userService:      userService,
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
			// fmt.Sprintf(". Command: %s\n", args[0])+
			// fmt.Sprintf("Offset: %s\n", args[1]),
		)

		r.bot.Send(msg)
		return
	}

	if update.Message == nil {
		r.showUnexpectedError(update.Message, errors.New("update message is nil"))
		return
	}

	user, err := r.checkUser(*update.Message.From)

	if err != nil {
		r.showUnexpectedError(update.Message, err)
		return
	}

	d, ok := r.userStateService.GetUserDialog(strconv.FormatInt(user.Id.Int64, 10))

	if ok == false {
		d = bot_state_machine.NewDialog(user.Id.Int64, r.subtitlesService)
		r.userStateService.SetUserDialog(d)
	}
	switch update.Message.Command() {
	case "help":
		r.helpCommand(*update.Message, user)
	case "list":
		r.listCommand(*update.Message, user)
	case "add":
		r.addCommand(*update.Message, user)
	case "debug":
		r.debugCommand(*update.Message, user)
	default:
		r.defaultBehavior(*update.Message, user)
	}
}

func (r CommandRouter) showUnexpectedError(inputMessage *tgbotapi.Message, err error) {
	r.logger.Error(err)

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Unexpected error occurred. Please try again or contact administrator.",
	)
	msg.ParseMode = tgbotapi.ModeHTML
	r.bot.Send(msg)
}

func (r CommandRouter) getAvailableCommandListString() string {
	return "<strong>Available commands:</strong>\n\n" +
		"/add - add a new text entry to study\n" +
		"/list - get a list of text entries to study\n" +
		"/language - set interface language\n" +
		"/help - get list of available commands\n"
}

func (r CommandRouter) checkUser(tgUser tgbotapi.User) (dto.UserDatabaseDto, error) {
	user, err := r.userService.GetUserByLogin(tgUser.UserName)

	if err != nil {
		return user, errors.New("getting user by login error: " + err.Error() + ". Login: " + tgUser.UserName)
	}

	if !user.Id.Valid || user.Id.Int64 < 1 {
		user, err = r.userService.SaveUser(entity.User{
			Login:     tgUser.UserName,
			FirstName: tgUser.FirstName,
			LastName:  tgUser.LastName,
			IsDeleted: false,
		})

		if err != nil {
			return user, errors.New("user registration error: " + err.Error() + ". Login: " + tgUser.UserName)
		}

		if !user.Id.Valid || user.Id.Int64 < 1 {
			return user, errors.New("user registration error: " + err.Error() + ". Login: " + tgUser.UserName)
		}
	}

	return user, nil
}
