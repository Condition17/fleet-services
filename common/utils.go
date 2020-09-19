package common

import (
	"fmt"
	"os"
)

func GetFullExternalServiceName(baseServiceName string) string {
	if os.Getenv("ENV_NAME") == string(ProdEnv) {
		return baseServiceName
	}

	return fmt.Sprintf("go.micro.api.%s", baseServiceName)
}
