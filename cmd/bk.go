package bk

import (
	"fmt"

	"github.com/s111ew/bk/internal/ctrl"
	"github.com/s111ew/bk/internal/files"
)

func Run(args []string, aliasFilePath, configFilePath string) error {

	if len(args) == 0 || len(args) > 3 {
		return fmt.Errorf(USAGE_TEXT)
	}

	if err := files.MakeAliasFileIfNotExists(aliasFilePath); err != nil {
		return err
	}

	if err := files.EnsureZshrcConfigured(configFilePath); err != nil {
		return err
	}

	switch args[0] {

	case "-g", "--get":
		resolvedPath, err := ctrl.ResolveAlias(args[1:], aliasFilePath)
		if err != nil {
			return err
		}
		fmt.Println(resolvedPath)

	case "-a", "--add":
		if err := ctrl.AddAlias(args[1:], aliasFilePath); err != nil {
			return err
		}

	case "-r", "--remove":
		if err := ctrl.RemoveAlias(args[1:], aliasFilePath); err != nil {
			return err
		}

	case "-u", "--update":
		if err := ctrl.UpdateAlias(args[1:], aliasFilePath); err != nil {
			return err
		}

	case "-l", "--list":
		if err := ctrl.ListAliases(aliasFilePath); err != nil {
			return err
		}

	case "-h", "--help":
		return fmt.Errorf(HELP_TEXT)

	default:
		return fmt.Errorf(USAGE_TEXT)

	}

	return nil
}
