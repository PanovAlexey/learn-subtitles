package bot_state_machine

import (
	"errors"
	"github.com/PanovAlexey/learn-subtitles/internal/domain/entity"
)

type ReadyToAddSubtitlesProhibitedWordsState struct {
	dialog *Dialog
}

func (s *ReadyToAddSubtitlesProhibitedWordsState) GetCode() StatesEnum {
	return ReadyToAddSubtitlesProhibitedWords
}

func (s *ReadyToAddSubtitlesProhibitedWordsState) AddSubtitles() error {
	s.dialog.SetReadyToAddSubtitlesNameState()

	return nil
}

func (s *ReadyToAddSubtitlesProhibitedWordsState) AddSubtitlesName(name string) error {
	return errors.New("not available command for current state")
}

func (s *ReadyToAddSubtitlesProhibitedWordsState) AddSubtitlesText(text string) error {
	return errors.New("not available command for current state")
}

func (s *ReadyToAddSubtitlesProhibitedWordsState) AddForbiddenPartsAndSaveSubtitles(subtitles entity.Subtitle, forbiddenPartsString string) error {
	subtitles.ForbiddenParts = s.dialog.subtitlesService.GetForbiddenPartsMapByString(forbiddenPartsString)
	err := s.dialog.subtitlesService.Add(777, subtitles) //@ToDo: change for real data

	if err != nil {
		return err
	}

	s.dialog.SetRestState()

	return nil
}

func (s *ReadyToAddSubtitlesProhibitedWordsState) GetSubtitlesList() ([]entity.Subtitle, error) {
	return []entity.Subtitle{}, nil
}

func (s *ReadyToAddSubtitlesProhibitedWordsState) GetSubtitlesById() (*entity.Subtitle, error) {
	return &entity.Subtitle{}, nil
}

func (s *ReadyToAddSubtitlesProhibitedWordsState) DeleteSubtitlesById() error {
	return nil
}

func (s *ReadyToAddSubtitlesProhibitedWordsState) GetRandomPhraseBySubtitlesId() (*entity.Phrase, error) {
	return &entity.Phrase{}, nil
}

func (s *ReadyToAddSubtitlesProhibitedWordsState) GetTranslateByPhraseId() (*entity.PhraseTranslation, error) {
	return &entity.PhraseTranslation{}, nil
}

func (s *ReadyToAddSubtitlesProhibitedWordsState) SetTranslateByPhraseId(translation string) error {
	return nil
}

func (s *ReadyToAddSubtitlesProhibitedWordsState) HideThisPhraseById() error {
	return nil
}
