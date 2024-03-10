package test

import (
	"livespace/src/helpers"

	"github.com/maxatome/go-testdeep/td"
)

type Some struct {
	Field1 string `json:"field1"`
	Field2 int    `json:"field2"`
}

func (s *IntegrationTest) TestToJsonOrPanic(t *td.T) {
	some := Some{"blah", 1}

	json := helpers.ToJsonOrPanic(some)

	t.Cmp(json, "{\"field1\":\"blah\",\"field2\":1}")
}
