package test

import (
	dao "livespace/src/space/dao"
	planet "livespace/src/space/domain"

	"github.com/maxatome/go-testdeep/td"
)

func (s *IntegrationTest) TestInsertAndGet(t *td.T) {
	pl := planet.Planet{Id: 0, Title: "Mercury", PlanetType: planet.AGRARIAN, Population: 12_122_244, Version: 1}
	err := dao.Save(s.Ctx, s.DB, &pl)

	t.CmpNoError(err)

	got, err := dao.Get(s.Ctx, s.DB, pl.Id)
	if t.CmpNoError(err) {
		t.Cmp(*got, pl)
	}
}

func (s *IntegrationTest) TestUpdateAndGet(t *td.T) {
	pl := planet.Planet{Id: 0, Title: "Mercury", PlanetType: planet.AGRARIAN, Population: 12_122_244, Version: 1}
	dao.Save(s.Ctx, s.DB, &pl)

	pl.Title = "Changed mercury"
	dao.Save(s.Ctx, s.DB, &pl)

	got, err := dao.Get(s.Ctx, s.DB, pl.Id)
	if t.CmpNoError(err) {
		t.Cmp(*got, pl)
	}
}

func (s *IntegrationTest) TestGetAll(t *td.T) {
	pl := planet.Planet{Id: 0, Title: "Mercury", PlanetType: planet.AGRARIAN, Population: 12_122_244, Version: 1}
	err := dao.Save(s.Ctx, s.DB, &pl)
	t.CmpNoError(err)

	pl2 := planet.Planet{Id: 0, Title: "Pluto", PlanetType: planet.MIXED, Population: 10_122_244, Version: 1}
	err = dao.Save(s.Ctx, s.DB, &pl2)
	t.CmpNoError(err)

	got, err := dao.GetAll(s.Ctx, s.DB)
	expect := []planet.Planet{pl, pl2}

	if t.CmpNoError(err) {
		t.Cmp(
			got,
			expect,
		)
	}
}
