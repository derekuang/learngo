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
// EXERCISE: Double Guesses
//
//  Let the player guess two numbers instead of one.
//
// HINT:
//  Generate random numbers using the greatest number
//  among the guessed numbers.
//
// EXAMPLES
//  go run main.go 5 6
//  Player wins if the random number is either 5 or 6.
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

	if len(args) != 2 {
		println(usage)
		return
	}

	guess1, err1 := strconv.Atoi(args[0])
	guess2, err2 := strconv.Atoi(args[1])

	if err1 != nil || err2 != nil {
		println(notNumErr)
		return
	}

	if guess1 <= 0 || guess2 <= 0 {
		println(notPosErr)
	}

	var g int
	if guess1 > guess2 {
		g = guess1
	} else {
		g = guess2
	}

	for turn := 1; turn <= maxTurns; turn++ {
		res := rand.Intn(g) + 1

		if guess1 != res && guess2 != res {
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
