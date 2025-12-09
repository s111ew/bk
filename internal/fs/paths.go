package fs

import (
	"os"
	"path/filepath"
)

func GeneratePaths(aliasFileName, configFileName, zshFuncsFileName string) (string, string, string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", "", "", err
	}

	aliasFilePath := filepath.Join(home, aliasFileName)
	configFilePath := filepath.Join(home, configFileName)
	zshFuncsFilePath := filepath.Join(home, zshFuncsFileName)

	return aliasFilePath, configFilePath, zshFuncsFilePath, nil
}
