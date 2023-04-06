package main

import (
	"fmt"
	"os"
	//"log"
	"errors"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("Goblin!\n Usage: -bin|-lib <project_name>")

	if len(os.Args) == 1 {
		fmt.Println("Enter a project name.")
		return
	}

	// get project name
	arguments := os.Args

	project_name := arguments[1]

	// checked if the argument supplied is a directory
	fileOrDirectory, err := IsArgumentDirectory(project_name)

	if err != nil {
		fmt.Println(err)
	}

	if fileOrDirectory {
		dir_status := AlreadyExist(project_name)

		if dir_status != nil {
			fmt.Println(dir_status)
		}

	} else {
		// get current directory or where to save
		// project

		path, err := os.Getwd()

		if err != nil {
			fmt.Println(err)
		}

		// fmt.Println(path)

		// create project directory and project name
		project_directory := filepath.Join(path, project_name)

		// check if the directory does not exits
		file_status := AlreadyExist(project_directory)

		if file_status != nil {
			fmt.Println(file_status)
		}

	}

}

// is the provided argument a string
// or a directory

func IsArgumentDirectory(argument string) (bool, error) {

	if strings.Contains(argument, "\"") || strings.Contains(argument, "/") {
		return true, nil
	}
	return false, errors.New("a string not a directory")

}

// is the argument provided an existing directory
// project name

func AlreadyExist(cli_argument string) error {
	if _, err := os.Stat(cli_argument); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(cli_argument, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("project directory created here :", cli_argument)

		return nil

	} else {
		//fmt.Println(cli_argument, " already exits!")

		error_message := cli_argument + " already exits!"

		return errors.New(error_message)
	}
}