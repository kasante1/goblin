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

	// checked if the cli_argument supplied is a directory
	fileOrDirectory := IsArgumentDirectory(project_name)

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

		// create project directory and project name
		project_directory := filepath.Join(path, project_name)

		// check if the directory does not exits
		file_status := AlreadyExist(project_directory)

		if file_status != nil {
			fmt.Println(file_status)
		}

	}

}

// is the provided cli_argument a string
// or a directory

func IsArgumentDirectory(cli_argument string) bool {

	// split cli_argument and test if its a direcoty minus
	// the cli_argument is a directory

	if strings.Contains(cli_argument, "/") {

		dir_argument, _ := SplitArgument(cli_argument)
		// check if the dir_argument is a directory
		if _, err := os.Stat(dir_argument); errors.Is(err, os.ErrExist) {
			return true
		}
		//check if the cli_argument is an existing directory
	} else if _, err := os.Stat(cli_argument); errors.Is(err, os.ErrExist) {
		return true
	}
	
	return false
}

// split cli_argument if it contains a slash
func SplitArgument(cli_argument string) (string, string) {

	// split cli_argument by the slash delimiter
	split_argument := strings.Split(cli_argument, "/")

	//get string slice length
	split_argument_length := len(split_argument) - 1

	// get last cli_argument string
	argument_remainder := strings.Join(split_argument[split_argument_length:], "")

	return split_argument[split_argument_length], argument_remainder
}

// is the cli_argument provided an existing directory
// project name
func AlreadyExist(cli_argument string) error {

	// is the cli_argument provided an existing directory
	if _, err := os.Stat(cli_argument); errors.Is(err, os.ErrNotExist) {

		// create directory/folder
		create_project := CreateProjectDirectory(cli_argument)

		// stop if directory creation fails
		if create_project != nil {
			return create_project
		}

		return nil

	} else {

		error_message := cli_argument + " already exits!"

		return errors.New(error_message)
	}
}

// create project directory/folders
func CreateProjectDirectory(cli_argument string) error {
	err := os.Mkdir(cli_argument, os.ModePerm)
	if err != nil {
		return errors.New("project creation failed. check directory permission")
	}

	fmt.Println("project directory created here :", cli_argument)
	return nil
}
