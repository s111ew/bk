package alias

import (
	"errors"
	"os"

	"github.com/s111ew/bk/internal/aliasfile"
)

func AddAlias(args []string) error {
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

	aliases, err := aliasfile.LoadAliases()
	if err != nil {
		return err
	}

	for _, a := range aliases {
		if a.Name == aliasName {
			return errors.New("alias exists")
		}
	}

	newAlias := aliasfile.Alias{
		Name: aliasName,
		Path: path,
	}

	aliases = append(aliases, newAlias)
	aliasfile.WriteAliases(aliases)
}
