package dir_contents

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// file contents

//main file

var MainFileContents string = `package main

import "fmt"

func main(){
  fmt.Println("Hello world!")
}

`

var GoModFileContent string = `module newproject
go version 1.19.0
`

// write files into the directories created

// create

// go mod files

// test files

// sample hello world file

// package main file

// folder package file

// create project files
func CreateProjectFiles(SubDirectories, fileName, file_contents string) {
	// file directory
	file_path := filepath.Join(SubDirectories, fileName)

	// check if file exist or otherwise
	_, err := os.Stat(file_path)

	// if file does not exits, create file.
	if errors.Is(err, os.ErrNotExist) {
		// write content to files
		WriteProjectFiles(file_path, file_contents)
		fmt.Println("[OK]", fileName, "created succesfully")
	} else {
		fmt.Println("[X]", fileName, "failed")
	}
}

func WriteProjectFiles(fileName, file_contents string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	if _, err := file.WriteString(file_contents); err != nil {
		fmt.Println(err)
	}
}

func Create_mod_file() {}

func Create_test_files() {}

func Create_main_project_file() {}

// use to illustrate how packages
// work in go lang
func Create_folder_package() {}
