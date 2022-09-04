package bot_state_machine

import "github.com/PanovAlexey/learn-subtitles/internal/domain/entity"

type State interface {
	AddSubtitles() error
	AddSubtitlesName(name string) error
	AddSubtitlesText(text string) error
	AddForbiddenPartsAndSaveSubtitles(subtitles entity.Subtitle, forbiddenPartsString string) (*entity.Subtitle, error)
	GetSubtitlesList() ([]entity.Subtitle, error)
	GetSubtitlesById() (*entity.Subtitle, error)
	DeleteSubtitlesById() error
	GetRandomPhraseBySubtitlesId() (*entity.Phrase, error)
	GetTranslateByPhraseId() (*entity.PhraseTranslation, error)
	SetTranslateByPhraseId(translation string) error
	HideThisPhraseById() error
	GetCode() StatesEnum
}
