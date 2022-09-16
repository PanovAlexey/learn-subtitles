package repository

import (
	"database/sql"
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

func (r SubtitleRepository) Add(subtitles entity.Subtitle) (entity.Subtitle, error) {
	outSubtitles := entity.Subtitle{}
	query := "INSERT INTO " +
		postgresql.TableSubtitlesName +
		" (name, text, user_id, created_at, is_deleted) VALUES" +
		" ($1, $2, $3, $4, $5) RETURNING id, name, text, user_id, created_at, is_deleted"
	err := r.db.QueryRow(
		query,
		subtitles.Name.String,
		subtitles.Text.String,
		subtitles.Author.Int64,
		time.Now(),
		false,
	).Scan(
		&outSubtitles.Id,
		&outSubtitles.Name,
		&outSubtitles.Text,
		&outSubtitles.Author,
		&outSubtitles.CreatedAt,
		&outSubtitles.IsDeleted,
	)

	if err != nil {
		return outSubtitles, err
	}

	return outSubtitles, err
}

func (r SubtitleRepository) GetList(userId int) ([]entity.Subtitle, error) {
	var collection []entity.Subtitle
	outSubtitle := entity.Subtitle{}

	query := "SELECT id, name, text, created_at, is_deleted FROM " + postgresql.TableSubtitlesName + " WHERE user_id = $1"
	rows, err := r.db.Query(query, userId)

	if err != nil || rows.Err() != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&outSubtitle.Id,
			&outSubtitle.Name,
			&outSubtitle.Text,
			&outSubtitle.CreatedAt,
			&outSubtitle.IsDeleted,
		); err != nil {
			return collection, err
		}

		collection = append(collection, outSubtitle)
	}

	return collection, nil
}

func (r SubtitleRepository) GetById(id, userId int) (entity.Subtitle, error) {
	var outSubtitle entity.Subtitle

	err := r.db.Get(
		&outSubtitle,
		"SELECT id, name, text, user_id, created_at, is_deleted FROM "+
			postgresql.TableSubtitlesName+" WHERE id = $1 and user_id = $2 LIMIT 1",
		id,
		userId,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return outSubtitle, nil
		}

		return outSubtitle, err
	}

	return outSubtitle, err
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

func (r SubtitleRepository) Update(subtitle entity.Subtitle) (entity.Subtitle, error) {
	outSubtitle := entity.Subtitle{}
	query := "UPDATE " +
		postgresql.TableSubtitlesName +
		" SET name=$1, text=$2, user_id=$3, created_at=$4, is_deleted=$5" +
		" WHERE id=$6 AND user_id=$7" +
		" RETURNING id, name, text, user_id, created_at, is_deleted"
	err := r.db.QueryRow(
		query,
		subtitle.Name.String,
		subtitle.Text.String,
		subtitle.Author.Int64,
		subtitle.CreatedAt.Time.Format(time.RFC3339),
		subtitle.IsDeleted.Bool,
	).
		Scan(
			&outSubtitle.Id,
			&outSubtitle.Name,
			&outSubtitle.Text,
			&outSubtitle.Author,
			&outSubtitle.CreatedAt,
			&outSubtitle.IsDeleted,
		)

	if err != nil {
		return outSubtitle, err
	}

	return outSubtitle, err
}
