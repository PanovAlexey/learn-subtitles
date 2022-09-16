package phrase

import (
	"github.com/PanovAlexey/learn-subtitles/internal/domain/entity"
	"strings"
)

const maxPhraseLength = 100

type PhraseRepository interface {
	Add(text string, subtitleId int64) (entity.Phrase, error)
	GetById(id int64) (entity.Phrase, error)
	GetRandom(subtitleId, userId int64) (entity.Phrase, error)
}

type PhraseService struct {
	repository PhraseRepository
}

func NewPhraseService(repository PhraseRepository) PhraseService {
	return PhraseService{repository: repository}
}

func (s PhraseService) GetRandomPhraseBySubtitleId(subtitleId, userId int64) (entity.Phrase, error) {
	return s.repository.GetRandom(subtitleId, userId)
}

func (s PhraseService) GetTranslateByPhraseId() (*entity.PhraseTranslation, error) {
	return nil, nil
}

func (s PhraseService) SetTranslateByPhraseId() {

}

func (s PhraseService) HideThisPhraseById() {

}

func (s PhraseService) SaveTextInPhrases(subtitles entity.Subtitle) (error, int) {
	phrases := s.splitText(subtitles.Text.String)
	count := 0

	for _, v := range phrases {
		_, err := s.repository.Add(v, subtitles.Id.Int64)

		if err != nil {
			return err, count
		}

		count++
	}

	return nil, count
}

func (s PhraseService) splitText(text string) []string {
	var phrases []string
	text = s.clearText(text)

	phrases = s.splitTextIntoSentences(text)
	phrases = s.splitPhrasesByColon(phrases)
	phrases = s.splitPhrasesByDash(phrases)
	phrases = s.splitPhrasesByCommas(phrases)
	phrases = s.splitPhrasesBySpace(phrases)
	phrases = s.clearOfEmpty(phrases)

	return phrases
}

func (s PhraseService) clearText(text string) string {
	text = strings.TrimSpace(text)
	text = strings.Replace(text, "\n", " ", -1)
	text = strings.Replace(text, "\t", " ", -1)
	text = strings.Replace(text, "\r", " ", -1)
	text = strings.Replace(text, "...", ".", -1)
	text = strings.Replace(text, "..", ".", -1)
	text = strings.Replace(text, "  ", " ", -1)
	text = strings.Replace(text, "  ", " ", -1)
	text = strings.Replace(text, ". ", ".", -1)
	text = strings.Replace(text, "? ", "?", -1)
	text = strings.Replace(text, "! ", "!", -1)
	text = strings.Replace(text, "! ", "!", -1)
	text = strings.Replace(text, " - ", " ", -1)
	text = strings.Replace(text, " -", "", -1)
	text = strings.Replace(text, "‐ ", "", -1)
	text = strings.Replace(text, "  ", " ", -1)

	return text
}

func (s PhraseService) splitTextIntoSentences(text string) []string {
	sentences := s.splitPhrasesBySymbol([]string{text}, ".")
	sentences = s.splitPhrasesBySymbol(sentences, "!")
	sentences = s.splitPhrasesBySymbol(sentences, "?")

	return sentences
}

func (s PhraseService) splitPhrasesByCommas(phrases []string) []string {
	return s.splitPhrasesBySymbol(phrases, ",")
}

func (s PhraseService) splitPhrasesByDash(phrases []string) []string {
	return s.splitPhrasesBySymbol(phrases, "—")
}

func (s PhraseService) splitPhrasesByColon(phrases []string) []string {
	return s.splitPhrasesBySymbol(phrases, ":")
}

func (s PhraseService) splitPhrasesBySpace(phrases []string) []string {
	return s.splitPhrasesBySymbol(phrases, " ")
}

func (s PhraseService) splitPhrasesBySymbol(phrases []string, symbol string) []string {
	var outputPhrases []string

	for _, phrase := range phrases {
		if len(phrase) > maxPhraseLength {
			temporaryLine := ""
			phrasesBySymbol := strings.Split(phrase, symbol)

			for i := 0; i < len(phrasesBySymbol); i++ {
				if i < len(phrasesBySymbol)-1 {
					phrasesBySymbol[i] = phrasesBySymbol[i] + symbol
				}

				if len(temporaryLine)+len(phrasesBySymbol[i]) < maxPhraseLength {
					temporaryLine = temporaryLine + phrasesBySymbol[i]
				} else {
					outputPhrases = append(outputPhrases, strings.TrimSpace(temporaryLine))
					temporaryLine = phrasesBySymbol[i]
				}
			}

			if len(temporaryLine) > 0 {
				outputPhrases = append(outputPhrases, strings.TrimSpace(temporaryLine))
			}
		} else {
			outputPhrases = append(outputPhrases, strings.TrimSpace(phrase))
		}
	}

	return outputPhrases
}

func (s PhraseService) clearOfEmpty(phrases []string) []string {
	var outputPhrases []string

	for _, phrase := range phrases {
		if len(phrase) > 0 {
			outputPhrases = append(outputPhrases, strings.TrimSpace(phrase))
		}
	}

	return outputPhrases
}
