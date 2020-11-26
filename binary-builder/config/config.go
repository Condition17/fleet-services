package config

import "github.com/Condition17/fleet-services/lib/environment"

type Config struct {
	GoogleProjectID           string `json:"GOOGLE_PROJECT_ID"`
	RedisUrl                  string `json:"REDIS_URL"`
	FleetServicesGrpcProxyUrl string `json:"FLEET_SERVICES_GRPC_PROXY_URL"`
}

func GetConfig() Config {
	var cfg Config
	environment.LoadConfigFromEnv(&cfg)
	return cfg
}
