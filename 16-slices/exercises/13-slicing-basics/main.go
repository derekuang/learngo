// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Slice the numbers
//
//   We've a string that contains even and odd numbers.
//
//   1. Convert the string to an []int
//
//   2. Print the slice
//
//   3. Slice it for the even numbers and print it (assign it to a new slice variable)
//
//   4. Slice it for the odd numbers and print it (assign it to a new slice variable)
//
//   5. Slice it for the two numbers at the middle
//
//   6. Slice it for the first two numbers
//
//   7. Slice it for the last two numbers (use the len function)
//
//   8. Slice the evens slice for the last one number
//
//   9. Slice the odds slice for the last two numbers
//
//
// EXPECTED OUTPUT
//   go run main.go
//
//     nums        : [2 4 6 1 3 5]
//     evens       : [2 4 6]
//     odds        : [1 3 5]
//     middle      : [6 1]
//     first 2     : [2 4]
//     last 2      : [3 5]
//     evens last 1: [6]
//     odds last 2 : [3 5]
//
//
// NOTE
//
//  You can also use my prettyslice package for printing the slices.
//
//
// HINT
//
//  Find a function in the strings package for splitting the string into []string
//
// ---------------------------------------------------------

import (
	"strconv"
	"strings"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	// uncomment the declaration below
	data := "2 4 6 1 3 5"

	nums := make([]int, 0, 6)
	for _, n := range strings.Fields(data) {
		num, _ := strconv.Atoi(n)
		nums = append(nums, num)
	}

	s.Show("nums", nums)

	var (
		evens, odds []int
	)

	for _, num := range nums {
		if num%2 == 0 {
			evens = append(evens, num)
		} else {
			odds = append(odds, num)
		}
	}

	s.Show("evens", evens)
	s.Show("odds", odds)
	s.Show("middle", nums[2:4])
	s.Show("first 2", nums[:2])
	s.Show("last 2", nums[len(nums)-2:])
	s.Show("evens last1", evens[len(evens)-1:])
	s.Show("odds last 2", odds[len(odds)-2:])
}

func init() {
	s.MaxPerLine = 6
}
