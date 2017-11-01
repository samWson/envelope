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

package cmd

import (
	"fmt"
	"github.com/samWson/envelope/record"
	"strconv"
)

// Income adds income sources to a budget.
func Income(args []string, tran *record.Transaction) {

	if len(args) < 1 {
		fmt.Println(incomeHelp)
		return
	}

	amount, err := strconv.ParseFloat(args[0], 64)

	if err != nil {
		fmt.Printf("Amount requires a valid number.\n%v\n", err)
		return
	} else {
		tran.Amount = amount
	}

	fmt.Printf("Income called with the amount %s\n", args[0])
	fmt.Printf("Transaction created: %v\n", tran)
}

var incomeHelp string = `Income (envelope add income) adds income items to a budget.

Usage:
    envelope add income [amount]

Examples:
    1000    add a $1000.00 income to a budget`
