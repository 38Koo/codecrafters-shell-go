package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

var commands = []string{"echo", "type", "exit"}

func main() {
	dirPath := os.Getenv("PATH")
	paths := strings.Split(dirPath, ":")

	for {

		fmt.Fprint(os.Stdout, "$ ")
		
		// Wait for user input
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		command = command[:len(command)-1]

		if command == "exit 0" {
			os.Exit(0)
		}

		commands := strings.Split(command, " ")

		switch commands[0] {
			case "echo":
				fmt.Println(strings.Join(commands[1:], " "))
			case "type":
				ok := checkBuiltin(commands[1])
				if ok {
					fmt.Println(commands[1] + " is a shell builtin")
					break
				} 
				
				p := checkValid(commands[1], paths)
				if p != "" {
					fmt.Println(commands[1] + " is " + p)
					break
				}

				fmt.Println(commands[1] + ": not found")
				
			default:
				// commandから最後の文字(\n)を削除する
				fmt.Println(command + ": command not found")
				
		}
	}
}

func checkBuiltin(command string) bool {
	for _, c := range commands {
		if c == command {
			return true
		}
	}
	return false
}

func checkValid(command string, paths []string) string {
	for _, p := range paths {
		files, err := os.ReadDir(p)
		if err != nil {
			fmt.Println("Error reading file:", err)
			os.Exit(1)
		}

		for _, f := range files {
			if f.Name() == command {
				return p + "/" + f.Name()
			}
		}
	}

	return ""
}