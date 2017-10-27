package main

import (
	"fmt"
	"os"
)

func main() {

	switch args := os.Args[1:]; args {
	default:
		// No args or incorrect args
		fmt.Println(`Usage: 
    envelope [command]

Examples:
    envelope add    add a new item to a budget`)
	}
}
