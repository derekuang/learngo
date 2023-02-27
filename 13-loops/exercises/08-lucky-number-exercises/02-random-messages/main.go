// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Random Messages
//
//  Display a few different won and lost messages "randomly".
//
// HINTS
//  1. You can use a switch statement to do that.
//  2. Set its condition to the random number generator.
//  3. I would use a short switch.
//
// EXAMPLES
//  The Player wins: then randomly print one of these:
//
//  go run main.go 5
//    YOU WON
//  go run main.go 5
//    YOU'RE AWESOME
//
//  The Player loses: then randomly print one of these:
//
//  go run main.go 5
//    LOSER!
//  go run main.go 5
//    YOU LOST. TRY AGAIN?
// ---------------------------------------------------------

const (
	maxTurns = 5 // less is more difficult
	usage    = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.
Your mission is to guess one of those numbers.

The greater your number is, harder it gets.

Wanna play?
`
	notNumErr = `Not a number.`
	notPosErr = `Please pick a positive number.`
	firstMsg  = `🥇 FIRST TIME WINNER!!!`
	winMsg1   = `🎉  YOU WON!`
	winMsg2   = `🎉  YOU'RE AWESOME!`
	lostMsg1  = `☠️  YOU LOST... Try again?`
	lostMsg2  = `☠️  LOSER!`
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	if len(args) != 1 {
		println(usage)
		return
	}

	guess, err := strconv.Atoi(args[0])

	if err != nil {
		println(notNumErr)
		return
	}

	if guess <= 0 {
		println(notPosErr)
	}

	for turn := 1; turn <= maxTurns; turn++ {
		res := rand.Intn(guess) + 1

		if guess != res {
			continue
		}

		switch rand.Intn(2) {
		case 0:
			println(winMsg1)
		case 1:
			println(winMsg2)
		}
		return
	}

	switch rand.Intn(2) {
	case 0:
		println(lostMsg1)
	case 1:
		println(lostMsg2)
	}
}
