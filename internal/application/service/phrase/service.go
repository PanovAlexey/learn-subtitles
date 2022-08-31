package phrase

import "github.com/PanovAlexey/learn-subtitles/internal/domain/entity"

type PhraseService struct {
}

func NewPhraseService() PhraseService {
	return PhraseService{}
}

func (s PhraseService) GetRandomPhraseBySubtitleId(subtitleId int) *entity.Phrase {
	return nil
}

func (s PhraseService) GetTranslateByPhraseId() (*entity.PhraseTranslation, error) {
	return nil, nil
}

func (s PhraseService) SetTranslateByPhraseId() {

}

func (s PhraseService) HideThisPhraseById() {

}
