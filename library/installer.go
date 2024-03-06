package library

import "os"

func install(installationDir string) {
	os.Mkdir(installationDir+"./ljd", 644)
}
