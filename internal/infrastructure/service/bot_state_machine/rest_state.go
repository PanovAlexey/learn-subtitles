package bot_state_machine

import (
	"errors"
	"github.com/PanovAlexey/learn-subtitles/internal/domain/entity"
)

type RestState struct {
	dialog *Dialog
}

func (s *RestState) GetCode() StatesEnum {
	return Rest
}

func (s *RestState) AddSubtitles() error {
	s.dialog.SetReadyToAddSubtitlesNameState()

	return nil
}

func (s *RestState) AddSubtitlesName(name string) error {
	return errors.New("not available command for current state")
}

func (s *RestState) AddSubtitlesText(text string) error {
	return errors.New("not available command for current state")
}

func (s *RestState) AddForbiddenPartsAndSaveSubtitles(
	subtitles entity.Subtitle,
	forbiddenPartsString string,
) (*entity.Subtitle, error) {
	return nil, errors.New("not available command for current state")
}

func (s *RestState) GetSubtitlesList() ([]entity.Subtitle, error) {
	return []entity.Subtitle{}, nil
}

func (s *RestState) GetSubtitlesById() (*entity.Subtitle, error) {
	return &entity.Subtitle{}, nil
}

func (s *RestState) DeleteSubtitlesById() error {
	return nil
}

func (s *RestState) GetRandomPhraseBySubtitlesId() (*entity.Phrase, error) {
	return &entity.Phrase{}, nil
}

func (s *RestState) GetTranslateByPhraseId() (*entity.PhraseTranslation, error) {
	return &entity.PhraseTranslation{}, nil
}

func (s *RestState) SetTranslateByPhraseId(translation string) error {
	return nil
}

func (s *RestState) HideThisPhraseById() error {
	return nil
}
