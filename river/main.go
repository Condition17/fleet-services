package main

import (
	riverSdk "github.com/Condition17/fleet-services/river/sdk"
	"log"
)

func main() {
	exitCode, err := riverSdk.Run(
		"/Users/cristian_conache/Workspace/learning/river/River3/TestPrograms/crash_detection",
		"-secondsBetweenStats", "2",
		"-arch", "x64",
		"-max", "1",
		"-outputType", "textual",
	)

	if exitCode == 0 {
		log.Println("Command succeeded")
	} else {
		log.Println("Command finished with exit code:", exitCode)
		log.Printf("River command error: %v", err)
	}
}
