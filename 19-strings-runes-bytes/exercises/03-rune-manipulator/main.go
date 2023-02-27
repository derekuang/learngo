// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"unicode/utf8"
)

// ---------------------------------------------------------
// EXERCISE: Rune Manipulator
//
//  Please read the comments inside the following code.
//
// EXPECTED OUTPUT
//  Please run the solution.
// ---------------------------------------------------------

func main() {
	words := []string{
		"cool",
		"güzel",
		"jīntiān",
		"今天",
		"read 🤓",
	}

	for _, word := range words {
		fmt.Printf("%s\n", word)

		// Print the byte and rune length of the strings
		// Hint: Use len and utf8.RuneCountInString
		fmt.Printf("\t\thas %d bytes and %d runes\n", len(word), utf8.RuneCountInString(word))

		// Print the bytes of the strings in hexadecimal
		// Hint: Use % x verb
		fmt.Printf("\t\t%-8s: % x\n", "bytes", []byte(word))

		// Print the runes of the strings in hexadecimal
		// Hint: Use % x verb
		fmt.Printf("\t\t%-8s:", "runes")
		for _, c := range word {
			fmt.Printf("% x", c)
		}
		fmt.Println()

		// Print the runes of the strings as rune literals
		// Hint: Use for range
		fmt.Printf("\t\t%-8s: ", "runes")
		for _, c := range word {
			fmt.Printf("'%c'", c)
		}
		fmt.Println()

		// Print the first rune and its byte size of the strings
		// Hint: Use utf8.DecodeRuneInString
		c, size := utf8.DecodeRuneInString(word)
		fmt.Printf("\t\t%-8s: '%c' (%d bytes)\n", "first", c, size)

		// Print the last rune of the strings
		// Hint: Use utf8.DecodeLastRuneInString
		c, size = utf8.DecodeLastRuneInString(word)
		fmt.Printf("\t\t%-8s: '%c' (%d bytes)\n", "last", c, size)

		// Slice and print the first two runes of the strings
		fmt.Printf("\t\t%-8s: \"%s\"\n", "first 2", string([]rune(word)[:2]))

		// Slice and print the last two runes of the strings
		fmt.Printf("\t\t%-8s: \"%s\"\n", "last 2", string([]rune(word)[utf8.RuneCountInString(word)-2:]))

		// Convert the string to []rune
		// Print the first and last two runes
		word1 := []rune(word)
		fmt.Printf("\t\t%-8s: \"%s\"\n", "first 2", string(word1[:2]))
		fmt.Printf("\t\t%-8s: \"%s\"\n", "last 2", string(word1[len(word1)-2:]))
	}
}
