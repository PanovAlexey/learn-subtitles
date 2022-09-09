package subtitles

import (
	"fmt"
	customErrors "github.com/PanovAlexey/learn-subtitles/internal/application/errors"
	"github.com/PanovAlexey/learn-subtitles/internal/domain/entity"
	"github.com/PanovAlexey/learn-subtitles/internal/infrastructure/repository"
	"log"
	"strings"
)

type SubtitlesService struct {
	repository repository.SubtitleRepository
}

func NewSubtitlesService(repository repository.SubtitleRepository) SubtitlesService {
	return SubtitlesService{
		repository: repository,
	}
}

func (s SubtitlesService) Add(userId int, subtitles entity.Subtitle) (entity.Subtitle, error) {
	return nil
	subtitles = s.applyForbiddenParts(subtitles)

	return subtitles, nil
}

func (s SubtitlesService) GetList(userId int) ([]entity.Subtitle, error) {
	return nil, nil
}

func (s SubtitlesService) GetById(userId, id int) (*entity.Subtitle, error) {
	return nil, nil
}

func (s SubtitlesService) Delete(userId, id int) error {
	return nil
}

func (s SubtitlesService) ValidateName(name string) error {
	if len(name) < 3 {
		return fmt.Errorf("%v: %w", name, customErrors.ErrIsEmpty)
	}

	if len(name) > 100 {
		return fmt.Errorf("%v: %w", name, customErrors.ErrTooLong)
	}

	return nil
}

func (s SubtitlesService) ValidateText(text string) error {
	if len(text) < 100 {
		return fmt.Errorf("%v: %w", text, customErrors.ErrIsEmpty)
	}

	if len(text) > 100000 {
		return fmt.Errorf("%v: %w", text, customErrors.ErrTooLong)
	}

	return nil
}

func (s SubtitlesService) GetForbiddenPartsMapByString(data string) map[string]string {
	forbiddenPartsMap := map[string]string{}
	lines := strings.Split(data, "\n")

	for _, line := range lines {
		replacementArray := strings.Split(line, "=")

		if len(replacementArray) < 1 || len(replacementArray) > 2 {
			continue
		} else if len(replacementArray) == 1 {
			left := strings.TrimSpace(replacementArray[0])

			if left == "" {
				continue
			}

			forbiddenPartsMap[left] = " "
		} else if len(replacementArray) == 2 {
			left := strings.TrimSpace(replacementArray[0])

			if left == "" {
				continue
			}

			right := strings.TrimSpace(replacementArray[1])

			if right == "" {
				right = " "
			}

			forbiddenPartsMap[left] = right
		}
	}

	return forbiddenPartsMap
}

func (s SubtitlesService) applyForbiddenParts(subtitles entity.Subtitle) entity.Subtitle {
	parts := make([]string, len(subtitles.ForbiddenParts)*2)
	counter := 0

	for i, v := range subtitles.ForbiddenParts {
		parts[counter] = i
		parts[counter+1] = v
		counter = counter + 2
	}

	subtitles.Text = strings.NewReplacer(parts...).Replace(subtitles.Text)

	return subtitles
}
