package bk

import (
	"github.com/s111ew/bk/internal/ctrl"
	"github.com/s111ew/bk/internal/files"
)

func Run(args []string, path string) error {

	if len(args) == 0 || len(args) > 3 {
		// return usage manual
	}

	if err := files.EnsureZshrcConfigured(); err != nil {
		return err
	}

	if err := files.MakeAliasFileIfNotExists(path); err != nil {
		return err
	}

	switch args[0] {

	case "add":
		if err := ctrl.AddAlias(args[1:]); err != nil {
			return err
		}

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
