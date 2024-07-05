package database

import (
	"database/sql"
	"os"

	"github.com/3tagger/echo-sample-arch/internal/config"
)

func InitPostgreSQL(dbCfg config.PostgreSQLConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
