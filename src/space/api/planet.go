package api

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"livespace/src/helpers"
	dao "livespace/src/space/dao"
	_ "livespace/src/space/domain"
	planet "livespace/src/space/domain"

	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	"github.com/mariomac/gostream/stream"
)

type PlanetHandler struct {
	DB *sqlx.DB
}

type Type uint8

const ( // PlanetType
	NON_HABITAT = 1 + iota
	AGRARIAN
	SCIENTIFIC
	MIXED
	PRIMITIVE
)

type PlanetView struct {
	Id         int64  `json:"id"`
	Title      string `json:"title"`
	Type       Type   `json:"planet_type"`
	Population uint64 `json:"population"`
	Version    uint16 `json:"version"`
}

func (p *PlanetView) ToDomain() *planet.Planet {
	return &planet.Planet{
		Id:         p.Id,
		Title:      p.Title,
		PlanetType: planet.Type(p.Type),
		Population: p.Population,
		Version:    p.Version,
	}
}

func ToView(p *planet.Planet) *PlanetView {
	return &PlanetView{
		Id:         p.Id,
		Title:      p.Title,
		Type:       Type(p.PlanetType),
		Population: p.Population,
		Version:    p.Version,
	}
}

func (h *PlanetHandler) GetPlanetHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()
	slog.Info(fmt.Sprint("Get planet Handling url ", r.URL, " id ", params.ByName("id")))

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil {
		w.Write(helpers.ToError(err))
		return
	}

	res, err := dao.Get(r.Context(), h.DB, id)
	helpers.WriteResponse(w, ToView(res), err)
}

func (h *PlanetHandler) GetPlanetsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer r.Body.Close()
	slog.Info(fmt.Sprint("Get planets Handling url ", r.URL))

	list, err := dao.GetAll(r.Context(), h.DB)
	res := stream.
		Map(stream.OfSlice(list), func(p planet.Planet) *PlanetView {
			return ToView(&p)
		}).
		ToSlice()
	helpers.WriteResponse(w, res, err)
}

func (h *PlanetHandler) SaveHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer r.Body.Close()
	slog.Info(fmt.Sprint("Save planet Handling url ", r.URL))

	var plView PlanetView
	err := helpers.DecodeJSONBody(w, r, &plView)
	if err != nil {
		w.Write(helpers.ToError(err))
		return
	}

	pl := plView.ToDomain()
	err = dao.Save(r.Context(), h.DB, pl)
	helpers.WriteResponse(w, ToView(pl), err)
}
