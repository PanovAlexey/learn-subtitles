package bot_state_machine

import (
	"errors"
	"github.com/PanovAlexey/learn-subtitles/internal/domain/entity"
)

type HasRandomPhraseState struct {
	dialog *Dialog
}

func (s *HasRandomPhraseState) GetCode() StatesEnum {
	return HasRandomPhrase
}

func (s *HasRandomPhraseState) AddSubtitles() error {
	s.dialog.SetReadyToAddSubtitlesNameState()

	return nil
}

func (s *HasRandomPhraseState) AddSubtitlesName(name string) error {
	return errors.New("not available command for current state")
}

func (s *HasRandomPhraseState) AddSubtitlesText(text string) error {
	return errors.New("not available command for current state")
}

func (s *HasRandomPhraseState) AddForbiddenPartsAndSaveSubtitles(
	subtitles entity.Subtitle,
	forbiddenPartsString string,
) (*entity.Subtitle, error) {
	return nil, errors.New("not available command for current state")
}

func (s *HasRandomPhraseState) GetSubtitlesList() ([]entity.Subtitle, error) {
	return []entity.Subtitle{}, nil
}

func (s *HasRandomPhraseState) GetSubtitlesById() (*entity.Subtitle, error) {
	return &entity.Subtitle{}, nil
}

func (s *HasRandomPhraseState) DeleteSubtitlesById() error {
	return nil
}

func (s *HasRandomPhraseState) GetRandomPhraseBySubtitlesId() (*entity.Phrase, error) {
	return &entity.Phrase{}, nil
}

func (s *HasRandomPhraseState) GetTranslateByPhraseId() (*entity.PhraseTranslation, error) {
	return &entity.PhraseTranslation{}, nil
}

func (s *HasRandomPhraseState) SetTranslateByPhraseId(translation string) error {
	return nil
}

func (s *HasRandomPhraseState) HideThisPhraseById() error {
	return nil
}
