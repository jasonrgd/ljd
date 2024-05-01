package library

import (
	"os"
	"testing"
)

func TestInstall(t *testing.T) {
	// Temporary directory for testing
	testDir := "/tmp/ljd_test"

	// Mocking os.Stat to return not exist
	oldStat := osStat
	osStat = func(name string) (os.FileInfo, error) {
		return nil, os.ErrNotExist
	}
	defer func() { osStat = oldStat }()

	// Mocking os.Mkdir to return nil error
	oldMkdir := osMkdir
	osMkdir = func(name string, perm os.FileMode) error {
		return nil
	}
	defer func() { osMkdir = oldMkdir }()

	// Call the function under test
	result := Install()

	// Check if the directory is created
	if _, err := os.Stat(testDir); os.IsNotExist(err) {
		t.Errorf("Install() did not create the installation directory: %v", err)
	}

	// Check if the result is false
	if result != false {
		t.Errorf("Install() did not return false as expected")
	}
}

func TestInstallAlreadyInstalled(t *testing.T) {
	// Mocking os.Stat to return exist
	oldStat := osStat
	osStat = func(name string) (os.FileInfo, error) {
		return nil, nil
	}
	defer func() { osStat = oldStat }()

	// Call the function under test
	result := Install()

	// Check if the result is false
	if result != false {
		t.Errorf("Install() did not return false as expected")
	}
}

func TestInstallPermissionError(t *testing.T) {
	// Temporary directory for testing
	testDir := "/tmp/ljd_test"

	// Mocking os.Stat to return not exist
	oldStat := osStat
	osStat = func(name string) (os.FileInfo, error) {
		return nil, os.ErrPermission
	}
	defer func() { osStat = oldStat }()

	// Call the function under test
	result := Install()

	// Check if the directory is created
	if _, err := os.Stat(testDir); os.IsNotExist(err) {
		t.Errorf("Install() did not attempt to create the installation directory with permission error: %v", err)
	}

	// Check if the result is false
	if result != false {
		t.Errorf("Install() did not return false as expected")
	}
}
