package planet

type Type uint8

const ( // PlanetType
	NON_HABITAT = 1 + iota
	AGRARIAN
	SCIENTIFIC
	MIXED
	PRIMITIVE
)

type Planet struct {
	Id         int64  `db:"id"`
	Title      string `db:"title"`
	PlanetType Type   `db:"planet_type"`
	Population uint64 `db:"population"`
	Version    uint16 `db:"version"`
}
