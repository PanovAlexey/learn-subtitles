package bot_state_machine

import (
	"errors"
	"github.com/PanovAlexey/learn-subtitles/internal/domain/entity"
)

type ReadyToAddTranslationRandomPhraseState struct {
	dialog *Dialog
}

func (s *ReadyToAddTranslationRandomPhraseState) GetCode() StatesEnum {
	return ReadyToAddTranslationRandomPhrase
}

func (s *ReadyToAddTranslationRandomPhraseState) AddSubtitles() error {
	s.dialog.SetReadyToAddSubtitlesNameState()

	return nil
}

func (s *ReadyToAddTranslationRandomPhraseState) AddSubtitlesName(name string) error {
	return errors.New("not available command for current state")
}

func (s *ReadyToAddTranslationRandomPhraseState) AddSubtitlesText(text string) error {
	return errors.New("not available command for current state")
}

func (s *ReadyToAddTranslationRandomPhraseState) AddForbiddenPartsAndSaveSubtitles(
	subtitles entity.Subtitle,
	forbiddenPartsString string,
) (*entity.Subtitle, error) {
	return nil, errors.New("not available command for current state")
}

func (s *ReadyToAddTranslationRandomPhraseState) GetSubtitlesList() ([]entity.Subtitle, error) {
	return []entity.Subtitle{}, nil
}

func (s *ReadyToAddTranslationRandomPhraseState) GetSubtitlesById() (*entity.Subtitle, error) {
	return &entity.Subtitle{}, nil
}

func (s *ReadyToAddTranslationRandomPhraseState) DeleteSubtitlesById() error {
	return nil
}

func (s *ReadyToAddTranslationRandomPhraseState) GetRandomPhraseByCurrentSubtitles() (entity.Phrase, error) {
	return entity.Phrase{}, nil
}

func (s *ReadyToAddTranslationRandomPhraseState) GetTranslateByPhraseId() (*entity.PhraseTranslation, error) {
	return &entity.PhraseTranslation{}, nil
}

func (s *ReadyToAddTranslationRandomPhraseState) SetTranslateByPhraseId(translation string) error {
	return nil
}

func (s *ReadyToAddTranslationRandomPhraseState) HideThisPhraseById() error {
	return nil
}
