package repository

import (
	"github.com/PanovAlexey/learn-subtitles/internal/domain/entity"
	"github.com/PanovAlexey/learn-subtitles/internal/infrastructure/service/database/postgresql"
	"github.com/jmoiron/sqlx"
	"time"
)

type PhraseRepository struct {
	db *sqlx.DB
}

func NewPhraseRepository(db *sqlx.DB) PhraseRepository {
	return PhraseRepository{db: db}
}

func (r PhraseRepository) Add(text string, subtitleId int64) (entity.Phrase, error) {
	phrase := entity.Phrase{}
	query := "INSERT INTO " +
		postgresql.TablePhrasesName +
		" (text, subtitle_id, created_at, is_deleted) VALUES" +
		" ($1, $2, $3, $4) RETURNING id, text, subtitle_id, created_at, is_deleted"
	err := r.db.QueryRow(
		query,
		text,
		subtitleId,
		time.Now(),
		false,
	).Scan(
		&phrase.Id,
		&phrase.Text,
		&phrase.Subtitle,
		&phrase.CreatedAt,
		&phrase.IsDeleted,
	)

	if err != nil {
		return phrase, err
	}

	return phrase, err
}

func (r PhraseRepository) GetById(id int64) (entity.Phrase, error) {
	return entity.Phrase{}, nil
}

func (r PhraseRepository) GetRandom(id, subtitleId int64) (entity.Phrase, error) {
	return entity.Phrase{}, nil
}
