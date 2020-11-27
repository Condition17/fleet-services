package config

import "github.com/Condition17/fleet-services/lib/environment"

type Config struct {
	ServiceName           string `json:"SERVICE_NAME"`
	GoogleProjectID       string `json:"GOOGLE_PROJECT_ID"`
	FileBuilderServiceUrl string `json:"FILE_BUILDER_SERVICE_URL"`
}

func GetConfig() Config {
	var cfg Config
	environment.LoadConfigFromEnv(&cfg)
	return cfg
}
