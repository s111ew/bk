package bk

import (
	"errors"
	"fmt"
)

func Run(args []string) error {

	if len(args) < 2 {
		return errors.New("usage")
	}

	switch args[0] {

	case "add":
		fmt.Println(args[1])

	case "rm":
		fmt.Println(args[1])

	case "update":
		fmt.Println(args[1])

	case "list":
		fmt.Println(args[1])

	default:
		return errors.New("usage")

	}

	return nil
}
