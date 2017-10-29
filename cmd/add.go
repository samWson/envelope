// Package cmd implements the commands of Envelope.
package cmd

import "fmt"

// Add adds a new item to a budget depending on the subcommand.
func Add(args []string) {
	fmt.Printf("add called with args: %s\n", args)
}
