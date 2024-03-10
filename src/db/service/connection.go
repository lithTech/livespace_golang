package db

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"livespace/src/config"
	log "log/slog"
)

func Connect(connectionConfig *config.DbConnectionConfig) (db *sqlx.DB, err error) {
	log.Info("Connecting to db " + connectionConfig.Url)

	db, err = sqlx.Connect("pgx", connectionConfig.Url)
	if err != nil {
		return nil, err
	}

	var version []string
	err = db.Select(&version, "select version();")
	if err != nil {
		log.Error("Failed to get db version", err)
	} else {
		log.Info("Db version is " + version[0])
	}

	return db, err
}
