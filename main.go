package main

import (
	"fmt"
	"os"

	"github.com/s111ew/bk/cmd"
)

func main() {
	input := os.Args[1:]

	if err := cmd.Run(input); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
