package dir_contents

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"regexp"
)

// file contents
//main file

func MainFileContents() string {
	var mainFileContents string = `package main

import "fmt"

func main(){
  fmt.Println("Hello world!")
}

`
	return mainFileContents
}

// modify go version output
// from go1.19.2 to go 1.19
func modifyVersionOutput() string {

	go_version := runtime.Version()
	
	//space between go and version numbers
	go_space := strings.NewReplacer("go", "go ")
	
	//replace go version with spaced go version
	spaced_content:= go_space.Replace(go_version)
	
	//remove trailing minor semver
	version_output  := regexp.MustCompile(`[.]\d$`).Split(spaced_content, -1)

	// return string of version_output which is a slice
	return version_output[0]
}

// adapter for the write to file function
func GoModFileContent(projectName string) string {

	var go_version = modifyVersionOutput()

	var templateGoModFile = "module name\n\nversion"

	var replace_template = strings.NewReplacer("name", projectName, "version", string(go_version))

	var GoModFileContent = replace_template.Replace(templateGoModFile)

	return GoModFileContent

}


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
		fmt.Println("[X]", fileName, "failed. already exits!")
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

