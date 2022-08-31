package subtitles

import (
	"github.com/PanovAlexey/learn-subtitles/internal/config"
	"github.com/PanovAlexey/learn-subtitles/internal/domain/entity"
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
