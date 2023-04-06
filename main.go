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
		if _, err := os.Stat(project_name); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(project_name, os.ModePerm)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("project directory created here :", project_name)

		} else {
			fmt.Println(project_name, " already exits! dir")
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

		// if err := os.Mkdir("a", os.ModePerm); err != nil {
		//     log.Fatal(err)
		// }

		// check if the directory does not exits
		if _, err := os.Stat(project_directory); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(project_directory, os.ModePerm)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("project directory created here : ", project_directory)

		} else {

			fmt.Println(project_directory, " already exits file")
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

func AlreadyExist(argument string) (bool, error) {
	return true, nil
}