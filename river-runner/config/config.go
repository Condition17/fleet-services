package config

import "github.com/Condition17/fleet-services/lib/environment"

type Config struct {
	GoogleProjectID           string `json:"GOOGLE_PROJECT_ID"`
	FleetServicesGrpcProxyUrl string `json:"FLEET_SERVICES_GRPC_PROXY_URL"`
	FleetServicesHttpApiUrl   string `json:"FLEET_SERVICES_HTTP_API_URL"`
	ServerPort                string `json:"SERVER_PORT"`
}

func GetConfig() Config {
	var cfg Config
	environment.LoadConfigFromEnv(&cfg)
	return cfg
}
