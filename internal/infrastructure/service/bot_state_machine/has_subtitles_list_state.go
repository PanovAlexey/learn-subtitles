package bot_state_machine

import (
	"errors"
	"github.com/PanovAlexey/learn-subtitles/internal/domain/entity"
)

type HasSubtitlesListState struct {
	dialog *Dialog
}

func (s *HasSubtitlesListState) GetCode() StatesEnum {
	return HasSubtitlesList
}

func (s *HasSubtitlesListState) AddSubtitles() error {
	s.dialog.SetReadyToAddSubtitlesNameState()

	return nil
}

func (s *HasSubtitlesListState) AddSubtitlesName(name string) error {
	return errors.New("not available command for current state")
}

func (s *HasSubtitlesListState) AddSubtitlesText(text string) error {
	return errors.New("not available command for current state")
}

func (s *HasSubtitlesListState) AddForbiddenPartsAndSaveSubtitles(
	subtitles entity.Subtitle,
	forbiddenPartsString string,
) (*entity.Subtitle, error) {
	return nil, errors.New("not available command for current state")
}

func (s *HasSubtitlesListState) GetSubtitlesList() ([]entity.Subtitle, error) {
	return []entity.Subtitle{}, nil
}

func (s *HasSubtitlesListState) GetSubtitlesById() (*entity.Subtitle, error) {
	return &entity.Subtitle{}, nil
}

func (s *HasSubtitlesListState) DeleteSubtitlesById() error {
	return nil
}

func (s *HasSubtitlesListState) GetRandomPhraseByCurrentSubtitles() (entity.Phrase, error) {
	return entity.Phrase{}, nil
}

func (s *HasSubtitlesListState) GetTranslateByPhraseId() (*entity.PhraseTranslation, error) {
	return &entity.PhraseTranslation{}, nil
}

func (s *HasSubtitlesListState) SetTranslateByPhraseId(translation string) error {
	return nil
}

func (s *HasSubtitlesListState) HideThisPhraseById() error {
	return nil
}
