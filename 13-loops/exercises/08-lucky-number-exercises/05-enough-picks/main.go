// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Enough Picks
//
//  If the player's guess number is below 10;
//  then make sure that the computer generates a random
//  number between 0 and 10.
//
//  However, if the guess number is above 10; then the
//  computer should generate the numbers
//  between 0 and the guess number.
//
// WHY?
//  This way the game will be harder.
//
//  Because, in the current version of the game, if
//  the player's guess number is for example 3; then the
//  computer picks a random number between 0 and 3.
//
//  So, currently a player can easily win the game.
//
// EXAMPLE
//  Suppose that the player runs the game like this:
//    go run main.go 9
//
//  Or like this:
//    go run main.go 2
//
//    Then the computer should pick a random number
//    between 0-10.
//
//  Or, if the player runs it like this:
//    go run main.go 15
//
//    Then the computer should pick a random number
//    between 0-15.
// ---------------------------------------------------------

const (
	maxTurns  = 5 // less is more difficult
	minBorder = 10
	usage     = `Welcome to the Lucky Number Game! 🍀

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

	border := math.Max(float64(guess), float64(minBorder))
	for turn := 1; turn <= maxTurns; turn++ {
		res := rand.Intn(int(border)) + 1

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
