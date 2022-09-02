package bot_state_machine

import (
	"github.com/PanovAlexey/learn-subtitles/internal/application/service/subtitles"
	"github.com/PanovAlexey/learn-subtitles/internal/domain/entity"
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
		err := d.AddSubtitlesName(data)

		if err != nil {
			return "", err
		}

		d.subtitles.Name = data

		return "Please enter the text to study:", nil
	case ReadyToAddSubtitlesText:
		err := d.AddSubtitlesText(data)

		if err != nil {
			return "", err
		}

		d.subtitles.Text = data

		return "Please enter the forbidden parts to replace:", nil
	case ReadyToAddSubtitlesProhibitedWords:
		err := d.AddForbiddenPartsAndSaveSubtitles(d.subtitles, data)

		if err != nil {
			return "", err
		}

		info := "You have successfully added text!\n" +
			"<strong>name:</strong> " + d.subtitles.Name + "\n" +
			"<strong>length:</strong> " + strconv.Itoa(len(d.subtitles.Text)) + "\n" +
			"<strong>spoiler substitution map:</strong> " + fmt.Sprintf("%+v\n", d.subtitles.ForbiddenParts) + "\n"

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

func (d *Dialog) AddSubtitlesName(name string) error {
	return d.currentState.AddSubtitlesName(name)
}

func (d *Dialog) AddSubtitlesText(text string) error {
	return d.currentState.AddSubtitlesText(text)
}

func (d *Dialog) AddForbiddenPartsAndSaveSubtitles(subtitles entity.Subtitle, forbiddenPartsString string) error {
	return d.currentState.AddForbiddenPartsAndSaveSubtitles(subtitles, forbiddenPartsString)
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
