package main

import (
	"log"
	"os"

	bk "github.com/s111ew/bk/cmd"
)

const ALIAS_FILE_PATH = "~/.bk"

func main() {
	err := bk.Run(os.Args[1:], ALIAS_FILE_PATH)

	if err != nil {
		log.Fatal(err)
	}

}
