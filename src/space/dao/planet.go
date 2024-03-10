package planet

import (
	"context"
	"fmt"
	domainDb "livespace/src/db"
	planet "livespace/src/space/domain"
	"log/slog"

	"github.com/jmoiron/sqlx"
)

func Get(ctx context.Context, db *sqlx.DB, id int64) (res *planet.Planet, err error) {
	slog.Debug(fmt.Sprint("Getting planet with id ", id))
	res = &planet.Planet{}
	err = db.GetContext(ctx, res, "select * from planet where id = $1", id)
	if err != nil {
		res = nil
	}
	return res, err
}

func GetAll(ctx context.Context, db *sqlx.DB) ([]planet.Planet, error) {
	slog.Debug(fmt.Sprint("Getting planets"))
	var planets = []planet.Planet{}
	err := db.SelectContext(ctx, &planets, "select * from planet")
	if err != nil {
		planets = nil
	}

	return planets, err
}

func Save(ctx context.Context, db *sqlx.DB, planet *planet.Planet) (err error) {
	tx := db.MustBegin()
	defer func() {
		if err != nil {
			slog.Debug("Rolling back " + err.Error())
			tx.Rollback()
		} else {
			slog.Debug("Commit!")
			err = tx.Commit()
		}
	}()

	if planet.Id == 0 {
		slog.Debug("Inserting planet " + planet.Title)
		var newId int64
		query := "insert into planet (title, population, planet_type) " +
			"values (:title, :population, :planet_type) returning id"
		named, err := tx.PrepareNamed(query)
		err = named.GetContext(ctx, &newId, &planet)
		if err != nil {
			return err
		}
		planet.Id = newId
		planet.Version = 1
	} else {
		slog.Debug(fmt.Sprint("Update planet ", planet))
		query := "update planet set " +
			"title = :title, population = :population, planet_type = :planet_type, version = version + 1 " +
			"where id = :id and version = :version"
		res, err := tx.NamedExecContext(ctx, query, &planet)
		if err != nil {
			return err
		}
		cnt, err := res.RowsAffected()
		slog.Debug(fmt.Sprint("count is ", cnt))
		if cnt == 0 {
			err = &domainDb.ConcurrentModificationError{}
		} else {
			planet.Version++
		}
	}
	return nil
}
