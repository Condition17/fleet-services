package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

const riverImageName string = "cconache/river3:latest"

func runRiverContainer(volumePath string, cmdArgs []string, finishChan chan<- bool, errorChan chan<- error) {
	dockerCmd := fmt.Sprintf("sudo docker run -v %s:/mount %s %s", volumePath, riverImageName, strings.Join(cmdArgs, " "))
	if err := exec.Command("/bin/sh", "-c", dockerCmd).Run(); err != nil {
		errorChan <- err
	}
	finishChan <- true
}

func main() {
	volumePath := "/Users/cristian_conache/Workspace/learning/river/mount"
	runFinishChan := make(chan bool)
	runErrorChan := make(chan error)

	go runRiverContainer(volumePath,
		[]string{
			"-bp", "/mount/sage",
			"-secondsBetweenStats", "2",
			"-arch", "x64",
			"-max", "1",
			"-outputType", "textual",
		}, runFinishChan, runErrorChan);

	select {
	case _ = <-runFinishChan:
		log.Println("Run successfully finished")
	case err := <-runErrorChan:
		log.Printf("River container run encountered error: %v", err)
	}
}
