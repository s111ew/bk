package alias

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/s111ew/bk/internal/fs"
)

func Resolve(args []string, aliasFilePath string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("bk: 'bk --get' requires 1 argument '<alias>'. See 'bk --help'.")
	}

	aliasName := args[0]

	aliases, err := fs.LoadAll(aliasFilePath)
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

func UnsafeResolve(args []string, aliasFilePath string) string {
	aliasName := args[0]

	alias, _ := fs.LoadOne(aliasName, aliasFilePath)

	return alias.Path
}

func Add(args []string, aliasFilePath string) error {

	if len(args) < 1 {
		return fmt.Errorf("bk: 'bk --add' requires atleast 1 argument '<alias>' and optionally '<path>'. See 'bk --help'.")
	}

	if len(args) > 2 {
		return fmt.Errorf("bk: 'bk --add' accepts 1 argument '<alias>' and optionally '<path>'. See 'bk --help'.")
	}

	aliasName := args[0]
	if err := sanitize(aliasName); err != nil {
		return err
	}

	var aliasPath string

	if len(args) == 1 {

		currWd, err := os.Getwd()

		if err != nil {
			return err
		}

		aliasPath = currWd

	} else {

		aliasPath = args[1]
		if err := sanitize(aliasPath); err != nil {
			return err
		}

	}

	if err := fs.WriteOne(aliasName, aliasPath, aliasFilePath); err != nil {
		return err
	}

	return nil
}

func Update(args []string, aliasFilePath string) error {
	if len(args) < 1 {
		return fmt.Errorf("bk: 'bk --update' requires atleast 1 argument '<alias>' and optionally '<path>'. See 'bk --help'.")
	}

	if len(args) > 2 {
		return fmt.Errorf("bk: 'bk --update' accepts 1 argument '<alias>' and optionally '<path>'. See 'bk --help'.")
	}

	aliasName := args[0]
	if err := sanitize(aliasName); err != nil {
		return err
	}

	var path string

	if len(args) == 1 {

		currWd, err := os.Getwd()

		if err != nil {
			return err
		}

		path = currWd

	} else {

		path = args[1]
		if err := sanitize(path); err != nil {
			return err
		}

	}

	aliases, err := fs.LoadAll(aliasFilePath)
	if err != nil {
		return err
	}

	for _, a := range aliases {
		if a.Name == aliasName {
			a.Path = path
		}
		aliases = append(aliases, a)
	}

	if err := fs.WriteAll(aliases, aliasFilePath); err != nil {
		return err
	}

	return nil
}

func Remove(args []string, aliasFilePath string) error {
	if len(args) != 1 {
		return fmt.Errorf("bk: 'bk --remove' requires 1 argument '<alias>'. See 'bk --help'.")
	}

	aliasName := args[0]

	err := fs.RemoveOne(aliasName, aliasFilePath)
	if err != nil {
		return err
	}

	return nil
}

func List(aliasFilePath string) error {
	aliases, err := fs.LoadAll(aliasFilePath)
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

func sanitize(s string) error {
	if strings.Contains(s, "=") {
		return fmt.Errorf("bk: input cannot contain '='")
	}

	return nil
}
