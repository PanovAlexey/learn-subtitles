package bot_state_machine

import (
	"errors"
	"github.com/PanovAlexey/learn-subtitles/internal/domain/entity"
)

type ReadyToAddSubtitlesNameState struct {
	dialog *Dialog
}

func (s *ReadyToAddSubtitlesNameState) GetCode() StatesEnum {
	return ReadyToAddSubtitlesName
}

func (s *ReadyToAddSubtitlesNameState) AddSubtitles() error {
	s.dialog.SetReadyToAddSubtitlesNameState()

	return nil
}

func (s *ReadyToAddSubtitlesNameState) AddSubtitlesName(name string) error {
	err := s.dialog.subtitlesService.ValidateName(name)

	if err != nil {
		return err
	}

	s.dialog.SetReadyToAddSubtitlesTextState()

	return nil
}

func (s *ReadyToAddSubtitlesNameState) AddSubtitlesText(text string) error {
	return errors.New("not available command for current state")
}

func (s *ReadyToAddSubtitlesNameState) AddForbiddenPartsAndSaveSubtitles(subtitles entity.Subtitle, forbiddenPartsString string) error {
	return errors.New("not available command for current state")
}

func (s *ReadyToAddSubtitlesNameState) GetSubtitlesList() ([]entity.Subtitle, error) {
	return []entity.Subtitle{}, nil
}

func (s *ReadyToAddSubtitlesNameState) GetSubtitlesById() (*entity.Subtitle, error) {
	return &entity.Subtitle{}, nil
}

func (s *ReadyToAddSubtitlesNameState) DeleteSubtitlesById() error {
	return nil
}

func (s *ReadyToAddSubtitlesNameState) GetRandomPhraseBySubtitlesId() (*entity.Phrase, error) {
	return &entity.Phrase{}, nil
}

func (s *ReadyToAddSubtitlesNameState) GetTranslateByPhraseId() (*entity.PhraseTranslation, error) {
	return &entity.PhraseTranslation{}, nil
}

func (s *ReadyToAddSubtitlesNameState) SetTranslateByPhraseId(translation string) error {
	return nil
}

func (s *ReadyToAddSubtitlesNameState) HideThisPhraseById() error {
	return nil
}
