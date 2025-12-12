package alias

import (
	"errors"
	"fmt"
	"os"

	"github.com/s111ew/bk/internal/fs"
)

func ResolveAlias(args []string, aliasFilePath string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("bk: 'bk --get' requires 1 argument '<alias>'. See 'bk --help'.")
	}

	aliasName := args[0]

	aliases, err := fs.LoadAliases(aliasFilePath)
	if err != nil {
		return "", err
	}

	for _, a := range aliases {
		if a.Name == aliasName {
			return a.Path, nil
		}
	}

	return "", errors.New("alias not found")
}

func UnsafeResolveAlias(args []string, aliasFilePath string) string {
	aliasName := args[0]

	alias, _ := fs.LoadSingleAlias(aliasName, aliasFilePath)

	return alias.Path
}

func AddAlias(args []string, aliasFilePath string) error {

	if len(args) < 1 {
		return fmt.Errorf("bk: 'bk --add' requires atleast 1 argument '<alias>' and optionally '<path>'. See 'bk --help'.")
	}

	if len(args) > 2 {
		return fmt.Errorf("bk: 'bk --add' accepts 1 argument '<alias>' and optionally '<path>'. See 'bk --help'.")
	}

	aliasName := args[0]

	var aliasPath string

	if len(args) == 1 {

		currWd, err := os.Getwd()

		if err != nil {
			return err
		}

		aliasPath = currWd

	} else {

		aliasPath = args[1]

	}

	if err := fs.WriteSingleAlias(aliasName, aliasPath, aliasFilePath); err != nil {
		return err
	}

	return nil
}

func UpdateAlias(args []string, aliasFilePath string) error {
	if len(args) < 1 {
		return fmt.Errorf("bk: 'bk --update' requires atleast 1 argument '<alias>' and optionally '<path>'. See 'bk --help'.")
	}

	if len(args) > 2 {
		return fmt.Errorf("bk: 'bk --update' accepts 1 argument '<alias>' and optionally '<path>'. See 'bk --help'.")
	}

	aliasName := args[0]

	var path string

	if len(args) == 1 {

		currWd, err := os.Getwd()

		if err != nil {
			return err
		}

		path = currWd

	} else {

		path = args[1]

	}

	aliases, err := fs.LoadAliases(aliasFilePath)
	if err != nil {
		return err
	}

	for _, a := range aliases {
		if a.Name == aliasName {
			a.Path = path
		}
		aliases = append(aliases, a)
	}

	if err := fs.WriteAliases(aliases, aliasFilePath); err != nil {
		return err
	}

	return nil
}

func RemoveAlias(args []string, aliasFilePath string) error {
	if len(args) != 1 {
		return fmt.Errorf("bk: 'bk --remove' requires 1 argument '<alias>'. See 'bk --help'.")
	}

	aliasName := args[0]

	err := fs.RemoveSingleAlias(aliasName, aliasFilePath)
	if err != nil {
		return err
	}

	return nil
}

func ListAliases(aliasFilePath string) error {
	aliases, err := fs.LoadAliases(aliasFilePath)
	if err != nil {
		return err
	}

	fmt.Printf("%-10s %-40s\n", "Alias", "Path")
	fmt.Println("-----------------------------------------------")

	for _, a := range aliases {
		fmt.Printf("%-10s %-40s\n", a.Name, a.Path)
	}

	return nil
}
