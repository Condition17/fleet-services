package main

import (
	riverSdk "github.com/Condition17/fleet-services/river/sdk"
	"log"
)

func main() {
	output, err := riverSdk.Run(
		"/Users/cristian_conache/Workspace/learning/river/River3/TestPrograms/crackme_xor",
		"-secondsBetweenStats", "2",
		"-arch", "x64",
		"-max", "1",
		"-outputType", "textual",
	)

	if err != nil {
		log.Printf("River command error: %v", err)
	}
	log.Printf("River command output: %s", output)
}
