package main

import (
	"fmt"
	"goblin/cli_args"
	"os"
)

func main() {
	fmt.Println(`goblin! -- create a go project

Usage:
  goblin <project_name>

  cd project_name
  go run project_name or go run .
  go test -v
  go build project_name`)


	if len(os.Args) == 1 {
		fmt.Println("\n[X] enter a project name.\n")
		return
	}

	// get fetch cli_arguments
	cli_arguments := os.Args

	// get cli project argument
	cli_project_argument := cli_arguments[1]

	// checked if the cli_argument supplied is a directory
	var fileOrDirectory bool = cli_args.IsArgumentDirectory(cli_project_argument)

	// if cli is a directory notify user or
	// create project in the existing directory
	if fileOrDirectory {
		var dir_status error = cli_args.AlreadyExist(cli_project_argument)

		if dir_status != nil {
			fmt.Println(dir_status)
		}

	} else {
		// if cli_project_argument is not a directory
		// create new directory in CWD
		cli_args.NewProjectDirectry(cli_project_argument)
	}

}
