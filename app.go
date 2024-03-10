package main

import (
	"livespace/assets"
	"livespace/src/config"
	logger "livespace/src/config"
	connection "livespace/src/db/service"
	"livespace/src/server"
	log "log/slog"
	"os"
)

func Run() {
	logger.SetLogger(os.Args)
	log.Info("Starting application...")
	assets := assets.GetAssets()
	applicationConfig := config.LoadConfig(assets)
	connectionPool, err := connection.Connect(&applicationConfig.Database.Connection)
	if err != nil {
		panic(err)
	}

	defer connectionPool.Close()
	if err == nil && connectionPool != nil {
		log.Info("Successfully obtained connection")
		_, err := connection.DoMigration(assets, connectionPool)
		if err != nil {
			log.Error("Failed to do sql migration", err)
			return
		}
	}

	server, _ := server.CreateServer(connectionPool)
	server.ListenAndServe()
}
