package bot_state_machine

import (
	"errors"
	"github.com/PanovAlexey/learn-subtitles/internal/domain/entity"
)

type SelectedSubtitlesState struct {
	dialog *Dialog
}

func (s *SelectedSubtitlesState) GetCode() StatesEnum {
	return SelectedSubtitles
}

func (s *SelectedSubtitlesState) AddSubtitles() error {
	s.dialog.SetReadyToAddSubtitlesNameState()

	return nil
}

func (s *SelectedSubtitlesState) AddSubtitlesName(name string) error {
	return errors.New("not available command for current state")
}

func (s *SelectedSubtitlesState) AddSubtitlesText(text string) error {
	return errors.New("not available command for current state")
}

func (s *SelectedSubtitlesState) AddForbiddenPartsAndSaveSubtitles(subtitles entity.Subtitle, forbiddenPartsString string) error {
	return errors.New("not available command for current state")
}

func (s *SelectedSubtitlesState) GetSubtitlesList() ([]entity.Subtitle, error) {
	return []entity.Subtitle{}, nil
}

func (s *SelectedSubtitlesState) GetSubtitlesById() (*entity.Subtitle, error) {
	return &entity.Subtitle{}, nil
}

func (s *SelectedSubtitlesState) DeleteSubtitlesById() error {
	return nil
}

func (s *SelectedSubtitlesState) GetRandomPhraseBySubtitlesId() (*entity.Phrase, error) {
	return &entity.Phrase{}, nil
}

func (s *SelectedSubtitlesState) GetTranslateByPhraseId() (*entity.PhraseTranslation, error) {
	return &entity.PhraseTranslation{}, nil
}

func (s *SelectedSubtitlesState) SetTranslateByPhraseId(translation string) error {
	return nil
}

func (s *SelectedSubtitlesState) HideThisPhraseById() error {
	return nil
}
