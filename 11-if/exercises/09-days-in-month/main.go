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
	"os"
	"strings"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Days in a Month
//
//  Print the number of days in a given month.
//
// RESTRICTIONS
//  1. On a leap year, february should print 29. Otherwise, 28.
//
//     Set your computer clock to 2020 to see whether it works.
//
//  2. It should work case-insensitive. See below.
//
//     Search on Google: golang pkg strings ToLower
//
//  3. Get the current year using the time.Now()
//
//     Search on Google: golang pkg time now year
//
//
// EXPECTED OUTPUT
//
//  -----------------------------------------
//  Your solution should not accept invalid months
//  -----------------------------------------
//  go run main.go
//    Give me a month name
//
//  go run main.go sheep
//    "sheep" is not a month.
//
//  -----------------------------------------
//  Your solution should handle the leap years
//  -----------------------------------------
//  go run main.go january
//    "january" has 31 days.
//
//  go run main.go february
//    "february" has 28 days.
//
//  go run main.go march
//    "march" has 31 days.
//
//  go run main.go april
//    "april" has 30 days.
//
//  go run main.go may
//    "may" has 31 days.
//
//  go run main.go june
//    "june" has 30 days.
//
//  go run main.go july
//    "july" has 31 days.
//
//  go run main.go august
//    "august" has 31 days.
//
//  go run main.go september
//    "september" has 30 days.
//
//  go run main.go october
//    "october" has 31 days.
//
//  go run main.go november
//    "november" has 30 days.
//
//  go run main.go december
//    "december" has 31 days.
//
//  -----------------------------------------
//  Your solution should be case insensitive
//  -----------------------------------------
//  go run main.go DECEMBER
//    "DECEMBER" has 31 days.
//
//  go run main.go dEcEmBeR
//    "dEcEmBeR" has 31 days.
// ---------------------------------------------------------

func main() {
	if len(os.Args) == 1 {
		println("Give me a month name")
	}

	var (
		now   = time.Now()
		year  = now.Year()
		month = os.Args[1]
		m     = strings.ToLower(month)
		f     = "%q has %d days.\n"
		days  = 0
	)

	if m == "january" {
		days = 31
	} else if m == "february" {
		if isLeapYear(year) {
			days = 29
		} else {
			days = 28
		}
	} else if m == "march" {
		days = 31
	} else if m == "april" {
		days = 30
	} else if m == "may" {
		days = 31
	} else if m == "june" {
		days = 30
	} else if m == "july" {
		days = 31
	} else if m == "august" {
		days = 31
	} else if m == "september" {
		days = 30
	} else if m == "october" {
		days = 31
	} else if m == "november" {
		days = 30
	} else if m == "december" {
		days = 31
	} else {
		fmt.Printf("%q is not a month.\n", month)
		return
	}

	fmt.Printf(f, month, days)
}

func isLeapYear(year int) bool {
	return ((year%4 == 0) && ((year%100 != 0) || (year%400 == 0)))
}
