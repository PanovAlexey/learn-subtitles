package subtitles

import (
	"fmt"
	customErrors "github.com/PanovAlexey/learn-subtitles/internal/application/errors"
	"github.com/PanovAlexey/learn-subtitles/internal/config"
	"github.com/PanovAlexey/learn-subtitles/internal/domain/entity"
	"log"
	"strings"
)

type SubtitlesService struct {
	config config.Config
}

func NewSubtitlesService(config config.Config) SubtitlesService {
	return SubtitlesService{
		config: config,
	}
}

func (s SubtitlesService) Add(userId int, subtitle string, forbiddenParts []string) error {
	return nil
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