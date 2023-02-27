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
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Housing Prices and Averages
//
//  Use the previous exercise to solve this exercise (Housing Prices).
//
//  Your task is to print the averages of the sizes, beds, baths, and prices.
//
//
// EXPECTED OUTPUT
//
//  Location       Size           Beds           Baths          Price
//  ===========================================================================
//  New York       100            2              1              100000
//  New York       150            3              2              200000
//  Paris          200            4              3              400000
//  Istanbul       500            10             5              1000000
//
//  ===========================================================================
//                 237.50         4.75           2.75           425000.00
//
// ---------------------------------------------------------

func main() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	rows := strings.Split(data, "\n")

	var (
		locations                  []string
		sizes, beds, baths, prices []int
	)

	for _, row := range rows {
		item := strings.Split(row, separator)

		locations = append(locations, item[0])

		size, _ := strconv.Atoi(item[1])
		sizes = append(sizes, size)

		bed, _ := strconv.Atoi(item[2])
		beds = append(beds, bed)

		bath, _ := strconv.Atoi(item[3])
		baths = append(baths, bath)

		price, _ := strconv.Atoi(item[4])
		prices = append(prices, price)
	}

	for _, head := range strings.Split(header, separator) {
		fmt.Printf("%-15s", head)
	}
	fmt.Printf("\n%s\n", strings.Repeat("=", 75))

	for i := range locations {
		fmt.Printf("%-15s", locations[i])
		fmt.Printf("%-15d", sizes[i])
		fmt.Printf("%-15d", beds[i])
		fmt.Printf("%-15d", baths[i])
		fmt.Printf("%-15d\n", prices[i])
	}

	fmt.Printf("\n%s\n", strings.Repeat("=", 75))
	fmt.Printf("%-15s", "")

	datas := [][]int{sizes, beds, baths, prices}
	for _, data := range datas {
		sum := 0.
		for _, n := range data {
			sum += float64(n)
		}
		fmt.Printf("%-15.2f", sum/float64(len(data)))
	}
	println()
}
