package main

import (
	"fmt"
	"os"
	// "time"
)

var directoriesNames []string
var pathProject string

func init() {
	// init the list of directories
	directoriesNames = []string{"controllers", "db", "deploy", "models", "router", "interceptors", "script", "tests"}
}

// -------------------------------------------Base functions---------------------------------------------
func NewGoProject(projectName string) bool {
	// create the project directoy
	if !createProjectDirectory(projectName) {
		return false
	}
	// create the project struct
	if !createDirectories("./" + projectName + "/") {
		return false
	}
	pathProject = "./" + projectName + "/"
	// create all base files
	if !createFiles() {
		return false
	}
	// return success if all is good
	fmt.Println("New Golang API project created")
	return true
}

func createProjectDirectory(projectName string) bool {
	if _, err := os.Stat("./" + projectName); os.IsNotExist(err) {
		err = os.Mkdir("./"+projectName, 0777)
		if err != nil {
			fmt.Println(err)
			return false
		}
		return true
	}
	fmt.Printf("directory name (%s) already exists\n", projectName)
	return false
}

func createDirectories(path string) bool {
	for _, directoryNames := range directoriesNames {
		if !createSimpleDir(path+""+directoryNames, 0777) {
			return false
		}
	}
	return true
}
func createFiles() bool {
	if !createServerFile(pathProject + "server.go") {
		return false
	}
	return true
}

// -------------------------------------------Sub functions---------------------------------------------

func createSimpleDir(path string, mode int) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, os.FileMode(mode))
		if err != nil {
			fmt.Println(err)
			return false
		}
		return true
	}
	return false
}
func createServerFile(pathFile string) bool {
	// detect if file exists
	var _, err = os.Stat(pathFile)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(pathFile)
		if err != nil {
			fmt.Println(err.Error())
			return false
		}
		defer file.Close()
	}
	return true
}
