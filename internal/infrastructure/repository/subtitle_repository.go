package repository

import (
	"database/sql"
	"github.com/PanovAlexey/learn-subtitles/internal/domain/dto"
	"github.com/PanovAlexey/learn-subtitles/internal/domain/entity"
	"github.com/PanovAlexey/learn-subtitles/internal/infrastructure/service/database/postgresql"
	"github.com/jmoiron/sqlx"
	"time"
)

type SubtitleRepository struct {
	db *sqlx.DB
}

func NewSubtitleRepository(db *sqlx.DB) SubtitleRepository {
	return SubtitleRepository{db: db}
}

func (r SubtitleRepository) Add(subtitles entity.Subtitle, userId int64) (dto.SubtitleDatabaseDto, error) {
	subtitleDatabaseDto := dto.SubtitleDatabaseDto{}
	query := "INSERT INTO " +
		postgresql.TableSubtitlesName +
		" (name, text, user_id, created_at, is_deleted) VALUES" +
		" ($1, $2, $3, $4, $5) RETURNING id, name, text, user_id, created_at, is_deleted"
	err := r.db.QueryRow(
		query,
		subtitles.Name,
		subtitles.Text,
		userId,
		time.Now(),
		false,
	).Scan(
		&subtitleDatabaseDto.Id,
		&subtitleDatabaseDto.Name,
		&subtitleDatabaseDto.Text,
		&subtitleDatabaseDto.Author,
		&subtitleDatabaseDto.CreatedAt,
		&subtitleDatabaseDto.IsDeleted,
	)

	if err != nil {
		return subtitleDatabaseDto, err
	}

	return subtitleDatabaseDto, err
}

func (r SubtitleRepository) GetList(userId int) ([]dto.SubtitleDatabaseDto, error) {
	var collection []dto.SubtitleDatabaseDto
	subtitleDatabaseDto := dto.SubtitleDatabaseDto{}

	query := "SELECT id, name, text, created_at, is_deleted FROM " + postgresql.TableSubtitlesName + " WHERE user_id = $1"
	rows, err := r.db.Query(query, userId)

	if err != nil || rows.Err() != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&subtitleDatabaseDto.Id,
			&subtitleDatabaseDto.Name,
			&subtitleDatabaseDto.Text,
			&subtitleDatabaseDto.CreatedAt,
			&subtitleDatabaseDto.IsDeleted,
		); err != nil {
			return collection, err
		}

		collection = append(collection, subtitleDatabaseDto)
	}

	return collection, nil
}

func (r SubtitleRepository) GetById(id, userId int) (dto.SubtitleDatabaseDto, error) {
	var subtitleDatabaseDto dto.SubtitleDatabaseDto

	err := r.db.Get(
		&subtitleDatabaseDto,
		"SELECT id, name, text, user_id, created_at, is_deleted FROM "+
			postgresql.TableSubtitlesName+" WHERE id = $1 and user_id = $2 LIMIT 1",
		id,
		userId,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return subtitleDatabaseDto, nil
		}

		return subtitleDatabaseDto, err
	}

	return subtitleDatabaseDto, err
}

func (r SubtitleRepository) Delete(id, userId int) error {
	result := ""
	err := r.db.Get(
		&result,
		"DELETE FROM "+postgresql.TableSubtitlesName+" WHERE id = $1 AND user_id = $2",
		id,
		userId,
	)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

func (r SubtitleRepository) Update() (dto.SubtitleDatabaseDto, error) {
	subtitleDatabaseDto := dto.SubtitleDatabaseDto{}
	query := "UPDATE " +
		postgresql.TableSubtitlesName +
		" SET name=$1, text=$2, user_id=$3, created_at=$4, is_deleted=$5" +
		" WHERE id=$6 AND user_id=$7" +
		" RETURNING id, name, text, user_id, created_at, is_deleted"
	err := r.db.QueryRow(
		query,
		subtitleDatabaseDto.Name,
		subtitleDatabaseDto.Text,
		subtitleDatabaseDto.Author,
		subtitleDatabaseDto.CreatedAt.Time.Format(time.RFC3339),
		subtitleDatabaseDto.IsDeleted,
	).
		Scan(
			&subtitleDatabaseDto.Id,
			&subtitleDatabaseDto.Name,
			&subtitleDatabaseDto.Text,
			&subtitleDatabaseDto.Author,
			&subtitleDatabaseDto.CreatedAt,
			&subtitleDatabaseDto.IsDeleted,
		)

	if err != nil {
		return subtitleDatabaseDto, err
	}

	return subtitleDatabaseDto, err
}
