package repository

import "github.com/jmoiron/sqlx"

type SubtitleRepository struct {
	db *sqlx.DB
}

func NewSubtitleRepository(db *sqlx.DB) SubtitleRepository {
	return SubtitleRepository{db: db}
}
