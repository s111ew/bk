package bk

import (
	"errors"
	"fmt"

	"github.com/s111ew/bk/internal/ctrl"
	"github.com/s111ew/bk/internal/files"
)

func Run(args []string, aliasFilePath, configFilePath string) error {

	if len(args) == 0 || len(args) > 3 {
		return errors.New("usage")
	}

	if err := files.MakeAliasFileIfNotExists(aliasFilePath); err != nil {
		return err
	}

	if err := files.EnsureZshrcConfigured(configFilePath); err != nil {
		return err
	}

	switch args[0] {

	case "res":
		resolvedPath, err := ctrl.ResolveAlias(args[1:], aliasFilePath)
		if err != nil {
			return err
		}
		fmt.Println(resolvedPath)

	case "add":
		if err := ctrl.AddAlias(args[1:], aliasFilePath); err != nil {
			return err
		}

	case "rm":
		if err := ctrl.RemoveAlias(args[1:], aliasFilePath); err != nil {
			return err
		}

	case "fix":
		if err := ctrl.UpdateAlias(args[1:], aliasFilePath); err != nil {
			return err
		}

	case "list":
		if err := ctrl.ListAliases(aliasFilePath); err != nil {
			return err
		}

	default:
		// return usage manual

	}

	return nil
}
