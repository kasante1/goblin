package dir_contents

import (
	"errors"
	"fmt"
	"log"
	"os"
)

// write files into the directories created

// create

// go mod files

// test files

// sample hello world file

// package main file

// folder package file

func Create_project_files(SubDirectories, fileName string) {
	_, err := os.Stat(SubDirectories)

	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("file does not exit")
	} else {
		fmt.Println("file exits")
	}
}

func Write_project_files(fileName string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	if _, err := file.WriteString("some project stuff"); err != nil {
		log.Fatal(err)
	}
}

func Create_mod_file() {}

func Create_test_files() {}

func Create_main_project_file()

// use to illustrate how packages
// work in go lang
func Create_folder_package() {}
