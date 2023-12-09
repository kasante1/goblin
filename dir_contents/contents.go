package dir_contents

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// file contents
//main file

var MainFileContents string = `package main
import "fmt"

func main(){

  fmt.Println("Hello world!")

  result := AddNumbers(3, 4)

  fmt.Println(result)
}


`



var AddNumbersFileContents string = `
package main

func AddNumbers(x, y int) int {
	return x + y
}

`

var TestsFileContents string = `package main

import "testing"
	
func TestAddNumbers(t *testing.T){
	result := AddNumbers(2, 3)

	if result != 5{
		t.Error("incorrect results: expected 5 got ",result)
	}
}


`



// create go project mod file
func CreateModFile(projectName, projectDirectory string) error{

	cmd := exec.Command("go", "mod", "init", projectName)

	cmd.Dir = projectDirectory

	err := cmd.Run()


	if err != nil {
		return err
	}

	return nil

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
		fmt.Println("[ OK ]", "create", fileName, "project success")
		
	} else {
		fmt.Println("[ X ]", "create", fileName, "project failed!")
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

// create file, folder, tests and test directories