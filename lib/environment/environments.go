package environment

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/ghodss/yaml"
)

type EnvironmentName string

const (
	LocalEnv EnvironmentName = "local"
	ProdEnv  EnvironmentName = "prod"
)

var errInvalidSpecification = errors.New("The env config specification must be a pointer to a struct.")

func getConfigFilePath(envName string) string {
	workDir, _ := os.Getwd()
	return filepath.Join(workDir, "env", fmt.Sprintf("%s.yml", envName))
}

func getEnvVar(key, defaultValue string) string {
	envVar := os.Getenv(key)
	if len(envVar) == 0 {
		return defaultValue
	}
	return envVar
}

func getConfigFromYAML(filename string, config interface{}) {
	if len(filename) == 0 {
		return
	}

	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	return
}

func LoadConfigFromEnv(cfgSpecification interface{}) {
	var specVal reflect.Value = reflect.ValueOf(cfgSpecification)
	var envName string = strings.ToLower(getEnvVar("ENV_NAME", string(LocalEnv)))

	if specVal.Kind() != reflect.Ptr || specVal.Elem().Kind() != reflect.Struct {
		panic(errInvalidSpecification)
	}

	getConfigFromYAML(getConfigFilePath("default"), &cfgSpecification)
	getConfigFromYAML(getConfigFilePath(envName), &cfgSpecification)
}
