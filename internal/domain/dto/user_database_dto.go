package dto

import (
	"database/sql"
)

type UserDatabaseDto struct {
	Id        sql.NullInt64  `db:"id"`
	FirstName sql.NullString `db:"first_name"`
	LastName  sql.NullString `db:"last_name"`
	Login     sql.NullString `db:"login"`
	CreatedAt sql.NullTime   `db:"created_at"`
	IsDeleted sql.NullString `db:"is_deleted"`
}
