package db

import (
	"livespace/assets"
	"log/slog"

	sql "github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"
)

func DoMigration(assets *assets.Assets, db *sql.DB) (int, error) {
	slog.Info("Executing db migration scripts")
	migration := migrate.EmbedFileSystemMigrationSource{FileSystem: *&assets.Sql, Root: "sql"}

	r, err := migrate.Exec(db.DB, "postgres", migration, migrate.Up)
	if err != nil {
		return 0, err
	}
	return r, nil
}
