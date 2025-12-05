package alias

import (
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

	aliasString := aliasfile.ConstructAliasString(aliasName, path)

}
