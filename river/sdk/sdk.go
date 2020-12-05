package sdk

import (
	"errors"
	"fmt"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

const dockerImageName = "cconache/river3:latest"
const dockerContainerVolumeDstPath = "/mount/"

func Run(targetFilePath string, args ...string) ([]byte, error) {
	fileDir, fileName := filepath.Split(targetFilePath)

	if fileName == "" {
		return nil, errors.New("invalid target file path")
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

	return exec.Command("/bin/sh", "-c", dockerCmdStr).CombinedOutput()
}
