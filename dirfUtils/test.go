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
