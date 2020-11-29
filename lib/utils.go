package lib

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	environment "github.com/Condition17/fleet-services/lib/environment"
)

func GetFullExternalServiceName(baseServiceName string) string {
	if os.Getenv("ENV_NAME") == string(environment.ProdEnv) {
		expr := regexp.MustCompile("([a-z])([A-Z])")
		return fmt.Sprintf("%s:8080", strings.ToLower(expr.ReplaceAllString(baseServiceName, "$1-$2")))
	}

	return fmt.Sprintf("go.micro.api.%s", baseServiceName)
}
