package bot_state_machine

import "github.com/PanovAlexey/learn-subtitles/internal/domain/entity"

type State interface {
	AddSubtitles() (*entity.Subtitle, error)
	AddSubtitlesName(name string) (*entity.Subtitle, error)
	AddSubtitlesText(text string) (*entity.Subtitle, error)
	AddSubtitlesForbiddenParts(forbiddenParts []string) (*entity.Subtitle, error)
	GetSubtitlesList() ([]entity.Subtitle, error)
	GetSubtitlesById() (*entity.Subtitle, error)
	DeleteSubtitlesById() error
	GetRandomPhraseBySubtitlesId() (*entity.Phrase, error)
	GetTranslateByPhraseId() (*entity.PhraseTranslation, error)
	SetTranslateByPhraseId(translation string) error
	HideThisPhraseById() error
	GetCode() int
}
