package fs

import (
	"os"
	"strings"
)

func Setup(aliasFilePath, configFilePath, zshFuncsFilePath string) error {
	if err := MakeAliasFileIfNotExists(aliasFilePath); err != nil {
		return err
	}

	if err := makeBkZshFileIfNotExists(zshFuncsFilePath); err != nil {
		return err
	}

	if err := EnsureZshrcConfigured(configFilePath); err != nil {
		return err
	}

	return nil
}

func MakeAliasFileIfNotExists(aliasFilePath string) error {
	_, err := os.Stat(aliasFilePath)

	if err == nil {
		return nil
	}

	if os.IsNotExist(err) {
		f, err := os.Create(aliasFilePath)
		if err != nil {
			return err
		}
		return f.Close()
	}

	return err
}

func makeBkZshFileIfNotExists(zshFuncsFilePath string) error {
	_, err := os.Stat(zshFuncsFilePath)

	if err == nil {
		return nil
	}

	if os.IsNotExist(err) {
		f, err := os.OpenFile(zshFuncsFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = f.WriteString(bkZshContents)
		return err
	}

	return err
}

func EnsureZshrcConfigured(configFilePath string) error {
	data, _ := os.ReadFile(configFilePath)
	contents := string(data)

	if strings.Contains(contents, "# >>> bk init >>>") {
		return nil
	}

	f, err := os.OpenFile(configFilePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(zshInsert)
	return err
}
