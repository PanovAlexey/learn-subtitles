package bot_state_machine

import (
	"errors"
	"github.com/PanovAlexey/learn-subtitles/internal/domain/entity"
)

type ReadyToAddSubtitlesTextState struct {
	dialog *Dialog
}

func (s *ReadyToAddSubtitlesTextState) GetCode() StatesEnum {
	return ReadyToAddSubtitlesText
}

func (s *ReadyToAddSubtitlesTextState) AddSubtitles() error {
	s.dialog.SetReadyToAddSubtitlesNameState()

	return nil
}

func (s *ReadyToAddSubtitlesTextState) AddSubtitlesName(name string) error {
	return errors.New("not available command for current state")
}

func (s *ReadyToAddSubtitlesTextState) AddSubtitlesText(text string) error {
	err := s.dialog.subtitlesService.ValidateText(text)

	if err != nil {
		return err
	}

	s.dialog.SetReadyToAddSubtitlesProhibitedWordsState()

	return nil
}

func (s *ReadyToAddSubtitlesTextState) AddForbiddenPartsAndSaveSubtitles(subtitles entity.Subtitle, forbiddenPartsString string) error {
	return errors.New("not available command for current state")
}

func (s *ReadyToAddSubtitlesTextState) GetSubtitlesList() ([]entity.Subtitle, error) {
	return []entity.Subtitle{}, nil
}

func (s *ReadyToAddSubtitlesTextState) GetSubtitlesById() (*entity.Subtitle, error) {
	return &entity.Subtitle{}, nil
}

func (s *ReadyToAddSubtitlesTextState) DeleteSubtitlesById() error {
	return nil
}

func (s *ReadyToAddSubtitlesTextState) GetRandomPhraseBySubtitlesId() (*entity.Phrase, error) {
	return &entity.Phrase{}, nil
}

func (s *ReadyToAddSubtitlesTextState) GetTranslateByPhraseId() (*entity.PhraseTranslation, error) {
	return &entity.PhraseTranslation{}, nil
}

func (s *ReadyToAddSubtitlesTextState) SetTranslateByPhraseId(translation string) error {
	return nil
}

func (s *ReadyToAddSubtitlesTextState) HideThisPhraseById() error {
	return nil
}
