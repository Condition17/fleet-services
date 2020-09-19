package common

import (
	"fmt"
	"os"
)

func getFullExternalServiceName(baseServiceName string) string {
	if os.Getenv("ENV_NAME") == string(prodEnv) {
		return baseServiceName
	}

	return fmt.Sprintf("go.micro.api.%s", baseServiceName)
}
