package main

import (
	"fmt"
)

func repl() {
	for {
		var input string
		fmt.Scanln(&input)

		if cmd, exists := commands[input]; exists {
			err := cmd.callback()
			if err != nil {
				fmt.Printf("Error executing command: %v\n", err)
			}
		} else {
			fmt.Println("Unknown command. Type 'help' for available commands.")
		}
	}
}
