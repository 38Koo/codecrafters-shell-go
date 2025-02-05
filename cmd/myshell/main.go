package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
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
			default:
			// commandから最後の文字(\n)を削除する
			fmt.Println(command + ": command not found")
				
		}


	}
}
