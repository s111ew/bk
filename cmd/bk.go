package bk

import (
	"errors"

	"github.com/s111ew/bk/internal/ctrl"
	"github.com/s111ew/bk/internal/files"
)

func Run(args []string, alias_file_path, config_file_path string) error {

	if len(args) == 0 || len(args) > 3 {
		return errors.New("usage")
	}

	if err := files.MakeAliasFileIfNotExists(alias_file_path); err != nil {
		return err
	}

	if err := files.EnsureZshrcConfigured(config_file_path); err != nil {
		return err
	}

	switch args[0] {

	case "add":
		if err := ctrl.AddAlias(args[1:], alias_file_path); err != nil {
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
