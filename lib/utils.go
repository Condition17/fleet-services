package lib

import (
	"fmt"
	"os"

	environment "github.com/Condition17/fleet-services/lib/environment"
)

func GetFullExternalServiceName(baseServiceName string) string {
	if os.Getenv("ENV_NAME") == string(environment.ProdEnv) {
		return fmt.Sprintf("%s:8080", baseServiceName)
	}

	return fmt.Sprintf("go.micro.api.%s", baseServiceName)
}
