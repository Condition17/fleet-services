package nfsModule

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func MountVolume(ipAddr string, fileShareName string, destinationPath string) error {
	// Ensure mount directory is created and ignore any other issue
	_ = os.Mkdir(destinationPath, 0777)

	// Mount volume at the specified destination
	fileSharePath := fmt.Sprintf(":/%s", fileShareName)
	log.Printf("Mounting volume '%s%s' to path '%s'\n", ipAddr, fileSharePath, destinationPath)
	if err := syscall.Mount(fileSharePath, destinationPath, "nfs", 0, fmt.Sprintf("nolock,addr=%s", ipAddr)); err != nil {
		return err
	}
	log.Printf("Volume successfully mount at path '%s'\n", destinationPath)

	// Set mount directory permissions
	if err := setDirFilesPermissions("777", destinationPath); err != nil {
		return err
	}

	return nil
}

func UmountVolume(dirPath string) error {
	log.Printf("Umounting volume at path '%s'\n", dirPath)
	if err := exec.Command("umount", "-l", dirPath).Run(); err != nil {
		return err
	}
	_ = os.Remove(dirPath)
	fmt.Printf("Successfully unmounted volume at path '%v'\n", dirPath)

	return nil
}

func setDirFilesPermissions(mode string, dirPath string) error {
	chmodCmdStr := fmt.Sprintf("chmod -R %s %s", mode, dirPath)
	log.Printf("Executting command: '%s'\n", chmodCmdStr)

	return exec.Command("/bin/sh", "-c", chmodCmdStr).Run()
}