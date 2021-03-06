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

// Package cmd implements the commands of Envelope.
package cmd

import (
	"fmt"
	"github.com/samWson/envelope/record"
)

// Add adds a new item to a budget depending on the sub-command.
func Add(args []string, tran *record.Transaction) {

	if len(args) < 1 {
		fmt.Println(help)
		return
	}

	switch command := args[0]; command {
	case "income":
		tran.Category = "income"
		Income(args[1:], tran)
	default:
		// Unrecognised args
		fmt.Println(help)
		return
	}

}

var help string = `Add (envelope add) adds items to a budget.

Usage:
    envelope add [item]

Examples:
    income    add new income to a budget`
