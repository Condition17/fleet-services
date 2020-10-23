package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/Condition17/fleet-services/lib"
	"github.com/tkanos/gonfig"
)

type Config struct {
	ServiceName string `json:"SERVICE_NAME"`
	RedisUrl    string `json:"REDIS_URL"`
}

func getEnvironmentName() string {
	return strings.ToLower(getEnvVar("ENV_NAME", string(lib.LocalEnv)))
}

func getDefaultConfigFilePath() string {
	workingDir, _ := os.Getwd()
	return filepath.Join(workingDir, "/env/default.json")
}

func getEnvVar(key, defaultValue string) string {
	envVar := os.Getenv(key)
	if len(envVar) == 0 {
		return defaultValue
	}
	return envVar
}

func GetConfig() Config {
	defaultConfig := Config{}
	err := gonfig.GetConf(getDefaultConfigFilePath(), &defaultConfig)

	if err != nil {
		panic(err)
	}

	return Config{
		ServiceName: getEnvVar("SERVICE_NAME", defaultConfig.ServiceName),
		RedisUrl:    getEnvVar("REDIS_URL", defaultConfig.RedisUrl),
	}
}
