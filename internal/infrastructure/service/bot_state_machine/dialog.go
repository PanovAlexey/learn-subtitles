package bot_state_machine

import (
	"errors"
	"fmt"
	customErrors "github.com/PanovAlexey/learn-subtitles/internal/application/errors"
	"github.com/PanovAlexey/learn-subtitles/internal/application/service/subtitles"
	"github.com/PanovAlexey/learn-subtitles/internal/domain/entity"
	"strconv"
)

type Dialog struct {
	subtitlesService subtitles.SubtitlesService

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
	d := &Dialog{}
	d.subtitlesService = subtitlesService
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

func (d *Dialog) TryToHandleUserData(data string) (string, error) {
	stateCode := d.currentState.GetCode()

	switch stateCode {
	case Rest:
		return "",
			errors.New(
				"no command selected for interaction. Input: " + data,
			)
	case ReadyToAddSubtitlesName:
		info, err := d.AddSubtitlesName(data)

		if err != nil {
			return info, err
		}

		d.subtitles.Name = data

		return info, nil
	case ReadyToAddSubtitlesText:
		info, err := d.AddSubtitlesText(data)

		if err != nil {
			return info, err
		}

		d.subtitles.Text = data

		return info, nil
	case ReadyToAddSubtitlesProhibitedWords:
		info, err := d.AddForbiddenPartsAndSaveSubtitles(d.subtitles, data)

		if err != nil {
			return info, err
		}

		return info, nil
	case HasSubtitlesList:

	case SelectedSubtitles:

	case HasRandomPhrase:

	case ReadyToAddTranslationRandomPhrase:
	default:
		d.SetRestState()

		return "", errors.New(
			"wrong current state code: " + strconv.Itoa(int(stateCode)),
		)
	}

	return "", nil
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

func (d *Dialog) AddForbiddenPartsAndSaveSubtitles(subtitles entity.Subtitle, forbiddenPartsString string) (string, error) {
	resultSubtitles, err := d.currentState.AddForbiddenPartsAndSaveSubtitles(subtitles, forbiddenPartsString)
	info := ""

	if err != nil {
		if errors.As(err, &customErrors.ErrTooLong) {
			return "The length of replacement phrases should not exceed the length of the main text. Please re-enter.", nil //@ToDo: move magic number to constants
		}

		return info, err
	} else {
		d.subtitles = *resultSubtitles

		if len(d.subtitles.ForbiddenParts) < 1 {
			if len(forbiddenPartsString) < 1 {
				info = "You have not specified a phrase to replace.\n\n"
			} else {
				info = "Failed to parse string with replacement phrases.\n\n"
			}
		}

		info += "You have successfully added text!\n" +
			"<strong>name:</strong> " + d.subtitles.Name + "\n" +
			"<strong>length of text:</strong> " + strconv.Itoa(len(d.subtitles.Text)) + "\n" +
			"<strong>spoiler substitution map:</strong> " + fmt.Sprintf("%+v\n", d.subtitles.ForbiddenParts) + "\n" +
			"<strong>Text:</strong> " + fmt.Sprintf("%+v\n", d.subtitles.Text) + "\n"
	}

	return info, err
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

func (d *Dialog) GetRandomPhraseBySubtitlesId() (*entity.Phrase, error) {
	return d.currentState.GetRandomPhraseBySubtitlesId()
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
