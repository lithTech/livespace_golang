package assets

import (
	"embed"
	log "log/slog"
)

type Assets struct {
	Sql embed.FS
	Yaml embed.FS
}

//go:embed sql
var sql embed.FS

//go:embed yaml
var yaml embed.FS

var assets = Assets{Sql: sql, Yaml: yaml}

func GetAssets() *Assets {
	log.Debug("Returning FS")
	return &assets
}