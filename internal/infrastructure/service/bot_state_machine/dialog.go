package bot_state_machine

import (
	"errors"
	"fmt"
	customErrors "github.com/PanovAlexey/learn-subtitles/internal/application/errors"
	"github.com/PanovAlexey/learn-subtitles/internal/application/service/phrase"
	"github.com/PanovAlexey/learn-subtitles/internal/application/service/subtitles"
	"github.com/PanovAlexey/learn-subtitles/internal/domain/entity"
	"github.com/PanovAlexey/learn-subtitles/internal/infrastructure/dto"
	"strconv"
)

type Dialog struct {
	subtitlesService subtitles.SubtitlesService
	phraseService    phrase.PhraseService

	rest                               State
	readyToAddSubtitlesName            State
	readyToAddSubtitlesText            State
	readyToAddSubtitlesProhibitedWords State
	hasSubtitlesList                   State
	selectedSubtitles                  State
	hasRandomPhrase                    State
	readyToAddTranslationRandomPhrase  State

	currentState      State
	userId            int64
	subtitles         entity.Subtitle
	phrase            entity.Phrase
	phraseTranslation entity.PhraseTranslation
}

func NewDialog(userId int64, subtitlesService subtitles.SubtitlesService) *Dialog {
func NewDialog(userId int64, subtitlesService subtitles.SubtitlesService, phraseService phrase.PhraseService) *Dialog {
	d := &Dialog{}
	d.subtitlesService = subtitlesService
	d.phraseService = phraseService

	d.userId = userId
	d.rest = &RestState{dialog: d}
	d.readyToAddSubtitlesName = &ReadyToAddSubtitlesNameState{dialog: d}
	d.readyToAddSubtitlesText = &ReadyToAddSubtitlesTextState{dialog: d}
	d.readyToAddSubtitlesProhibitedWords = &ReadyToAddSubtitlesProhibitedWordsState{dialog: d}
	d.hasSubtitlesList = &HasSubtitlesListState{dialog: d}
	d.selectedSubtitles = &SelectedSubtitlesState{dialog: d}
	d.hasRandomPhrase = &HasRandomPhraseState{dialog: d}
	d.readyToAddTranslationRandomPhrase = &ReadyToAddTranslationRandomPhraseState{dialog: d}

	d.SetRestState() // set state by default

	return d
}

func (d *Dialog) TryToHandleUserData(data string) (string, []dto.CommandButton, error) {
	stateCode := d.currentState.GetCode()

	switch stateCode {
	case Rest:
		return "",
			nil,
			errors.New(
				"no command selected for interaction. Input: " + data,
			)
	case ReadyToAddSubtitlesName:
		info, err := d.AddSubtitlesName(data)

		if err != nil {
			return info, nil, err
		}

		d.subtitles.Name.String = data
		d.subtitles.Author.Int64 = d.userId

		return info, nil, nil
	case ReadyToAddSubtitlesText:
		info, err := d.AddSubtitlesText(data)

		if err != nil {
			return info, nil, err
		}

		d.subtitles.Text.String = data
		d.subtitles.Author.Int64 = d.userId

		return info, nil, nil
	case ReadyToAddSubtitlesProhibitedWords:
		info, buttons, err := d.AddForbiddenPartsAndSaveSubtitles(d.subtitles, data)

		if err != nil {
			return info, buttons, err
		}

		return info, buttons, nil
	case HasSubtitlesList:

	case SelectedSubtitles:

	case HasRandomPhrase:

	case ReadyToAddTranslationRandomPhrase:
	default:
		d.SetRestState()

		return "", nil, errors.New(
			"wrong current state code: " + strconv.Itoa(int(stateCode)),
		)
	}

	return "", nil, nil
}

func (d *Dialog) AddSubtitles() error {
	return d.currentState.AddSubtitles()
}

func (d *Dialog) AddSubtitlesName(name string) (string, error) {
	err := d.currentState.AddSubtitlesName(name)

	if err != nil {
		if errors.As(err, &customErrors.ErrIsEmpty) {
			return "The name must be at least 5 letters. Please re-enter.", nil //ToDo: move magic number to constants
		}

		if errors.As(err, &customErrors.ErrTooLong) {
			return "The name must be less than 100 characters. Please re-enter.", nil //ToDo: move magic number to constants
		}
	}

	info := "Please enter the text to study:"

	return info, err
}

func (d *Dialog) AddSubtitlesText(text string) (string, error) {
	err := d.currentState.AddSubtitlesText(text)

	if err != nil {
		if errors.As(err, &customErrors.ErrIsEmpty) {
			return "The text must be at least 100 letters. Please re-enter.", nil //@ToDo: move magic number to constants
		}

		if errors.As(err, &customErrors.ErrTooLong) {
			return "The text must be less than 100 000 characters. Please re-enter.", nil //@ToDo: move magic number to constants
		}
	}

	info := "Please enter the forbidden parts to replace:"

	return info, err
}

func (d *Dialog) AddForbiddenPartsAndSaveSubtitles(
	subtitles entity.Subtitle,
	forbiddenPartsString string,
) (string, []dto.CommandButton, error) {
	resultSubtitles, err := d.currentState.AddForbiddenPartsAndSaveSubtitles(subtitles, forbiddenPartsString)
	info := ""

	if err != nil {
		if errors.As(err, &customErrors.ErrTooLong) {
			return "The length of replacement phrases should not " +
				"exceed the length of the main text. Please re-enter.", nil, nil //@ToDo: move magic number to constants
		}

		return info, nil, err
	} else {
		d.subtitles = *resultSubtitles

		if len(d.subtitles.ForbiddenParts) < 1 {
			if len(forbiddenPartsString) < 1 {
				info = "You have not specified a phrase to replace.\n\n"
			}
		}

		info += "You have successfully added text!\n" +
			"<strong>name:</strong> " + d.subtitles.Name.String + "\n" +
			"<strong>length of text:</strong> " + strconv.Itoa(len(d.subtitles.Text.String)) + "\n" +
			"<strong>spoiler substitution map:</strong> " + fmt.Sprintf("%+v\n", d.subtitles.ForbiddenParts) + "\n" +
			"<strong>Text:</strong> " + fmt.Sprintf("%+v\n", d.subtitles.Text.String) + "\n"
	}

	var buttons []dto.CommandButton

	buttons = append(
		buttons,
		dto.CommandButton{Id: d.subtitles.Id.Int64, Command: "get_p", Text: "Check phrase"},
		dto.CommandButton{Id: d.subtitles.Id.Int64, Command: "del_sub", Text: "Delete text"},
		dto.CommandButton{Id: d.subtitles.Id.Int64, Command: "help", Text: "To menu"},
	)

	return info, buttons, err
}

func (d *Dialog) GetSubtitlesList() ([]entity.Subtitle, error) {
	return d.currentState.GetSubtitlesList()
}

func (d *Dialog) GetSubtitlesById() (*entity.Subtitle, error) {
	return d.currentState.GetSubtitlesById()
}

func (d *Dialog) DeleteSubtitlesById() error {
	return d.currentState.DeleteSubtitlesById()
}

func (d *Dialog) GetRandomPhraseByCurrentSubtitles() (entity.Phrase, []dto.CommandButton, error) {
	var buttons []dto.CommandButton

	buttons = append(
		buttons,
		dto.CommandButton{Id: d.subtitles.Id.Int64, Command: "get_p", Text: "Check phrase"},
		dto.CommandButton{Id: d.subtitles.Id.Int64, Command: "del_p", Text: "Delete phrase"},
		dto.CommandButton{Id: d.subtitles.Id.Int64, Command: "help", Text: "To menu"},
	)

	phrase, err := d.currentState.GetRandomPhraseByCurrentSubtitles()

	return phrase, buttons, err
}

func (d *Dialog) GetTranslateByPhraseId() (*entity.PhraseTranslation, error) {
	return d.currentState.GetTranslateByPhraseId()
}

func (d *Dialog) SetTranslateByPhraseId(translation string) error {
	return d.currentState.SetTranslateByPhraseId(translation)
}

func (d *Dialog) HideThisPhraseById() error {
	return d.currentState.HideThisPhraseById()
}

func (d *Dialog) SetRestState() {
	d.setState(&RestState{dialog: d})
}

func (d *Dialog) SetReadyToAddSubtitlesNameState() {
	d.setState(&ReadyToAddSubtitlesNameState{dialog: d})
}

func (d *Dialog) SetReadyToAddSubtitlesTextState() {
	d.setState(&ReadyToAddSubtitlesTextState{dialog: d})
}

func (d *Dialog) SetReadyToAddSubtitlesProhibitedWordsState() {
	d.setState(&ReadyToAddSubtitlesProhibitedWordsState{dialog: d})
}

func (d *Dialog) SetHasSubtitlesListState() {
	d.setState(&HasSubtitlesListState{dialog: d})
}

func (d *Dialog) SetSelectedSubtitlesState() {
	d.setState(&SelectedSubtitlesState{dialog: d})
}

func (d *Dialog) SetHasRandomPhraseState() {
	d.setState(&HasRandomPhraseState{dialog: d})
}

func (d *Dialog) SetReadyToAddTranslationRandomPhraseState() {
	d.setState(&ReadyToAddTranslationRandomPhraseState{dialog: d})
}

func (d *Dialog) setState(s State) {
	d.currentState = s
}
