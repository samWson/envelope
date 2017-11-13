// MIT License
//
// Copyright (c) 2017 Samuel Wilson <samWson@users.noreply.github.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"fmt"
	"github.com/samWson/envelope/cmd"
	"github.com/samWson/envelope/record"
	"log"
	"os"
	"os/user"
)

const headers = "category,amount,date\n"

// Main passes arguments on to subcommands.
func main() {

	verifyBudget()

	var args []string = os.Args[1:]

	if len(args) < 1 {
		fmt.Println(help)
		return
	}

	switch first := args[0]; first {
	case "add":
		var trans record.Transaction
		cmd.Add(args[1:], &trans)
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

// verifyBudget checks that the budget file exists, or creates a default file.
func verifyBudget() {

	usr, err := user.Current()
	if err != nil {
		log.Fatalf("Error getting the current user: %v", err)
	}

	home := usr.HomeDir
	budget := home + "/.budget"

	if _, err := os.Stat(budget); os.IsNotExist(err) {

		fmt.Println(budget, "file does not exist. Creating default file at:", budget)

		file, err := os.Create(budget)
		defer file.Close()

		if err != nil {
			log.Fatalf("Error creating budget file: %v", err)
		}

		writeHeaders(file, headers)
	}
}

// writeHeaders writes column headers to an empty budget file.
func writeHeaders(f *os.File, h string) {

	_, err := f.WriteString(h)

	if err != nil {
		log.Fatalf("Error writing headers to budget file: %v", err)
	}
}
