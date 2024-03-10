package config

import (
	embed "embed"
	"livespace/assets"
	log "log/slog"

	yaml "gopkg.in/yaml.v3"
)

type DbConnectionConfig struct {
	Url string
}

type DbConfig struct {
	Connection DbConnectionConfig
}

type ApplicationConfig struct {
	Database DbConfig
}

var applicationConfig = ApplicationConfig{}

func LoadConfig(assets *assets.Assets) *ApplicationConfig {
	applicationYaml, err := embed.FS.ReadFile(assets.Yaml, "yaml/application.yaml")
	if err != nil {
		log.Error("Failed to find application config", err)
		panic(err)
	}
	if err := yaml.Unmarshal(applicationYaml, &applicationConfig); err != nil {
		log.Error("Failed to parse application config:", err)
		panic(err)
	}
	
	return &applicationConfig
}
