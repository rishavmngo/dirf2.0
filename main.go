package main

import (
	"bufio"
	"dirf/dirfUtils"
	"fmt"
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
	var globaVar string

	paths := []string{""}
	temp := []string{}

	// fmt.Println(len(subCommands))

	for commandIndex, subCommand := range subCommands {
		splitByPlus := strings.Split(subCommand, "+")
		for _, dir := range splitByPlus {
			for _, path := range paths {

				if len(dir) > 0 && string(dir[0]) == "$" {

					dir = globaVar + dir[1:]
				} else {
					globaVar = dir
				}

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

	if len(commands) < 1 {
		fmt.Println("Error: path is required")
		return
	}

	for _, command := range commands {
		if string(command[0]) == "-" {
			switch string(command[1:]) {
			case "-noArgs":
				fmt.Print("Paths: ")
				reader := bufio.NewReader(os.Stdin)
				// ReadString will block until the delimiter is entered
				input, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println("An error occured while reading input. Please try again", err)
					return
				}

				// remove the delimeter from the string
				input = strings.TrimSuffix(input, "\n")
				co := strings.Split(input, " ")

				for _, command := range co {

					ExpandPath(command)
				}
			case "h", "-help":
				dirfUtils.PrintHelpMenu()
			default:
				fmt.Println("Error: " + "\"" + string(command[1:]) + "\"" + " " + "not recognized")
				fmt.Println("-h, --help for help menu")
			}
		} else {
			ExpandPath(command)
		}

	}

}
