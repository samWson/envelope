package main

import (
	"fmt"
	"os"
)

func main() {

	var args []string = os.Args[1:]

	if len(args) < 1 {
		fmt.Println(help)
		return
	}

	switch first := args[0]; first {
	case "add":
		fmt.Println("Add called")
	default:
		// No args or incorrect args
		fmt.Println(help)
		return
	}

}

var help string = `Envelope is a budgeting tool implementing the envelope method.

Usage:
    envelope [command]

Examples:
    add    add a new item to a budget`
