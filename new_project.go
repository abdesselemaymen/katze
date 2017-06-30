package main

import (
	"fmt"
	"os"
	// "time"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
)

type Server struct {
	UrlDbPackage     string
	UrlScriptPackage string
	UrlRouterPackage string
}

var directoriesNames []string
var pathProject string
var pathDir string
var currentPath string

func init() {
	// get katze package full path
	_, katzePackage, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	pathDir = path.Dir(katzePackage)

	// get the current project path
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	if strings.Contains(dir, os.Getenv("GOPATH")+"/src/") {
		currentPath = strings.Replace(pathDir, os.Getenv("GOPATH")+"/src/", "", -1)
	} else {
		panic("The current structure is not supported: " + dir + " must be under gopath " + os.Getenv("GOPATH") + "/src")
	}
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
	file, err := os.OpenFile(pathFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer file.Close()

	server := Server{
		UrlDbPackage:     currentPath + "/db",
		UrlScriptPackage: currentPath + "/script",
		UrlRouterPackage: currentPath + "/routes",
	}
	t := template.Must(template.New("server.tmpl").ParseFiles(pathDir + "/templates/server.tmpl"))
	err2 := t.Execute(file, server)
	if err2 != nil {
		fmt.Println(err2)
		return false
	}

	defer file.Close()
	return true
}
