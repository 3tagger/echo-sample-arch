package database

import (
	"database/sql"
	"fmt"

	"github.com/3tagger/echo-sample-arch/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitPostgreSQL(dbCfg config.PostgreSQLConfig) (*sql.DB, error) {
	dburl := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s",
		dbCfg.Host,
		dbCfg.Port,
		dbCfg.Name,
		dbCfg.Username,
		dbCfg.Password,
	)

	db, err := sql.Open("pgx", dburl)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
