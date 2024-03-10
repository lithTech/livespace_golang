package test

import (
	"context"
	"fmt"
	"livespace/assets"
	"livespace/src/config"
	con "livespace/src/db/service"
	"livespace/src/server"
	"livespace/src/test/testcontainers"
	"log/slog"
	"net/http"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/maxatome/go-testdeep/helpers/tdsuite"
	"github.com/maxatome/go-testdeep/td"
)

type IntegrationTest struct {
	Ctx      context.Context
	DB       *sqlx.DB
	Routes   http.Handler
	Server   *http.Server
	Postgres *testcontainers.PostgreSQLContainer
}

func (s *IntegrationTest) PreTest(t *td.T, testName string) error {
	slog.Debug(fmt.Sprint("Clearing test data before ", testName))
	tx := s.DB.MustBegin()
	_, err := s.DB.Exec("truncate planet")
	slog.Debug(fmt.Sprint("Error: ", err))
	return tx.Commit()
}

func (s *IntegrationTest) Setup(t *td.T) error {
	slog.Info("Setup tests")
	config.SetLogger([]string{"log.level", "debug"})
	ctx, cancelContext := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancelContext()

	postgres, err := testcontainers.NewPostgreSQLContainer(ctx)
	if err != nil {
		panic(err)
	}

	db, err := con.Connect(&config.DbConnectionConfig{Url: postgres.GetConnectUrl()})
	if err != nil {
		panic(err)
	}
	s.DB = db
	s.Ctx = context.Background()
	s.Postgres = postgres

	if err == nil {
		_, err = con.DoMigration(assets.GetAssets(), db)
	}

	server, routes := server.CreateServer(db)

	s.Routes = routes
	s.Server = server

	return err
}

func (s *IntegrationTest) Destroy(t *td.T) error {
	slog.Debug("Destroy tests")

	context, cancelContext := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelContext()

	s.Postgres.Terminate(context)

	return s.DB.Close()
}

func TestIntegration(t *testing.T) {
	slog.Debug("Running integration tests")
	tdsuite.Run(t, &IntegrationTest{})
}
