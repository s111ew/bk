package aliasfile

import "os"

func MakeAliasFileIfNotExists(path string) error {
	_, err := os.Stat(path)

	if err == nil {
		return nil
	}

	if os.IsNotExist(err) {
		return nil
	}

	return err
}
