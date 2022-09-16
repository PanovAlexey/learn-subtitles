package telegram

import (
	"encoding/json"
	"errors"
	"github.com/PanovAlexey/learn-subtitles/internal/application/service/phrase"
	"github.com/PanovAlexey/learn-subtitles/internal/application/service/subtitles"
	service "github.com/PanovAlexey/learn-subtitles/internal/application/service/user"
	"github.com/PanovAlexey/learn-subtitles/internal/domain/dto"
	"github.com/PanovAlexey/learn-subtitles/internal/domain/entity"
	infrastructureDto "github.com/PanovAlexey/learn-subtitles/internal/infrastructure/dto"
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

	user := *r.getUserByUpdate(update)

	d, ok := r.userStateService.GetUserDialog(strconv.FormatInt(user.Id.Int64, 10))

	if ok == false {
		d = bot_state_machine.NewDialog(user.Id.Int64, r.subtitlesService, r.phraseService)
		r.userStateService.SetUserDialog(d)
	}

	command := r.getCommandByUpdate(update)

	/////////////
		r.showUnexpectedError(update.Message, err)
		return
	}

	d, ok := r.userStateService.GetUserDialog(strconv.FormatInt(user.Id.Int64, 10))

	switch command {
	case "help":
		r.helpCommand(*update.Message, user)
	case "list":
		r.listCommand(*update.Message, user)
	case "add":
		r.addCommand(*update.Message, user)
	case "del_sub":
		r.deleteSubtitlesCommand(*update.Message, user)
	case "get_p":
		r.getPhraseCommand(*update.Message, user)
	case "debug":
		r.debugCommand(*update.Message, user)
	default:
		r.defaultBehavior(*update.Message, user)
	}
}

func (r CommandRouter) getUserByUpdate(update tgbotapi.Update) *dto.UserDatabaseDto {
	var inputUser tgbotapi.User

	if update.Message != nil {
		inputUser = *update.Message.From
	} else if update.CallbackQuery.From != nil {
		inputUser = *update.CallbackQuery.From
	} else {
		r.showUnexpectedError(update.Message, errors.New("update message and callback query both are nil"))
		return nil
	}

	user, err := r.checkUser(inputUser)

	if err != nil {
		r.showUnexpectedError(update.Message, err)
		return nil
	}

	return &user
}

func (r CommandRouter) getCommandByUpdate(update tgbotapi.Update) string {
	var command string

	if update.Message != nil {
		command = update.Message.Command()
	} else if update.CallbackQuery.From != nil {
		parsedData := infrastructureDto.CommandButton{}
		json.Unmarshal([]byte(update.CallbackQuery.Data), &parsedData)
		command = parsedData.Command
	} else {
		r.showUnexpectedError(update.Message, errors.New("update message and callback query both are nil"))
		return ""
	}

	return command
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
