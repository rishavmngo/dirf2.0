package main

import (
	"dirf/dirfUtils"
	"os"
	"strings"
)

func addSlashOrNot(commandIndex int, subCommands []string) string {

	if commandIndex == len(subCommands)-1 {
		return ""
	} else {
		return "/"
	}

}

func ExpandPath(command string) {

	subCommands := strings.Split(command, "/")
	// lastFileDir := subCommand[len(subCommand) - 1] == ""

	paths := []string{""}
	temp := []string{}

	for commandIndex, subCommand := range subCommands {
		splitByPlus := strings.Split(subCommand, "+")
		if len(splitByPlus) < 2 {

		}
		for _, dir := range splitByPlus {
			for _, path := range paths {

				temp = append(temp, path+dir+addSlashOrNot(commandIndex, subCommands))
			}
		}
		paths = temp
		temp = nil
	}

	dirfUtils.CreateFromPaths(paths)

}

func main() {
	commands := os.Args[1:]

	for _, command := range commands {
		ExpandPath(command)
	}

}
