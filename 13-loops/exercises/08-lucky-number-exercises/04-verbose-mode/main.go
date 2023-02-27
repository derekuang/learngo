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
// EXERCISE: Verbose Mode
//
//  When the player runs the game like this:
//
//    go run main.go -v 4
//
//  Display each generated random number:

//    1 3 4 🎉  YOU WIN!
//
//  In this example, computer picks 1, 3, and 4. And the
//  player wins.
//
// HINT
//  You need to get and interpret the command-line arguments.
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

	verbose := false

	if len(args) == 2 {
		if args[0] == "-v" {
			verbose = true
		}
		args = args[1:]
	}

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

		if verbose {
			print(res, " ")
		}

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
