package main

import (
	"log"
	"os"
	"path/filepath"

	bk "github.com/s111ew/bk/cmd"
)

const ALIAS_FILE = ".bk"
const CONFIG_FILE = ".zshrc"

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	aliasFilePath := filepath.Join(home, ALIAS_FILE)
	configFilePath := filepath.Join(home, CONFIG_FILE)

	if err := bk.Run(os.Args[1:], aliasFilePath, configFilePath); err != nil {
		log.Fatal(err)
	}
}
