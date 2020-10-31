package config

import "github.com/Condition17/fleet-services/lib/environment"

type Config struct {
	ServiceName     string `json:"SERVICE_NAME"`
	DbHost          string `json:"DB_HOST"`
	DbUser          string `json:"DB_USER"`
	DbName          string `json:"DB_NAME"`
	DbPassword      string `json:"DB_PASSWORD"`
	GoogleProjectID string `json:"GOOGLE_PROJECT_ID"`
}

func GetConfig() Config {
	var cfg Config
	environment.LoadConfigFromEnv(&cfg)
	return cfg
}
