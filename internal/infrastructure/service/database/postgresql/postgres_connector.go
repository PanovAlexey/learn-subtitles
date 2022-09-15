package postgresql

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"time"
)

const TableUsersName = `users`
const TableSubtitlesName = `subtitles`

func GetPostgresConnector(
	databaseUser, databasePassword, databaseAddress, databasePort, databaseName string,
	maxOpenConnections, maxIdleConnections int,
	connectionMaxIdleTime, connectionMaxLifeTime time.Duration,
) (*sqlx.DB, error) {
	db, err := sqlx.Open(
		"pgx",
		"postgresql://"+databaseUser+":"+databasePassword+"@"+databaseAddress+
			":"+databasePort+"/"+databaseName+"?sslmode=disable",
	)

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConnections)
	db.SetMaxIdleConns(maxIdleConnections)
	db.SetConnMaxIdleTime(connectionMaxIdleTime)
	db.SetConnMaxLifetime(connectionMaxLifeTime)

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
