package cmd

import (
	"fmt"

	"github.com/s111ew/bk/internal/alias"
	"github.com/s111ew/bk/internal/fs"
)

type Files struct {
	AliasFile  string
	ConfigFile string
	ZshFile    string
}

var f = Files{
	AliasFile:  ".bk",
	ConfigFile: ".zshrc",
	ZshFile:    ".bk.zsh",
}

func Run(input []string) error {

	if len(input) == 0 {
		return fmt.Errorf("bk: command required. See 'bk --help'.")
	}

	if len(input) > 3 {
		return fmt.Errorf("bk: too many arguments. See 'bk --help'.")
	}

	aliasFp, configFp, zshFp, err := fs.GeneratePaths(f.AliasFile, f.ConfigFile, f.ZshFile)
	if err != nil {
		return err
	}

	if err := fs.Setup(aliasFp, configFp, zshFp); err != nil {
		return err
	}

	command := input[0]
	args := input[1:]

	switch command {

	case "-g", "--get":
		path, err := alias.Resolve(args, aliasFp)
		if err != nil {
			return err
		}
		fmt.Println(path)

	case "-a", "--add":
		if err := alias.Add(args, aliasFp); err != nil {
			return err
		}

	case "-r", "--remove":
		if err := alias.Remove(args, aliasFp); err != nil {
			return err
		}

	case "-u", "--update":
		if err := alias.Update(args, aliasFp); err != nil {
			return err
		}

	case "-l", "--list":
		if err := alias.List(aliasFp); err != nil {
			return err
		}

	case "-h", "--help":
		return fmt.Errorf(HELP_TEXT)

	case "--resolve":
		fmt.Println(alias.UnsafeResolve(args, aliasFp))

	default:
		return fmt.Errorf("bk: '%s' is not a command. See 'bk --help'.", command)

	}

	return nil
}
