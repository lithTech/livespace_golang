package test

import (
	"fmt"
	"github.com/mariomac/gostream/stream"
	"livespace/src/space/api"
	dao "livespace/src/space/dao"
	planet "livespace/src/space/domain"
	"net/http"

	"github.com/maxatome/go-testdeep/helpers/tdhttp"
	"github.com/maxatome/go-testdeep/td"
)

func (s *IntegrationTest) TestGetSingle(t *td.T) {
	pl := &planet.Planet{Id: 0, Title: "Mercury", PlanetType: planet.AGRARIAN, Population: 12_122_244, Version: 1}
	err := dao.Save(s.Ctx, s.DB, pl)

	t.CmpNoError(err)

	testApi := tdhttp.NewTestAPI(t, s.Routes)

	testApi.Get(fmt.Sprint("/planets/", pl.Id)).
		CmpStatus(http.StatusOK).
		CmpJSONBody(api.ToView(pl))
}

func (s *IntegrationTest) TestGetList(t *td.T) {
	pl := planet.Planet{Id: 0, Title: "Mercury", PlanetType: planet.AGRARIAN, Population: 12_122_244, Version: 1}
	err := dao.Save(s.Ctx, s.DB, &pl)
	t.CmpNoError(err)

	pl2 := planet.Planet{Id: 0, Title: "Pluto", PlanetType: planet.MIXED, Population: 12_122_244, Version: 1}
	err = dao.Save(s.Ctx, s.DB, &pl2)
	t.CmpNoError(err)

	testApi := tdhttp.NewTestAPI(t, s.Routes)

	expected := stream.
		Map(stream.OfSlice([]planet.Planet{pl, pl2}), func(pl planet.Planet) *api.PlanetView {
			return api.ToView(&pl)
		}).
		ToSlice()

	testApi.Get("/planets").
		CmpStatus(http.StatusOK).
		CmpJSONBody(expected)
}

func (s *IntegrationTest) TestSave(t *td.T) {
	pl := planet.Planet{Id: 0, Title: "Mercury", PlanetType: planet.AGRARIAN, Population: 12_122_244, Version: 1}
	plView := api.ToView(&pl)
	testApi := tdhttp.NewTestAPI(t, s.Routes)

	testApi.PostJSON("/planets", plView).
		CmpStatus(http.StatusOK).
		CmpJSONBody(td.SStruct(
			api.ToView(&pl),
			td.StructFields{
				"Id": td.NotZero(),
			}))
}
