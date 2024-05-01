package library

import (
	"fmt"
	"os"
)

func Install() bool {
	installRoot := getInstallationDir()
	if _, err := os.Stat(installRoot); os.IsNotExist(err) {
		err := os.Mkdir(installRoot, 0755)
		if err != nil {
			fmt.Println("Already Installed or no Permissions ")
		}
	}
	return false
}

func getHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err == nil {
		return homeDir
	}
	return "/"
}

func getInstallationDir() string {
	return getHomeDir() + "./ljd"
}
