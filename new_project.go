package main

import (
	"fmt"
	"os"
	// "time"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	template "text/template"
)

type Imports struct {
	UrlDbPackage          string
	UrlScriptPackage      string
	UrlRouterPackage      string
	UrlConfigPackage      string
	UrlControllerPackage  string
	UrlModelPackage       string
	UrlInterceptorPackage string
}

var directoriesNames []string
var pathProject string
var pathDir string
var currentPath string
var Import Imports

func init() {
	// set default GOLANG_ENV = dev
	value := os.Getenv("GOLANG_ENV")
	if len(value) == 0 {
		os.Setenv("GOLANG_ENV", "dev")
	}

	// get katze full path package
	_, katzePackage, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	pathDir = path.Dir(katzePackage)

	// get current project path
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	if strings.Contains(dir, os.Getenv("GOPATH")+"/src/") {
		currentPath = strings.Replace(dir, os.Getenv("GOPATH")+"/src/", "", -1)
	} else {
		panic("The current structure is not supported: " + dir + " must be under gopath " + os.Getenv("GOPATH") + "/src")
	}

	// init the list of directories
	directoriesNames = []string{"controllers", "config", "db", "deploy", "models", "router", "interceptors", "script", "tests"}
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
	currentPath = currentPath + "/" + projectName
	// init imports
	Import = Imports{
		UrlDbPackage:          currentPath + "/db",
		UrlScriptPackage:      currentPath + "/script",
		UrlRouterPackage:      currentPath + "/router",
		UrlConfigPackage:      currentPath + "/config",
		UrlControllerPackage:  currentPath + "/controllers",
		UrlModelPackage:       currentPath + "/models",
		UrlInterceptorPackage: currentPath + "/interceptors",
	}
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
	if !createFile(pathProject+"server.go", "server.tmpl", "templates/server.tmpl") {
		return false
	}
	if !createFile(pathProject+"config/config.go", "config.tmpl", "templates/config.tmpl") {
		return false
	}
	if !createFile(pathProject+"db/db.go", "db.tmpl", "templates/db.tmpl") {
		return false
	}
	if !createFile(pathProject+"script/scriptDatabase.go", "scriptDatabase.tmpl", "templates/scriptDatabase.tmpl") {
		return false
	}
	if !createFile(pathProject+"script/seedDatabase.go", "seedDatabase.tmpl", "templates/seedDatabase.tmpl") {
		return false
	}
	if !createFile(pathProject+"router/routes.go", "routes.tmpl", "templates/routes.tmpl") {
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
func createFile(pathFile string, templateName string, templatePath string) bool {
	// detect if file exists
	file, err := os.OpenFile(pathFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer file.Close()

	t := template.Must(template.New(templateName).ParseFiles(pathDir + "/" + templatePath))
	err2 := t.Execute(file, Import)
	if err2 != nil {
		fmt.Println(err2)
		return false
	}

	defer file.Close()
	return true
}
