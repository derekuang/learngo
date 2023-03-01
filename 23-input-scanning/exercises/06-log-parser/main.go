// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Log Parser from Stratch
//
//  You've watched the lecture. Now, try to create the same
//  log parser program on your own. Do not look at the lecture,
//  and the existing source code.
//
//
// EXPECTED OUTPUT
//
//  go run main.go < log.txt
//
//   DOMAIN                             VISITS
//   ---------------------------------------------
//   blog.golang.org                        30
//   golang.org                             10
//   learngoprogramming.com                 20
//
//   TOTAL                                  60
//
//
//  go run main.go < log_err_missing.txt
//
//   wrong input: [golang.org] (line #3)
//
//
//  go run main.go < log_err_negative.txt
//
//   wrong input: "-100" (line #3)
//
//
//  go run main.go < log_err_str.txt
//
//   wrong input: "FOUR" (line #3)
//
// ---------------------------------------------------------

func main() {
	in := bufio.NewScanner(os.Stdin)

	cnt := map[string]int{}
	total := 0

	line := 0
	for in.Scan() {
		line++

		records := strings.Fields(in.Text())
		domain := records[0]
		if len(records) != 2 {
			fmt.Printf("wrong input: [%s] (line #%d)\n", domain, line)
			return
		}

		visits, err := strconv.Atoi(records[1])
		if err != nil {
			fmt.Printf("wrong input: [%q] (line #%d)\n", records[1], line)
			return
		}

		if visits < 0 {
			fmt.Printf("wrong input: [%q] (line #%d)\n", records[1], line)
			return
		}

		total += visits
		cnt[domain] += visits
	}

	fmt.Printf("%-30s%10s\n", "DOMAIN", "VISITS")
	fmt.Printf("%s\n", strings.Repeat("-", 45))

	for domain, visits := range cnt {
		fmt.Printf("%-30s%10d\n", domain, visits)
	}
	fmt.Printf("\n%-30s%10d\n", "TOTAL", total)
}
