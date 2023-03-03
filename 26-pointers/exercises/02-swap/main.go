// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Swap
//
//  Using funcs, swap values through pointers, and swap
//  the addresses of the pointers.
//
//
//  1- Swap the values through a func
//
//     a- Declare two float variables
//
//     b- Declare a func that can swap the variables' values
//        through their memory addresses
//
//     c- Pass the variables' addresses to the func
//
//     d- Print the variables
//
//
//  2- Swap the addresses of the float pointers through a func
//
//     a- Declare two float pointer variables and,
//        assign them the addresses of float variables
//
//     b- Declare a func that can swap the addresses
//        of two float pointers
//
//     c- Pass the pointer variables to the func
//
//     d- Print the addresses, and values that are
//        pointed by the pointer variables
//
// ---------------------------------------------------------

package main

import "fmt"

func main() {
	f1, f2 := 1., 2.
	swap(&f1, &f2)
	fmt.Printf("f1: %f f2: %f\n", f1, f2)

	p1, p2 := &f1, &f2
	p1, p2 = swapAddr(p1, p2)
	fmt.Printf("p1: %p %f p2: %p %f\n", p1, *p1, p2, *p2)
}

func swap(f1, f2 *float64) {
	*f1, *f2 = *f2, *f1
}

func swapAddr(p1, p2 *float64) (*float64, *float64) {
	return p2, p1
}
