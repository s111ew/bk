package alias

import (
	"os"
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

	// create alias string and append it to ~/.bk
}
