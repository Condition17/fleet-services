package config

import (
	"os"
	"path/filepath"
	"strings"

	common "github.com/Condition17/fleet-services/common"
	"github.com/tkanos/gonfig"
)

type Config struct {
	ServiceName string `json:"SERVICE_NAME"`
}

func getEnvironmentName() string {
	return strings.ToLower(getEnvVar("ENV_NAME", string(common.LocalEnv)))
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
	}
}
