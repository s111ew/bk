package main

import (
	"log"
	"os"

	bk "github.com/s111ew/bk/cmd"
)

func main() {
	err := bk.Run(os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}

}
