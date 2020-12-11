package sdk

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"syscall"
)

const dockerImageName = "cconache/river3:latest"
const dockerContainerVolumeDstPath = "/mount/"

func Run(targetFilePath string, args ...string) (int, error) {
	var err error
	var cmdOutput []byte
	fileDir, fileName := filepath.Split(targetFilePath)

	if fileName == "" {
		return 1, errors.New("invalid target file path")
	}

	cmdArgs := append(
		[]string{"-bp", path.Join("/mount/", fileName)},
		args...,
	)
	dockerCmdStr := fmt.Sprintf(
		"docker run -v %s:%s %s %s",
		fileDir,
		dockerContainerVolumeDstPath,
		dockerImageName,
		strings.Join(cmdArgs, " "),
	)

	if cmdOutput, err = exec.Command("/bin/sh", "-c", dockerCmdStr).CombinedOutput(); err != nil {
		log.Printf("%s\n", cmdOutput)
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				return status.ExitStatus(), err
			}
		} else {
			log.Printf("Error encountered while trying to execute command '%v': %v\n", dockerCmdStr, err)
			return 1, err
		}
	}

	log.Printf("%s\n", cmdOutput)
	return 0, nil
}
