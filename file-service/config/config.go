package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/tkanos/gonfig"
)

type Config struct {
	ServiceName string `json:"SERVICE_NAME"`
	RedisUrl    string `json:"REDIS_URL"`
}

func getEnvironmentName() string {
	env := os.Getenv("ENV_NAME")
	if "" == env {
		env = "local"
	}

	return strings.ToLower(env)
}

func getEnvFilePath() string {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}

	workingDir, _ := os.Getwd()
	return filepath.Join(workingDir, fmt.Sprintf("/env/%s_config.json", getEnvironmentName()))
}

func GetConfig() Config {
	config := Config{}
	fileName := getEnvFilePath()
	err := gonfig.GetConf(fileName, &config)

	if err != nil {
		panic(err)
	}

	return config
}
