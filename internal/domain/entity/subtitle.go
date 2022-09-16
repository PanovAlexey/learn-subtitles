package entity

import (
	"database/sql"
)

type Subtitle struct {
	Id             sql.NullInt64  `db:"id"`
	Name           sql.NullString `db:"name"`
	Text           sql.NullString `db:"text"`
	CreatedAt      sql.NullTime   `db:"created_at"`
	Author         sql.NullInt64  `db:"user_id"`
	IsDeleted      sql.NullBool   `db:"is_deleted"`
	ForbiddenParts map[string]string
}
