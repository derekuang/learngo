// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Refactor
//
//   In the headerOf function, instead of using a switch,
//   use a map. Declare it at the package-level.
//
// RESTRICTION
//
//   For panicking, use a deferred function and a named param.
//
// EXPECTED OUTPUT
//
//   Please run the solution.
//
// ---------------------------------------------------------

func main() {
	fmt.Println(headerOf("gif"))
}

var headers = map[string]string{
	"png": "\x89PNG\r\n\x1a\n",
	"jpg": "\xff\xd8\xff",
}

func headerOf(format string) (head string) {
	defer func() {
		if head == "" {
			panic("unknown format: " + format)
		}
	}()

	return headers[format]
}
