package dirfUtils

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func CreateFromPaths(paths []string) {

	workingDirectory, error := os.Getwd()

	if error != nil {
		log.Fatal("error while geting current working directory")
	}

	for _, path := range paths {

		ifTheLastCommandIsDir := string(path[len(path)-1]) == "/"
		if ifTheLastCommandIsDir {

			errorWhileCreatingDirectories := os.MkdirAll(path, 0777)

			if errorWhileCreatingDirectories != nil {
				fmt.Println(errorWhileCreatingDirectories)
			} else {

			}
		} else {
			splittedPath := strings.Split(path, "/")
			for index, subC := range splittedPath {
				if len(splittedPath) == 1 || index == len(splittedPath)-1 {
					os.Create(subC)
				} else {
					os.Mkdir(subC, 0777)
					os.Chdir(subC)
				}
			}
			fmt.Println("Created: ", path)
			os.Chdir(workingDirectory)
		}
	}
}

func PrintHelpMenu() {

	fmt.Println("USAGE:")
	fmt.Println("\tExamples:")
	fmt.Println("\t>> dirf --noArgs\n\t>> Paths: a/b+d/c will expand into ['a/b/c' 'a/d/c']")
	fmt.Println("\t>> dirf --noArgs\n\t>> Paths: a/$.component.jsx+$.styles.css (will expand into ['a/a.component.jsx' 'a/a.styles.css'])")
	fmt.Println("\tExamples:")
	fmt.Println("\t>> dirf a/b/c d/c+c/e (can pass the paths directly as arguments but '$' expansion) can be used only using escape sequence and single quotes")
	fmt.Println("OPTIONS:")
	fmt.Println("\t--noArgs \tCreate directory and fils with standard input")
	fmt.Println("\t--help,-h \tFor help")
}
