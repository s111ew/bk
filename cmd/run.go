package cmd

import (
	"fmt"

	"github.com/s111ew/bk/internal/alias"
	"github.com/s111ew/bk/internal/fs"
)

const ALIAS_FILE = ".bk"
const CONFIG_FILE = ".zshrc"

func Run(args []string) error {

	if len(args) == 0 {
		return fmt.Errorf("bk: bk requires a command. See 'bk --help'.")
	}

	if len(args) > 3 {
		return fmt.Errorf("bk: too many arguments. See 'bk --help'.")
	}

	aliasFilePath, configFilePath, err := fs.GeneratePaths(ALIAS_FILE, CONFIG_FILE)
	if err != nil {
		return err
	}

	if err := fs.Setup(aliasFilePath, configFilePath); err != nil {
		return err
	}

	switch args[0] {

	case "-g", "--get":
		resolvedPath, err := alias.ResolveAlias(args[1:], aliasFilePath)
		if err != nil {
			return err
		}
		fmt.Println(resolvedPath)

	case "-a", "--add":
		if err := alias.AddAlias(args[1:], aliasFilePath); err != nil {
			return err
		}

	case "-r", "--remove":
		if err := alias.RemoveAlias(args[1:], aliasFilePath); err != nil {
			return err
		}

	case "-u", "--update":
		if err := alias.UpdateAlias(args[1:], aliasFilePath); err != nil {
			return err
		}

	case "-l", "--list":
		if err := alias.ListAliases(aliasFilePath); err != nil {
			return err
		}

	case "-h", "--help":
		return fmt.Errorf(HELP_TEXT)

	case "--resolve":
		fmt.Println(alias.UnsafeResolveAlias(args[1:], aliasFilePath))

	default:
		return fmt.Errorf("bk: '%s' is not a bk command. See 'bk --help'.", args[0])

	}

	return nil
}
