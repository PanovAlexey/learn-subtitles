package repository

import (
	"database/sql"
	"github.com/PanovAlexey/learn-subtitles/internal/domain/dto"
	"github.com/PanovAlexey/learn-subtitles/internal/domain/entity"
	"github.com/PanovAlexey/learn-subtitles/internal/infrastructure/service/database/postgresql"
	"github.com/jmoiron/sqlx"
	"time"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return UserRepository{db: db}
}

func (r UserRepository) SaveUser(user entity.User) (dto.UserDatabaseDto, error) {
	userDatabaseDto := dto.UserDatabaseDto{}
	query := "INSERT INTO " +
		postgresql.TableUsersName +
		" (first_name, last_name, created_at, login) VALUES ($1, $2, $3, $4) RETURNING id, first_name, last_name, created_at, login"
	err := r.db.QueryRow(query, user.FirstName, user.LastName, time.Now(), user.Login).
		Scan(
			&userDatabaseDto.Id,
			&userDatabaseDto.FirstName,
			&userDatabaseDto.LastName,
			&userDatabaseDto.CreatedAt,
			&userDatabaseDto.Login,
		)

	if err != nil {
		return userDatabaseDto, err
	}

	return userDatabaseDto, err
}

func (r UserRepository) GetUserByLogin(login string) (dto.UserDatabaseDto, error) {
	userDatabaseDto := dto.UserDatabaseDto{}

	err := r.db.Get(
		&userDatabaseDto,
		"SELECT * FROM "+postgresql.TableUsersName+" WHERE login = $1 LIMIT 1",
		login,
	)

	if err == sql.ErrNoRows {
		return userDatabaseDto, nil
	}

	if err != nil {
		return userDatabaseDto, err
	}

	return userDatabaseDto, err
}
