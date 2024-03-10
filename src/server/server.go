package server

import (
	spaceApi "livespace/src/space/api"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
)

func CreateServer(db *sqlx.DB) (*http.Server, http.Handler) {
	routes := httprouter.New()

	planetHandler := spaceApi.PlanetHandler{DB: db}
	routes.GET("/planets/:id", planetHandler.GetPlanetHandler)
	routes.GET("/planets", planetHandler.GetPlanetsHandler)
	routes.POST("/planets", planetHandler.SaveHandler)

	server := http.Server{Addr: ":8080", Handler: routes}

	return &server, routes
}
