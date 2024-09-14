package main

import (
	"fmt"
	"os"
)

func commandPing() error {
	fmt.Println("Help: Displays a help message")
	return nil
}

func commandExit() error {
	fmt.Println("Goodbye!")
	os.Exit(0)
	return nil
}
