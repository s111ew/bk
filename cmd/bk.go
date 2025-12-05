package bk

import "github.com/s111ew/bk/internal/aliasfile"

func Run(args []string, path string) error {

	if len(args) == 0 {
		// return usage manual
	}

	aliasfile.MakeAliasFileIfNotExists(path)

	switch args[0] {

	case "add":
		// add a new alias/path pair

	case "rm":
		// remove an alias/path pair given an alias

	case "fix":
		// update an alias/path pair given an an alias and a new path

	case "list":
		// return a table of alias/path pairs

	default:
		// return usage manual

	}

	return nil
}
