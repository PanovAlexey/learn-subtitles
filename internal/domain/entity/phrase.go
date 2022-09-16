package entity

import (
	"database/sql"
)

type Phrase struct {
	Id        sql.NullInt64  `db:"id"`
	Text      sql.NullString `db:"text"`
	CreatedAt sql.NullTime   `db:"created_at"`
	Subtitle  sql.NullInt64  `db:"subtitle_id"`
	IsDeleted sql.NullBool   `db:"is_deleted"`
}
