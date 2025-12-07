package files

import (
	"os"
	"path/filepath"
)

func GeneratePaths(aliasFileName, configFileName string) (string, string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", "", err
	}

	aliasFilePath := filepath.Join(home, aliasFileName)
	configFilePath := filepath.Join(home, configFileName)

	return aliasFilePath, configFilePath, nil
}
