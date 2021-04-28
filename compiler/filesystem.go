package compiler

import (
	"os"
	"os/user"
	"path/filepath"
)

type GinkerDir struct {
	// User home directory.
	HomeDir string
	// Working directory
	Dir string
	// Working file.
	File string
	// Module name
	ModuleName string
}

func GetDirConfig() (*GinkerDir, error) {
	ginkerDir := GinkerDir{}

	usr, err := user.Current()
	if err != nil {
		return nil, err
	}

	ginkerDir.HomeDir = usr.HomeDir
	ginkerDir.Dir = filepath.Join(ginkerDir.HomeDir, ".ginker")
	ginkerDir.File = filepath.Join(ginkerDir.Dir, "ginker.go")
	ginkerDir.ModuleName = "example.com/ginker"

	return &ginkerDir, nil
}

func GetConfig() string {
	ginkerDir, err := GetDirConfig()
	if err != nil {
	}

	path := filepath.Join(ginkerDir.HomeDir, ginkerDir.Dir, "ginker.env")

	return path
}

func CreateWorkingDir() (bool, error) {
	dirConfig, err := GetDirConfig()
	if err != nil {
		return false, err
	}

	os.Mkdir(dirConfig.Dir, 0755)
	os.Create(dirConfig.File)

	return true, nil
}
