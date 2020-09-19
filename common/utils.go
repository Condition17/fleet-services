package common

import (
	"fmt"
	"os"
)

func GetFullExternalServiceName(baseServiceName string) string {
	if os.Getenv("ENV_NAME") == string(ProdEnv) {
		return fmt.Sprintf("%s:8080", baseServiceName)
	}

	return fmt.Sprintf("go.micro.api.%s", baseServiceName)
}
